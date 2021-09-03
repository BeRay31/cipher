<template>
  <div class="text-center flex flex-row justify-between border-teal-200 border-2px w-[100%]">
    <Menu
      :active-mode="menuState.activeMode"
      :active-type="menuState.activeType"
      @mode-clicked="handleModeChange"
      @type-clicked="handleTypeChange"
    ></Menu>
    <div class="flex flex-grow justify-center items-center bg-teal-100">
      <!-- TITLE -->
      <div class="flex flex-row justify-between items-center">
        <div class="font-bold text-2xl font-sans">{{ capitalize(title) }}</div>
      </div>
      

    </div>
  </div>
</template>

<script setup lang="ts">
import { MODE, TYPE } from '~/components/constant'
import Menu from '~/components/Menu.vue'
// TYPE
type MenuType = {
  activeMode: string,
  activeType: string
}

const menuState: MenuType = reactive({
  activeMode: MODE.ENCRYPT,
  activeType: TYPE.STANDARD_VIGENERE
})
const handleTypeChange = (type: string) => { menuState.activeType = type }
const handleModeChange = (mode: string) => { menuState.activeMode = mode }
const title = computed(() => `${menuState.activeType} ${menuState.activeMode === MODE.ENCRYPT ? 'ENCRYPTION' : 'DECRYPTION'}`)
const capitalize = (str: string) => {
  return str.toLowerCase().split(" ").map(el => el.charAt(0).toUpperCase() + el.slice(1)).join(" ")
}
</script>


