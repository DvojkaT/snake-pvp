<script setup lang="ts">

import {onMounted, onUnmounted, ref, useTemplateRef, watch} from "vue";
import {connect} from "@/lib/ws.ts";
import {renderCanvas} from "@/lib/render.ts";
import {type CellSub, directions} from "@/types";
import {Subscription} from "centrifuge";
import {getAuth} from "@/lib/auth.ts";
import {useRouter} from "vue-router";
import {useSwipe} from "@vueuse/core";

const canvasRef = ref<HTMLCanvasElement>();
const sub = ref<Subscription | null>(null)
const props = defineProps<{
  uuid: string
}>()
const emit = defineEmits(['gameStarted', 'setPlayers'])

const {direction} = useSwipe(canvasRef)
const winner = ref<string>()
const router = useRouter()

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

watch(direction, (newDirection) => {
  switch (newDirection) {
    case "up":
      sub.value?.publish({type: "snake_move", direction: directions.up});
      break;
    case "down":
      sub.value?.publish({type: "snake_move", direction: directions.down});
      break;
    case "left":
      sub.value?.publish({type: "snake_move", direction: directions.left});
      break;
    case "right":
      sub.value?.publish({type: "snake_move", direction: directions.right});
      break;
  }
})

onMounted(() => {
  if (canvasRef.value === undefined) {
    console.log("Can't find canvas");
    return
  }

  const {sub: subscription} = connect(`room:${props.uuid}`, getAuth())
  sub.value = subscription

  const ctx = canvasRef.value.getContext("2d");
  if (!ctx) {
    return;
  }

  const cellSize = 10;

  canvasRef.value.width = 50 * cellSize;
  canvasRef.value.height = 50 * cellSize;

  sub.value.on('publication', function (data: CellSub) {
    if (data.data.event === "room_state") {
      emit('gameStarted')
      renderCanvas(ctx, data.data.data.cells, cellSize);
    }
    if (data.data.event === "lobby_state") {
      emit('setPlayers', data.data.data.players)
      if (data.data.data.winner != null) {
        winner.value = data.data.data.winner.name
      }
    }
  });

  window.addEventListener('keydown', handleKeyDown);
})

onUnmounted(() => {
  window.removeEventListener("keydown", handleKeyDown);
})

function toMain() {
  return router.push("/")
}

</script>

<template>
  <canvas ref="canvasRef" class="touch-none relative aspect-square w-full border md:w-[500px] md:h-[500px]">
  </canvas>
  <div v-show="winner" class="absolute flex w-full h-full flex-row items-center justify-center">
    <div
      class="rounded-2xl w-[50%] md:w-[30%] h-[30%] bg-gray-50 border-1 shadow-2xl">
      <div class="p-4 flex flex-col gap-4">
        <span class="text-xl"> Победитель: {{ winner ?? "не найден" }}</span>
        <button @click.prevent="toMain"
                class="shadow-lg rounded-2xl border-gray-400 bg-gray-200 border p-2 cursor-pointer hover:bg-gray-500 transition">
          Выйти
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
