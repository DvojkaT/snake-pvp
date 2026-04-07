<script setup lang="ts">
import GameCanvas from "@/GameCanvas.vue";
import {getAuth} from "@/lib/auth.ts";
import {useRoute} from "vue-router";
import {onMounted, ref} from "vue";
import axios from "axios";

const route = useRoute()

const uuid = getAuth()
const gameUuid = route.params.gameId
const isConnected = ref(false)

onMounted(() => {
  join().then(() => {
    isConnected.value = true
  })
})

async function join() {
  const formData = ref({user_id: getAuth(), name: "test-name"});
  try {
    return await axios.post(`http://127.0.0.1:8080/${gameUuid}/join`, formData.value);
  } catch (error) {
    console.error(error);
  }
}

async function start() {
  try {
    return await axios.post(`http://127.0.0.1:8080/${gameUuid}/start`);
  } catch (error) {
    console.error(error);
  }
}
</script>

<template>
  <div class="top-0 absolute flex flex-col gap-2">
    <span>Подключен ли к игре: {{ isConnected ? "Да" : "Нет"}}</span>
    <span>Идентификатор игры: {{ gameUuid }} </span>
    <span>Ваш идентификатор: {{ uuid }}</span>
    <button @click.prevent="start">
      Начать
    </button>
  </div>
  <div
    class="min-h-full min-w-full flex flex-col gap-10 items-center content-center justify-center mt-20">
    <h1 class="font-bold text-2xl">Змейка</h1>
    <GameCanvas/>
  </div>
</template>

<style scoped>

</style>
