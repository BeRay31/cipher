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
          <InputField
            :content="text"
            :file-input="textFileRef"
            :text-placeholder="plainPlaceholder"
            :file-placeholder="filePlaceholder"
            :accept-all="menuState.activeType === TYPE.EXTENDED_VIGENERE"
            @update="(e) => {text = e}"
            @file-changed="(e) => {textFileRef = e}"
            @remove-file="() => {textFileRef = null}"
          />
          <div class="my-4"></div>
          <InputField 
            :content="key"
            :file-input="keyFileRef"
            :text-placeholder="`Input key here`"
            :file-placeholder="filePlaceholder"
            :accept-all="menuState.activeType === TYPE.EXTENDED_VIGENERE"
            @update="(e) => {key = e}"
            @file-changed="(e) => {keyFileRef = e}"
            @remove-file="() => {keyFileRef = null}"
          />
          <div @click="handleStartButton" class="flex flex-row text-secondary border-current border-[2px] rounded-md py-2 px-4 gap-x-2 select-none cursor-pointer w-[90%] justify-center items-center mt-10 hover:bg-current">
            <p class="text-primary font-spartan font-semibold text-base">{{ capitalize(menuState.activeMode) }}</p>
          </div>
        </div>
      </div>
    </div>
    <template v-if="result">
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
                  <div @click="() => handleDownload(resultLabel == resultTypes.NO_SPACE ? outputTypes.NO_SPACE : outputTypes.GROUPED)" class="text-secondary border-[2px] border-current select-none cursor-pointer hover:bg-current rounded-md px-3 py-2">
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
              <div @click="() => handleDownload(outputTypes.FILE)" class="text-secondary border-[2px] border-current select-none cursor-pointer hover:bg-current rounded-md px-3 py-2">
                <p class="font-spartan text-primary text-sm">Download</p>
              </div>
            </div>
          </template>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import { ComputedRef, Ref } from '@vue/reactivity'
import Hill from '~/api/Hill'
import VigenereAuto from '~/api/VigenereAuto'
import VigenereFull from '~/api/VigenereFull'
import VigenereStandard from '~/api/VigenereStandard'
import { MODE, TYPE, RESULT_STRING_TYPE, RESULT } from '~/components/constant'
import InputField from '~/components/InputField.vue'
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
const outputTypes = RESULT

const text: Ref<string> = ref('')
const key: Ref<string> = ref('')
let textFileRef: Ref<string|any> = ref(null)
let keyFileRef: Ref<string|any> = ref(null)

let result: Ref<Result|null> = ref(null)

/// Handlers
const handleTypeChange: Function = (type: string): void => { menuState.activeType = type }
const handleModeChange: Function = (mode: string): void => { menuState.activeMode = mode }
const handleDownload: Function = (type: RESULT): void => {
  switch (type) {
    case RESULT.NO_SPACE:
      downloadFile(result.value?.noSpace, title.value)
      console.log("🚀 ~ file: index.vue ~ line 144 ~ RESULT.NO_SPACE", RESULT.NO_SPACE)
      break;
    case RESULT.GROUPED:
      downloadFile(result.value?.grouped, title.value)
      console.log("🚀 ~ file: index.vue ~ line 144 ~ RESULT.GROUPED", RESULT.GROUPED)
      break;
    case RESULT.FILE:
      console.log("🚀 ~ file: index.vue ~ line 144 ~ RESULT.FILE", RESULT.FILE)
      break;
  
    default:
      break;
  }
}

const handleStartButton = async () => {
  if ((!key.value && !text.value) && (!keyFileRef.value && !textFileRef.value)) {
    // TODO Made the UI
    alert("Please fill the form correctly")
  }
  let APIClass: any = VigenereStandard
  // TODO handle file
  switch (menuState.activeType) {
    case types.STANDARD_VIGENERE:
      APIClass = VigenereStandard
      break;
    case types.FULL_VIGENERE:
      APIClass = VigenereFull
      break;
    case types.EXTENDED_VIGENERE:
      console.log("TODO", types.EXTENDED_VIGENERE)
      break;
    case types.AUTO_KEY_VIGENERE:
      APIClass = VigenereAuto
      break;
    case types.AFFINE:
      console.log("TODO", types.AFFINE)
      break;
    case types.PLAYFAIR:
      console.log("TODO", types.PLAYFAIR)
      break;
    case types.HILL:
      APIClass = Hill
      break;
    default:
      APIClass = VigenereStandard
      break;
  }
  // TODO Make sure API call work
  let resp: any
  if (menuState.activeMode === menus.ENCRYPT) {
    resp = await APIClass.encrypt(text.value)
  } else {
    resp = await APIClass.encrypt(text.value)
  }

  const content: string = String(resp.content)
  let groupedContent: string = ""
  let startIdx: number = 0
  let endIdx: number = 5
  if (endIdx >= content.length) {
    groupedContent = content
  } else {
    while (endIdx < content.length) {
      groupedContent += content.substring(startIdx, endIdx)
      startIdx = endIdx
      endIdx+=5
      if (endIdx >= content.length) {
        groupedContent += content.substring(startIdx, endIdx)
      }
    }
  }
  result.value = {
    noSpace: content.replaceAll(" ", ""),
    grouped: groupedContent
  }
}

// Computed
const title: ComputedRef<string> = computed(() => `${menuState.activeType} ${menuState.activeMode === MODE.ENCRYPT ? 'ENCRYPTION' : 'DECRYPTION'}`)
const plainPlaceholder: ComputedRef<string> = computed(() => `Input ${menuState.activeMode === MODE.ENCRYPT ? 'plain' : 'cipher'} text here`)
const filePlaceholder: ComputedRef<string> = computed(() => `Upload ${menuState.activeType === TYPE.EXTENDED_VIGENERE ? 'any file' : 'a .txt file'}`)

const capitalize = (str: string) => {
  return str.toLowerCase().split(" ").map(el => el.charAt(0).toUpperCase() + el.slice(1)).join(" ")
}

const downloadFile = (text: string | undefined, name: string) => {
  const element = document.createElement('a')
  element.setAttribute('href', `data:text/plain;charset=utf-8,${encodeURIComponent(text || '')}`)
  element.setAttribute('download', name)
  element.click()
  element.remove()
}
</script>


