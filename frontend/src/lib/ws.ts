import {Centrifuge} from "centrifuge";

export function connect(roomId: string) {
  const centrifuge = new Centrifuge('ws://127.0.0.1:8080/connection/websocket', {
    data: {
      user_id: "testPlayerId"
    }
  });

  const sub = centrifuge.newSubscription(roomId);

  sub.subscribe();
  centrifuge.connect();

  return { sub }
}
