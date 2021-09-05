<template>
  <div class="container mx-auto sm:(p-2 w-full) md:(p-6 w-2/3) lg:(p-8 w-1/2)">
    <div class="flex flex-row justify-evenly items-center w-full mb-8">
      <div class="w-1/3">
        <cipher-selection />
      </div>
      <div class="w-2/3">
        <mode-selection />
      </div>
    </div>
    <div v-if="mainStore.resultString !== ''" class="w-full">
      <result />
    </div>
  </div>
</template>

<script setup lang="ts">import { useMainStore } from '~/store'

const mainStore = useMainStore()
</script>
