<template>
  <div class="flex flex-col justify-start items-start text-left space-y-1 ">
    <div
      v-for="(cipherType, i) in cipherTypes"
      :key="cipherType"
      class=" text-lg select-none cursor-pointer"
      :class="[store.modeIndex === i ? 'text-primary font-semibold cursor-default text-indigo-400' : 'text-alternate opacity-25']"
      @click="() => {
        store.modeIndex = i
        store.mode = cipherTypes[i]
      }"
    >
      <div class="relative inline-flex items-center">
        <carbon-hashtag v-if="store.modeIndex === i" class="-left-8 absolute" />
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
