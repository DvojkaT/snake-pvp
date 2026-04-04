import {type Cell, objectTypes} from "@/types";

export function renderCanvas(ctx: CanvasRenderingContext2D, cells: Cell[][], cellSize: number) {
  for (let x = 0; x < cells.length; x++) {
    for (let y = 0; y < cells[x]!.length; y++) {
      switch (cells[x]![y]!.object) {
        case objectTypes.snakePart: {
          ctx.fillStyle = '#000000';
          ctx.fillRect(x * cellSize, y * cellSize, cellSize, cellSize);
          break;
        }
        case objectTypes.fruit: {
          ctx.fillStyle = '#ffcd00';
          ctx.fillRect(x * cellSize, y * cellSize, cellSize, cellSize);
          break;
        }
        case objectTypes.empty: {
          ctx.fillStyle = '#ffffff';
          ctx.fillRect(x * cellSize, y * cellSize, cellSize, cellSize);
          break;
        }
      }
    }
  }
}
