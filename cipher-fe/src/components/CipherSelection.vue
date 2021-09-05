<template>
  <div class="flex flex-col justify-start items-start text-left space-y-1 ">
    <div
      v-for="(cipherType, i) in cipherTypes"
      :key="cipherType"
      class=" text-lg select-none cursor-pointer transition-all duration-300"
      :class="[store.modeIndex === i ? 'font-semibold cursor-default text-indigo-400 text-lg' : 'opacity-25 text-sm']"
      @click="() => {
        store.modeIndex = i
        store.mode = cipherTypes[i]
      }"
    >
      <div class="relative inline-flex items-center">
        <carbon-hashtag v-if="store.modeIndex === i" class="-left-6 absolute" />
        {{ cipherType }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useMainStore } from '~/store'
import { cipherTypes } from '~/constants'

// STORE
const store = useMainStore()

</script>
