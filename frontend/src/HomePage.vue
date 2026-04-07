<script setup lang="ts">
import {getAuth, getName, setName} from "@/lib/auth.ts";
import {ref} from "vue";

const name = ref<string>(getName())
const currentName = ref<string>(getName())

function setUsername() {
  setName(name.value)
  currentName.value = getName();
}
</script>

<template>
  <div class="grid grid-cols-4">
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
          <input class="text-center rounded-2xl border-gray-400 bg-gray-200 border p-2" v-model="name" placeholder="Введите имя">
          <button class="rounded-2xl border-gray-400 bg-gray-200 border p-2 cursor-pointer hover:bg-gray-500 transition">Применить</button>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
