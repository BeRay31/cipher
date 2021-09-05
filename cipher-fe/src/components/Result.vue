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
    <pre v-if="mainStore.mode !== 'vigenereext'" class="w-full overflow-x-auto p-2 rounded-lg bg-gray-600">{{ display(mainStore.resultString) }}</pre>
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
const dataURItoBlob = (dataURI: string) => {
  const buffer = Buffer.from(dataURI, 'base64')
  const ab = Uint8Array.from(buffer).buffer
  // write the ArrayBuffer to a blob, and you're done
  const blob = new Blob([ab])
  return blob
}

const toBinary = (string: string) => {
  const codeUnits = new Uint16Array(string.length)
  for (let i = 0; i < codeUnits.length; i++)
    codeUnits[i] = string.charCodeAt(i)

  return btoa(String.fromCharCode(...new Uint8Array(codeUnits.buffer)))
}

const fromBinary = (encoded: string) => {
  const binary = atob(encoded)
  const bytes = new Uint8Array(binary.length)
  for (let i = 0; i < bytes.length; i++)
    bytes[i] = binary.charCodeAt(i)

  return String.fromCharCode(...new Uint16Array(bytes.buffer))
}

const handleDownload = (): void => {
  if (mainStore.mode !== 'vigenereext') { downloadFile(mainStore.resultString, new Date().getTime().toString()) }
  else {
    console.log(mainStore.resultString)
    const blob = dataURItoBlob(mainStore.resultString)
    const link = document.createElement('a')
    const fileName = new Date().getTime().toString()
    link.href = window.URL.createObjectURL(blob)
    link.download = fileName
    link.click()
    // downloadFile(bytes.buffer, new Date().getTime().toString())
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
