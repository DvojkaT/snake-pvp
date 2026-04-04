<script setup lang="ts">

import {onMounted, ref} from "vue";
import {connect} from "@/lib/ws.ts";
import {renderCanvas} from "@/lib/render.ts";
import type {CellSub} from "@/types";

const canvasRef = ref<HTMLCanvasElement>();

onMounted(() => {
  if (canvasRef.value === undefined) {
    return
  }

  const { sub } = connect("room:test-game-id")

  const ctx = canvasRef.value.getContext("2d");
  if (!ctx) {
    return;
  }

  const cellSize = 10;

  canvasRef.value.width = 50 * cellSize;
  canvasRef.value.height = 50 * cellSize;

  sub.on('publication', function (data: CellSub) {
    renderCanvas(ctx, data.data.cells, cellSize);
  });
})

</script>

<template>
  <canvas ref="canvasRef" class="border w-[500px] h-[500px]">

  </canvas>
</template>

<style scoped>

</style>
