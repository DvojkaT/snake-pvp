<script setup lang="ts">
import GameCanvas from "@/GameCanvas.vue";
import {getAuth, getName} from "@/lib/auth.ts";
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import axios from "axios";

const route = useRoute()
const router = useRouter()

const uuid = getAuth()
const gameUuid = route.params.gameId as string

const isConnected = ref(false)
const isGameStarted = ref(false)

onMounted(() => {
  join().then(() => {
    isConnected.value = true
  })
})

async function join() {
  const formData = ref({user_id: getAuth(), name: getName()});
  try {
    return await axios.post(`http://127.0.0.1:8080/${gameUuid}/join`, formData.value);
  } catch (error) {
    console.error(error);
    return router.push("/");
  }
}

async function start() {
  try {
    return await axios.post(`http://127.0.0.1:8080/${gameUuid}/start`);
  } catch (error) {
    console.error(error);
  }
}

function exit() {
  return router.push('/')
}

function gameStarted() {
  isGameStarted.value = true
}
</script>

<template>
  <div class="top-0 absolute grid grid-cols-4">
    <div class="bg-gray-300 shadow-lg border-gray-400 border flex flex-col rounded-2xl p-4 m-4">
      <div class="border-b">
        <button @click.prevent="exit"
                class="min-w-[40%] my-4 shadow-lg rounded-2xl border-gray-400 bg-gray-200 border p-2 cursor-pointer hover:bg-gray-500 transition">
          Выйти
        </button>
      </div>
      <h3 class="text-xl font-bold text-center my-4">
        Информация об игре
      </h3>
      <div class="flex flex-col gap-4">
        <div class="flex flex-col">
          <span class="text-nowrap">Подключен ли к игре: </span>
          <span class="font-semibold break-all">{{ isConnected ? "Да" : "Нет" }}</span>
        </div>
        <div class="flex flex-col">
          <span class="text-nowrap">Идентификатор игры: </span>
          <span class="font-semibold break-all">{{ gameUuid }}</span>
        </div>
        <div class="flex flex-col">
          <span class="text-nowrap">Ваш идентификатор: </span>
          <span class="font-semibold break-all">{{ uuid }}</span>
        </div>
        <div class="flex flex-col">
          <span>Ваше имя: </span>
          <span class="font-semibold break-all">{{ getName() }}</span>
        </div>
        <button v-if="!isGameStarted" @click.prevent="start"
                class="shadow-lg rounded-2xl border-gray-400 bg-gray-200 border p-2 cursor-pointer hover:bg-gray-500 transition">
          Начать
        </button>
      </div>
    </div>
  </div>
  <div
    class="min-h-full min-w-full flex flex-col gap-10 items-center content-center justify-center mt-20">
    <h1 class="font-bold text-2xl">Змейка</h1>
    <GameCanvas @game-started="gameStarted" :uuid="gameUuid"/>
  </div>
</template>

<style scoped>

</style>
