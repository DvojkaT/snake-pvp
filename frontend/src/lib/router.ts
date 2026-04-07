import {createMemoryHistory, createRouter, createWebHistory} from 'vue-router'

import HomePage from "@/HomePage.vue";
import GamePage from "@/GamePage.vue";

const routes = [
  {path: '/', component: HomePage},
  {path: '/:gameId', component: GamePage},
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})
