<template>
  <!-- MAIN-CONTAINER -->
  <div class="flex flex-col justify-center items-center p-12">
    <!-- INPUT-CONTAINER -->
    <div class="mt-4 w-[100%] flex flex-col justify-center items-center">
      <p class="text-3xl font-poppins text-primary mb-16 font-medium">{{ capitalize(title) }}</p>
      <!-- CONTENT-CONTAINER -->
      <div class="flex flex-row justify-center items-center w-[100%] gap-x-60">
        <!-- MENU-CONTAINER -->
        <div class="flex flex-col justify-center items-center gap-4 ">
          <div
            v-for="cipherType in types"
            :key="cipherType"
            class="font-spartan text-xl select-none cursor-pointer"
            :class="[menuState.activeType == cipherType ? 'text-primary font-semibold cursor-default' : 'text-alternate font-medium']"
            @click="() => handleTypeChange(cipherType)"
          >{{ capitalize(cipherType) }}</div>
        </div>
        <!-- INPUT-CONTENT-CONTAINER -->
        <div class="bg-secondary rounded-2xl flex flex-col p-6 py-10 justify-center items-center">
          <!-- MODE-CONTAINER -->
          <div class="flex flex-row justify-center items-center gap-x-20 mb-20">
            <div
              v-for="mode in menus"
              :key="mode"
              class="font-spartan text-xl select-none cursor-pointer pb-2"
              :class="[menuState.activeMode === mode ? 'text-secondary font-semibold border-primary border-b-[2px] border-current' : 'text-alternate font-medium']"
              @click="() => handleModeChange(mode)"
            >
              {{ capitalize(mode) }}
            </div>
          </div>
          <!-- INPUT-CONTAINER -->
          <div class="flex flex-col justify-center items-center">
            <input
              type="text"
              v-model="text"
              :placeholder="plainPlaceholder"
              class="px-2 py-2 rounded min-w-[35rem] font-spartan font-normal focus:outline-none">

              <p class="text-alternate font-spartan text-xl my-10">or</p>
              <!-- INPUT-FILE-CONTAINER -->
              <div
                class="flex flex-row text-secondary border-current border-[2px] rounded-md py-2 px-4 gap-x-2 select-none cursor-pointer hover:bg-current min-w-52 justify-center items-center"
                @click="(e) => handleClickInputFile(e, INPUT.TEXT)"
              >
                <img src="../assets/foundation_upload-cloud.svg">
                <p class="text-primary font-spartan font-normal text-base">{{ filePlaceholder }}</p>
              </div>
              <input id="hiddenInputFile1" type="file" class="hidden" @change="() => handleInputFile(1)">
          </div>
          <div class="my-4"></div>
          <div class="flex flex-col justify-center items-center">
            <input
              type="text"
              v-model="key"
              placeholder="Input key here"
              class="px-2 py-2 rounded min-w-[35rem] font-spartan font-normal focus:outline-none">

              <p class="text-alternate font-spartan text-xl my-10">or</p>
              <!-- INPUT-FILE-CONTAINER -->
              <div
                class="flex flex-row text-secondary border-current border-[2px] rounded-md py-2 px-4 gap-x-2 select-none cursor-pointer hover:bg-current min-w-52 justify-center items-center"
                @click="(e) => handleClickInputFile(e, INPUT.KEY)"
              >
                <img src="../assets/foundation_upload-cloud.svg">
                <p class="text-primary font-spartan font-normal text-base">{{ filePlaceholder }}</p>
              </div>
              <input id="hiddenInputFile2" type="file" class="hidden" @change="() => handleInputFile(2)">
          </div>
          <div class="flex flex-row text-secondary border-current border-[2px] rounded-md py-2 px-4 gap-x-2 select-none cursor-pointer min-w-52 justify-center items-center mt-10 hover:bg-current">
            <p class="text-primary font-spartan font-semibold text-base">{{ capitalize(menuState.activeMode) }}</p>
          </div>
        </div>
      </div>
    </div>
    <!-- RESULT CONTAINER -->
    <div class="mt-28 w-[100%] flex flex-col justify-center items-center">
      <p class="text-3xl font-poppins text-primary mb-16 font-medium">Result</p>
      <!-- RESULT CONTENT CONTAINER -->
      <div class="flex flex-row justify-center items-center bg-secondary rounded-xl" :class="[menuState.activeType !== TYPE.EXTENDED_VIGENERE ? 'w-[100%] p-10' : 'p-4']">
        <!-- TEXT RESULT -->
        <template v-if="menuState.activeType !== TYPE.EXTENDED_VIGENERE">
          <div class="flex flex-col justify-center items-start w-[100%]">
            <div
              v-for="resultLabel in resultTypes"
              :key="resultLabel"
              class="flex flex-col w-[100%] justify-center items-start"
            >
              <div class="flex flex-row justify-center items-center gap-x-8 mb-4">
                <p class="font-poppins font-semibold text-xl text-primary">{{ capitalize(resultLabel) }}</p>
                <div class="text-secondary border-[2px] border-current select-none cursor-pointer hover:bg-current rounded-md px-3 py-2">
                  <p class="font-spartan text-primary text-sm">Download</p>
                </div>
              </div>
              <p class="font-spartan text-alternate text-lg">Lorem ipsum dolor sit amet consectetur adipisicing elit. Facilis recusandae praesentium dolor odit tempora doloremque iusto vitae ratione deleniti iure nemo sint aliquam repudiandae excepturi aut tenetur, nesciunt quod neque.</p>
              <div v-if="resultLabel !== resultTypes.GROUPED_5_LETTERS" class="border w-[100%] my-8"></div>
            </div>
          </div>
        </template>
        <!-- FILE RESULT -->
        <template v-else>
          <div class="flex flex-col justify-center items-center gap-6">
            <p class="font-poppins font-semibold text-xl text-primary">File</p>
            <img class="w-16 h-16" src="../assets/file.svg">
            <div class="text-secondary border-[2px] border-current select-none cursor-pointer hover:bg-current rounded-md px-3 py-2">
              <p class="font-spartan text-primary text-sm">Download</p>
            </div>
          </div>
        </template>
      </div>
    </div>
    <template v-if="result">
    </template>
  </div>
</template>

<script setup lang="ts">
import { ComputedRef, Ref } from '@vue/reactivity'
import { MODE, TYPE, RESULT_STRING_TYPE, INPUT } from '~/components/constant'
// TYPE
type MenuType = {
  activeMode: string,
  activeType: string
}

type Result = {
  noSpace: string,
  grouped: string
}

const menuState: MenuType = reactive({
  activeMode: MODE.ENCRYPT,
  activeType: TYPE.STANDARD_VIGENERE
})

const types = TYPE
const menus = MODE
const resultTypes = RESULT_STRING_TYPE

const text: Ref<string|null> = ref(null)
const key: Ref<string|null> = ref(null)
let fileRef: Ref<string|any> = ref(null)
let fileKeyRef: Ref<string|any> = ref(null)
let result: Ref<Result|null> = ref(null)

/// Handlers
const handleTypeChange: Function = (type: string): void => { menuState.activeType = type }
const handleModeChange: Function = (mode: string): void => { menuState.activeMode = mode }
const handleInputFile: Function = (e: any, type: INPUT): void => {
  fileRef.value = e.target.files[0]
}
const handleClickInputFile: Function = (id: [1, 2]): void => {
  document.getElementById('hiddenInputFile'+id)?.click()
}
// Computed
const title: ComputedRef<string> = computed(() => `${menuState.activeType} ${menuState.activeMode === MODE.ENCRYPT ? 'ENCRYPTION' : 'DECRYPTION'}`)
const plainPlaceholder: ComputedRef<string> = computed(() => `Input ${menuState.activeMode === MODE.ENCRYPT ? 'plain' : 'cipher'} text here`)
const filePlaceholder: ComputedRef<string> = computed(() => `Upload ${menuState.activeType === TYPE.EXTENDED_VIGENERE ? 'any file' : 'a .txt file'}`)

const capitalize = (str: string) => {
  return str.toLowerCase().split(" ").map(el => el.charAt(0).toUpperCase() + el.slice(1)).join(" ")
}
</script>


