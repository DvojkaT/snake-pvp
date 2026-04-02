package ws

import (
	"context"
	"encoding/json"

	"github.com/centrifugal/centrifuge"
)

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
}
