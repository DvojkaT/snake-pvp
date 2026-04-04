import {Centrifuge} from "centrifuge";
import {renderCanvas} from "@/lib/render.ts";

export function connect(roomId: string) {
  const centrifuge = new Centrifuge('ws://127.0.0.1:8080/connection/websocket', {
    data: {
      user_id: "asd"
    }
  });

  const sub = centrifuge.newSubscription(roomId);

  sub.subscribe();
  centrifuge.connect();

  return { sub }
}
