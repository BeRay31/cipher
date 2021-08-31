<script setup lang="ts">
import { useClipboard } from '@vueuse/core'

const ciphertext = ref('')
onMounted(async() => {
  // TODO: fetch source
  const data = 'thisisthecipthertext'
  ciphertext.value = data
})
const { copy } = useClipboard()
const download = (filename: string) => {
  const element = document.createElement('a')
  element.setAttribute('href', `data:text/plain;charset=utf-8,${encodeURIComponent(ciphertext.value)}`)
  element.setAttribute('download', filename)
  element.click()
  element.remove()
}
</script>

<template>
  <div>
    <div class="mb-2">
      <div class="text-4xl font-bold">
        Affine Cipher
      </div>
      <div class="text-sm opacity-75 italic">
        Here's your result.
      </div>
    </div>
    <div class="font-mono w-64 mx-auto mb-2 text-lg">
      {{ ciphertext }}
    </div>
    <div class="flex space-x-2 justify-center mx-auto">
      <button class="btn inline-flex items-center" @click="download(`affine-${new Date().toISOString()}`)">
        <ant-design:save-outlined class="inline mr-2" />
        Save to file
      </button>
      <button class="btn inline-flex items-center" @click="copy(ciphertext)">
        <ant-design:copy-outlined class="inline mr-2" />
        Copy to Clipboard
      </button>
    </div>
  </div>
</template>
