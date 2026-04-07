import {Centrifuge} from "centrifuge";

export function connect(roomId: string, userId: string) {
  const centrifuge = new Centrifuge(`${import.meta.env.VITE_WS_URL}/connection/websocket`, {
    data: {
      user_id: userId
    }
  });

  const sub = centrifuge.newSubscription(roomId);

  sub.subscribe();
  centrifuge.connect();

  return { sub }
}
