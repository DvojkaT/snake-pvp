<script setup lang="ts">
import {getAuth, getName, setName} from "@/lib/auth.ts";
import {ref} from "vue";
import axios from "axios";
import {useRouter} from "vue-router";

const router = useRouter();

const name = ref<string>(getName())
const joinGameUuid = ref<string>()
const currentName = ref<string>(getName())

//errors
const connectionError = ref<string | null>(null)

function setUsername() {
  setName(name.value)
  currentName.value = getName();
}

async function createGame() {
  try {
    const {data: {data}} = await axios.post(`http://127.0.0.1:8080/create`);
    return router.push(`/${data.uuid}`)
  } catch (error) {
    console.error(error);
  }
}

// Подключение идемпотентное, так что можно и так обойтись
async function checkRoom() {
  connectionError.value = null;
  const formData = ref({user_id: getAuth(), name: getName()});
  axios.post(`http://127.0.0.1:8080/${joinGameUuid.value}/join`, formData.value).then(() => {
    return router.push(`/${joinGameUuid.value}`);
  }).catch(error => {
    console.error(error);
    connectionError.value = "Не удалось подключиться. Проверьте идентификатор";
  })
}

</script>

<template>
  <div class="flex flex-col md:grid md:grid-cols-4">
    <div class="bg-gray-300 shadow-lg border-gray-400 border flex flex-col rounded-2xl p-4 m-4">
      <h3 class="text-xl font-bold text-center mb-4">
        Информация об игроке
      </h3>
      <div class="flex flex-col gap-4 border-b">
        <div class="flex flex-col">
          <span class="text-nowrap">Ваш идентификатор: </span>
          <span class="font-semibold break-all">{{ getAuth() }}</span>
        </div>
        <div class="flex flex-col">
          <span>Ваше имя: </span>
          <span class="font-semibold break-all">{{ currentName }}</span>
        </div>
      </div>
      <div>
        <form class="flex flex-col gap-4" @submit.prevent="setUsername">
          <label class="text-lg font-bold text-center mt-2">
            Смена имени
          </label>
          <input class="text-center rounded-2xl border-gray-400 bg-gray-200 border p-2"
                 v-model="name" placeholder="Имя">
          <button
            class="shadow-lg rounded-2xl border-gray-400 bg-gray-200 border p-2 cursor-pointer hover:bg-gray-500 transition">
            Применить
          </button>
        </form>
      </div>
    </div>
    <div class="bg-gray-300 shadow-lg border-gray-400 border flex flex-col rounded-2xl p-4 m-4">
      <h3 class="text-xl font-bold text-center mb-4">
        Меню
      </h3>
      <div class="flex flex-col gap-4 border-b">
        <button
          class="mb-17 rounded-2xl border-gray-400 bg-gray-200 border p-2 cursor-pointer hover:bg-gray-500 transition"
          @click.prevent="createGame">Начать игру
        </button>
      </div>
      <div>
        <form class="flex flex-col gap-4" @submit.prevent="checkRoom">
          <label class="text-lg font-bold text-center mt-2">
            Присоединиться к существующей
          </label>
          <input class="text-center rounded-2xl border-gray-400 bg-gray-200 border p-2"
                 v-model="joinGameUuid" placeholder="Идентификатор игры">
          <button
            class="shadow-lg rounded-2xl border-gray-400 bg-gray-200 border p-2 cursor-pointer hover:bg-gray-500 transition">
            Присоединиться
          </button>
          <span v-show="connectionError"
                class="text-center text-red-500">{{ connectionError }}</span>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
