<template>
  <div class="flex flex-col w-[20%] m-6">
    <!-- HEADER SHOULD BE HERE -->
    <!-- MODE CONTAINER   -->
    <div class="flex flex-row justify-evenly gap-4">
      <div 
        v-for="mode in MODE"
        :key="mode"
        @click="emit('modeClicked', mode)"
        class="
          hover:cursor-pointer
          p-2 border-2px
          border-teal-200
          rounded-xl
          hover:bg-teal-100
          :hover:bg-black
          font-semibold
          font-sans
          select-none"
        :class="[
          mode === props.activeMode ?
            'bg-teal-600 text-light-600 hover:text-black' : ''
        ]">{{ mode }}</div>
    </div>
    <!-- TYPE CONTAINER -->
    <div class="grid auto-cols-auto gap-1 mt-4">
      <div
        v-for="cipherType in TYPE"
        :key="cipherType"
        @click="emit('typeClicked', cipherType)"
        class="
          hover:cursor-pointer
          p-2
          border-2px
          border-teal-200
          rounded-xl
          hover:bg-teal-100
          font-semibold
          font-sans
          select-none
          mt-2"
          :class="[
            cipherType === props.activeType ?
              'bg-teal-600 text-light-600 hover:text-black' : ''
          ]">{{ cipherType }}</div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { MODE, TYPE } from '~/components/constant'

const props = defineProps({
  activeMode: {
    type: String,
    default: MODE.ENCRYPT
  },
  activeType: {
    type: String,
    default: TYPE.STANDARD_VIGENERE
  }
});

const emit = defineEmits(['modeClicked', 'typeClicked'])

</script>
