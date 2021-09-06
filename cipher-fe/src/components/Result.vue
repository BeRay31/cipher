<template>
  <!-- RESULT CONTENT CONTAINER -->
  <div class="flex flex-col bg-gray-700 rounded-xl p-6">
    <div class="flex justify-between items-end mb-8">
      <div class="flex space-x-4 font-semibold">
        <div class="text-3xl">
          Result
        </div>
        <button
          class="btn"
          @click="handleDownload"
        >
          Download
        </button>
      </div>
      <div
        v-if="mainStore.mode !== 'vigenereext' && !mainStore.error"
        class="flex space-x-2"
      >
        <div
          v-for="(displayMode, i) in displayModes"
          :key="`${displayMode}-${i}`"
          class="cursor-pointer"
          :class="[currentDisplayMode === displayMode ? 'border-b-2 border-indigo-400': 'opacity-50']"
          @click="() => setDisplayMode(i)"
        >
          {{ displayMode }}
        </div>
      </div>
    </div>
    <!-- TEXT RESULT -->
    <pre v-else-if="mainStore.mode !== 'vigenereext'" class="w-full overflow-x-auto p-2 rounded-lg bg-gray-600">{{ display(mainStore.resultString) }}</pre>
    <pre v-else-if="!!mainStore.resultFile" class="w-full overflow-x-auto p-2 rounded-lg bg-gray-600">{{ `[${mainStore.isEncrypt ? 'Encrypted' : 'Decrypted'}]${mainStore.fileInputProperties?.name}` }}</pre>
  </div>
</template>

<script setup lang="ts">
import { displayModes } from '~/constants'
import { useMainStore } from '~/store'

const mainStore = useMainStore()

// REFS
const currentDisplayMode = ref(displayModes[0])

// HANDLERS
const setDisplayMode = (displayModeIndex: number) => {
  currentDisplayMode.value = displayModes[displayModeIndex]
}
const downloadFile = (text: string, name: string) => {
  const element = document.createElement('a')
  element.setAttribute('href', `data:text/plain;charset=utf-8,${encodeURIComponent(text || '')}`)
  element.setAttribute('download', name)
  element.click()
  element.remove()
}

const handleDownload = (): void => {
  if (mainStore.mode !== 'vigenereext') {
    downloadFile(mainStore.resultString, new Date().getTime().toString())
  }
  else {
    const link = document.createElement('a')
    const fileName = `[${new Date().getTime().toString()}]${mainStore.fileInputProperties?.name}`
    link.href = window.URL.createObjectURL(mainStore.resultFile)
    link.download = fileName
    link.click()
  }
}

// UTILS
const display = (text: string): string => {
  if (currentDisplayMode.value === '5charblock') {
    let result = ''
    for (let i = 0; i < text.length; i++)
      result += (i !== 0 && i % 5 === 0) ? `\n${text[i]}` : text[i]

    return result
  }
  return text
}
</script>
