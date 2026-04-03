package ws

import (
	"context"
	"encoding/json"
	"snake/internal/game"

	"github.com/centrifugal/centrifuge"
)

const RoomPrefix = "room:"

func HandleConnection(node *centrifuge.Node) {
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
