<template>
  <div class="text-center flex flex-row justify-between border-teal-200 border-2px w-[100%]">
    <Menu :active-mode="menuState.activeMode" :active-type="menuState.activeType" @mode-clicked="handleModeChange" @type-clicked="handleTypeChange"></Menu>
    <div class="flex flex-grow bg-teal-100 ml-4">
      <div>LADSD</div>
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
</script>


