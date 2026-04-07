<template>
  <main>
    <RouterView />
  </main>
</template>

<script setup lang="ts">
import {getAuth, getName} from "@/lib/auth.ts";
getAuth()
getName()
</script>

<style scoped></style>
