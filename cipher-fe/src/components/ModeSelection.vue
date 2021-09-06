<template>
  <div class="bg-gray-700 rounded-2xl flex flex-col p-6">
    <div class="inline-flex space-x-2 text-gray-400 mb-2">
      <div
        class="cursor-pointer"
        :class="[mainStore.isEncrypt ? 'border-b-2 border-indigo-400': 'opacity-50']"
        @click="() => setIsEncrypt(true)"
      >
        Encrypt
      </div>
      <div
        class="cursor-pointer"
        :class="[!mainStore.isEncrypt ? 'border-b-2 border-indigo-400': 'opacity-50']"
        @click="() => setIsEncrypt(false)"
      >
        Decrypt
      </div>
    </div>
    <div class="flex flex-col space-y-2 mb-2">
      <InputField
        :content="mainStore.inputString"
        :placeholder="mainStore.isEncrypt ? 'Plaintext...' :'Ciphertext...'"
        :accept-all="mainStore.mode === 'vigenereext'"
        type="content"
        @update="handleInputUpdate"
        @file-changed="handleInputFileChange"
        @remove-file="handleInputFileRemove"
      />
      <InputField
        :content="mainStore.keyString"
        placeholder="Key..."
        :accept-all="false"
        type="key"
        @update="handleKeyUpdate"
        @file-changed="handleKeyFileChange"
        @remove-file="handleKeyFileRemove"
      />
    </div>
    <div v-if="mainStore.mode === 'affine'" class="opacity-20 mb-2">
      Note: For affine ecnryption, key should be in format of m,b (key, offset)
    </div>
    <div v-if="mainStore.mode === 'hill'" class="opacity-20 mb-2">
      Note: For hill ecnryption, key should be have a quadratic length ex "bcef"
    </div>
    <button class="btn w-full py-2" @click="handleButton">
      {{ mainStore.isEncrypt ? 'Encrypt' :'Decrypt' }}
    </button>
  </div>
</template>

<script setup lang="ts">
import { useMainStore } from '~/store'
import { vigenereDecryptRequest, vigenereEncryptRequest, vigenereExtendedDecryptRequest, vigenereExtendedEncryptRequest } from '~/api/vigenere'
import { affineDecryptRequest, affineEncryptRequest } from '~/api/affine'
import { playfairDecryptRequest, playfairEncryptRequest } from '~/api/playfair'
import { hillDecryptRequest, hillEncryptRequest } from '~/api/hill'

const mainStore = useMainStore()

const setIsEncrypt = (status: boolean) => {
  mainStore.isEncrypt = status
  mainStore.resetIO()
}

// HANDLERS
const handleInputUpdate = (text: string) => {
  mainStore.inputString = text
}
const handleInputFileChange = (text: string) => {
  mainStore.inputFileString = text
}
const handleInputFileRemove = () => {
  mainStore.inputFileString = ''
}

const handleKeyUpdate = (text: string) => {
  mainStore.keyString = text
}
const handleKeyFileChange = (text: string) => {
  mainStore.keyFileString = text
}
const handleKeyFileRemove = () => {
  mainStore.keyFileString = ''
}

const handleButton = async() => {
  // get required data
  const mode = mainStore.mode
  const isInputFile = mainStore.inputFileString !== ''
  const input = isInputFile ? mainStore.inputFileString : mainStore.inputString
  const isKeyFile = mainStore.keyFileString !== ''
  const key = isKeyFile ? mainStore.keyFileString : mainStore.keyString
  try {
    // check mode
    if (mainStore.isEncrypt) {
      if (mode === 'vigenerestd') {
        const result = await vigenereEncryptRequest(input, key, false, 'standard')
        mainStore.resultString = result.content
      }
      if (mode === 'vigenerefull') {
        const result = await vigenereEncryptRequest(input, key, false, 'full')
        mainStore.resultString = result.content
      }
      if (mode === 'vigenereauto') {
        const result = await vigenereEncryptRequest(input, key, false, 'auto')
        mainStore.resultString = result.content
      }
      if (mode === 'vigenereext') {
        const result = await vigenereExtendedEncryptRequest(mainStore.fileInputProperties as File, key)
        const unsigned8byteArray = new Uint8Array(result)
        const file = new Blob([unsigned8byteArray])
        mainStore.resultFile = file as File
      }
      if (mode === 'playfair') {
        const result = await playfairEncryptRequest(input, key, false)
        mainStore.resultString = result.content
      }
      if (mode === 'hill') {
        const result = await hillEncryptRequest(input, key, false)
        mainStore.resultString = result.content
      }
      if (mode === 'affine') {
        const [m, b] = key.split(',')
        const result = await affineEncryptRequest(input, parseInt(m), parseInt(b), false)
        mainStore.resultString = result.content
      }
    }
    else {
      if (mode === 'vigenerestd') {
        const result = await vigenereDecryptRequest(input, key, false, 'standard')
        mainStore.resultString = result.content
      }
      if (mode === 'vigenerefull') {
        const result = await vigenereDecryptRequest(input, key, false, 'full')
        mainStore.resultString = result.content
      }
      if (mode === 'vigenereauto') {
        const result = await vigenereDecryptRequest(input, key, false, 'auto')
        mainStore.resultString = result.content
      }
      if (mode === 'vigenereext') {
        const result = await vigenereExtendedDecryptRequest(mainStore.fileInputProperties as File, key)
        const unsigned8byteArray = new Uint8Array(result)
        const file = new Blob([unsigned8byteArray])
        mainStore.resultFile = file as File
      }
      if (mode === 'playfair') {
        const result = await playfairDecryptRequest(input, key, false)
        mainStore.resultString = result.content
      }
      if (mode === 'hill') {
        const result = await hillDecryptRequest(input, key, false)
        mainStore.resultString = result.content
      }
      if (mode === 'affine') {
        const [m, b] = key.split(',')
        const result = await affineDecryptRequest(input, parseInt(m), parseInt(b), false)
        mainStore.resultString = result.content
      }
    }
  } catch(e) {
    mainStore.resultString = e.toString()
  }
}

</script>
