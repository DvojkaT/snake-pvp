<script setup lang="ts">
import GameCanvas from "@/GameCanvas.vue";
import {getAuth, getName} from "@/lib/auth.ts";
import {useRoute, useRouter} from "vue-router";
import {computed, onMounted, ref} from "vue";
import axios from "axios";
import {type Player, statuses} from "@/types";

const route = useRoute()
const router = useRouter()
const siteUrl = import.meta.env.VITE_API_URL

const uuid = getAuth()
const gameUuid = route.params.gameId as string

const isConnected = ref(false)

const playersList = ref<Player[]>([])
const statusId = ref<number>()

//todo Правильно распределить данные между GameCanvas и GamePage

onMounted(() => {
  join().then(({data: {data}}) => {
    isConnected.value = true
    playersList.value = data.players
    statusId.value = data.status_id
  })
})

const statusText = computed(() => {
  switch (statusId.value) {
    case statuses.lobby:
      return "Лобби"
    case statuses.active:
      return "Активно"
    case statuses.ended:
      return "Закончен"
    default:
      return "Неизвестно"
  }
})

async function join() {
  const formData = ref({user_id: getAuth(), name: getName()});
  try {
    return await axios.post(`${siteUrl}/${gameUuid}/join`, formData.value);
  } catch (error) {
    console.error(error);
    return router.push("/");
  }
}

async function start() {
  try {
    const data = await axios.post(`${siteUrl}/${gameUuid}/start`);
    setStatus(statuses.active)
    return data
  } catch (error) {
    console.error(error);
  }
}

function exit() {
  return router.push('/')
}

function setStatus(status: number) {
  statusId.value = status
}

function setPlayersList(players: string[]) {
  playersList.value = players
}
</script>

<template>
  <div class="md:top-0 md:absolute md:grid md:grid-cols-4">
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
          <span class="text-nowrap">Статус игры: </span>
          <span class="font-semibold break-all">{{ statusText }}</span>
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
        <button v-if="statusId == statuses.lobby" @click.prevent="start"
                class="shadow-lg rounded-2xl border-gray-400 bg-gray-200 border p-2 cursor-pointer hover:bg-gray-500 transition">
          Начать
        </button>
      </div>
    </div>
  </div>
  <div
    class="min-h-full min-w-full flex flex-col gap-10 items-center content-center justify-center mt-20">
    <h1 class="font-bold text-2xl">Змейка</h1>
    <GameCanvas @set-players="setPlayersList" @set-status="setStatus" :status="statusId"
                :uuid="gameUuid"/>
  </div>
  <div class="md:top-0 md:right-0 md:absolute">
    <div class="bg-gray-300 shadow-lg border-gray-400 border flex flex-col rounded-2xl p-4 m-4">
      <h3 class="text-xl font-bold text-center my-4">
        Список игроков
      </h3>
      <div class="flex flex-col gap-4">
        <ul class="flex flex-col">
          <li v-for="player in playersList" class="text-nowrap">- {{ player.name }}</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
