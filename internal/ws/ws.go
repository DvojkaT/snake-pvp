package ws

import (
	"context"
	"encoding/json"
	"log"
	"snake/internal/game"
	"strings"

	"github.com/centrifugal/centrifuge"
)

type UserMessage struct {
	Type string `json:"type"`
}

type SnakeDirectionMessage struct {
	Direction game.Direction `json:"direction"`
}

const RoomPrefix = "room:"

func HandleConnection(node *centrifuge.Node, list game.RoomList) {
	node.OnConnecting(func(ctx context.Context, e centrifuge.ConnectEvent) (centrifuge.ConnectReply, error) {
		dataType := struct {
			UserID string `json:"user_id"`
		}{}
		data := e.Data

		if err := json.Unmarshal(data, &dataType); err != nil {
			return centrifuge.ConnectReply{}, centrifuge.ErrorBadRequest
		}

		if dataType.UserID == "" {
			return centrifuge.ConnectReply{}, centrifuge.ErrorBadRequest
		}

		return centrifuge.ConnectReply{
			Credentials: &centrifuge.Credentials{
				UserID: dataType.UserID,
			},
		}, nil
	})

	node.OnConnect(func(client *centrifuge.Client) {
		client.OnSubscribe(func(event centrifuge.SubscribeEvent, callback centrifuge.SubscribeCallback) {
			//todo Сделать проверку на подключение к каналу
			if event.Channel == "" {
				callback(centrifuge.SubscribeReply{}, centrifuge.ErrorPermissionDenied)
				return
			}

			callback(centrifuge.SubscribeReply{
				Options: centrifuge.SubscribeOptions{
					Data: []byte(`{"msg": "welcome"}`),
				},
			}, nil)
		})

		client.OnPublish(func(event centrifuge.PublishEvent, callback centrifuge.PublishCallback) {
			log.Printf("[user %s] publishes into channel %s: %s", client.UserID(), event.Channel, string(event.Data))

			if !client.IsSubscribed(event.Channel) {
				callback(centrifuge.PublishReply{}, centrifuge.ErrorPermissionDenied)
				return
			}

			var msg UserMessage
			err := json.Unmarshal(event.Data, &msg)
			if err != nil {
				log.Printf("[user %s] error: %s", client.UserID(), err)
				callback(centrifuge.PublishReply{}, centrifuge.ErrorInternal)
				return
			}

			switch msg.Type {
			case "snake_move":
				{
					var msg SnakeDirectionMessage
					if err = json.Unmarshal(event.Data, &msg); err != nil {
						log.Printf("[user %s] error: %s", client.UserID(), err)
						callback(centrifuge.PublishReply{}, centrifuge.ErrorInternal)
						return
					}

					roomId := strings.TrimPrefix(event.Channel, RoomPrefix)
					room, ok := list[roomId]
					if !ok {
						log.Printf("[user %s] error. Room not found: %s", client.UserID(), err)
						callback(centrifuge.PublishReply{}, centrifuge.ErrorInternal)
						return
					}
					snake := room.Snakes[client.UserID()]
					snake.SetDirection(msg.Direction)
				}
			}

		})
	})
}

func PublishRoomState(node *centrifuge.Node, room *game.RoomView) error {
	data, err := json.Marshal(room)
	if err != nil {
		return err
	}

	_, err = node.Publish(RoomPrefix+room.ID, data)
	if err != nil {
		return err
	}

	return nil
}
