import {Centrifuge} from "centrifuge";

export function connect () {
  const centrifuge = new Centrifuge('ws://127.0.0.1:8080/connection/websocket', {
    data: {
      user_id: "asd"
    }
  });

  const sub = centrifuge.newSubscription('room:test-game-id');

  sub.on('publication', function(ctx) {
    console.log(ctx.data);
  });

// Trigger subscribe process.
  sub.subscribe();

// Trigger actual connection establishement.
  centrifuge.connect();
}
