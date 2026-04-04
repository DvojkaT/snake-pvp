<script setup lang="ts">

import {onMounted, onUnmounted, ref} from "vue";
import {connect} from "@/lib/ws.ts";
import {renderCanvas} from "@/lib/render.ts";
import {type CellSub, directions} from "@/types";
import {Subscription} from "centrifuge";

const canvasRef = ref<HTMLCanvasElement>();
const sub = ref<Subscription | null>(null)

const handleKeyDown = (event: KeyboardEvent) => {
  if (["ArrowUp", "ArrowDown", "ArrowLeft", "ArrowRight"].includes(event.key)) {
    event.preventDefault();
  }
  switch (event.key) {
    case "ArrowUp":
      sub.value?.publish({type: "snake_move", direction: directions.up});
      break;
    case "ArrowDown":
      sub.value?.publish({type: "snake_move", direction: directions.down});
      break;
    case "ArrowLeft":
      sub.value?.publish({type: "snake_move", direction: directions.left});
      break;
    case "ArrowRight":
      sub.value?.publish({type: "snake_move", direction: directions.right});
      break;
  }
}

onMounted(() => {
  if (canvasRef.value === undefined) {
    return
  }

  const {sub: subscription} = connect("room:test-game-id")
  sub.value = subscription

  const ctx = canvasRef.value.getContext("2d");
  if (!ctx) {
    return;
  }

  const cellSize = 10;

  canvasRef.value.width = 50 * cellSize;
  canvasRef.value.height = 50 * cellSize;

  sub.value.on('publication', function (data: CellSub) {
    renderCanvas(ctx, data.data.cells, cellSize);
  });

  window.addEventListener('keydown', handleKeyDown);
})

onUnmounted(() => {
  window.removeEventListener("keydown", handleKeyDown);
})

</script>

<template>
  <canvas ref="canvasRef" class="border w-[500px] h-[500px]">

  </canvas>
</template>

<style scoped>

</style>
