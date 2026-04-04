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

export type Cell = {
  object: number,
  snake: Snake,
}

export type Snake = {
  color: string
}

export type CellSub = {
  data: {
    cells: Cell[][]
  }
}
