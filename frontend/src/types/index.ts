export const objectTypes = {
  empty: 0,
  fruit: 1,
  snakePart: 2,
} as const;

export const directions = {
  up: 0,
  down: 1,
  left: 2,
  right: 3,
} as const;

export const statuses = {
  lobby: 0,
  active: 1,
  ended: 2
} as const;

export type Player = {
  name: string
  id: string
}

export type Cell = {
  object: number,
  snake: Snake,
}

export type Snake = {
  color: string
}

export type SubMessage = LobbyState | RoomState

export type LobbyState = {
  data: {
    event: "lobby_state",
    data: {
      players: Player[]
      winner: Player
      status_id: number
    }
  }
}

export type RoomState = {
  data: {
    event: "room_state",
    data: {
      id: string,
      cells: Cell[][],
    }
  }
}
