<template>
  <div v-if="!fileInput" class="flex flex-col justify-center items-center">
    <input
      type="text"
      :value="inputText"
      @input="changeValue"
      :placeholder="textPlaceholder"
      class="px-2 py-2 rounded min-w-[35rem] font-spartan font-normal focus:outline-none">

      <p class="text-alternate font-spartan text-xl my-10">or</p>
      <!-- INPUT-FILE-CONTAINER -->
      <div
        class="flex flex-row text-secondary border-current border-[2px] rounded-md py-2 px-4 gap-x-2 select-none cursor-pointer hover:bg-current min-w-52 justify-center items-center"
        @click="handleClickInputFile"
      >
        <img src="../assets/foundation_upload-cloud.svg">
        <p class="text-primary font-spartan font-normal text-base">{{ props.filePlaceholder }}</p>
      </div>
      <input id="hiddenInputFile" type="file" class="hidden" :accept="`${acceptAll ? '*' : '.txt'}`" @change="handleFileChange">
  </div>
  <div v-else>
    <div class="flex flex-col justify-center items-center gap-6">
      <p class="font-poppins font-semibold text-xl text-primary">{{ fileInput?.name }}</p>
      <img class="w-16 h-16" src="../assets/file.svg">
      <div @click="removeFile" class="text-secondary border-[2px] border-current select-none cursor-pointer hover:bg-current rounded-md px-3 py-2">
        <p class="font-spartan text-primary text-sm">Remove</p>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { Ref } from "@vue/reactivity"

const props = defineProps({
  content: {
    type: String,
    default: ''
  },
  fileInput: {
    type: Object,
    default: null
  },
  filePlaceholder: {
    type: String,
    default: ''
  },
  textPlaceholder: {
    type: String,
    default: ''
  },
  acceptAll: {
    type: Boolean,
    default: false
  }
})

const inputText: Ref<string> = ref(props.content)

const emit = defineEmits<{
  (event: 'fileChanged', files: Object): void
  (event: 'update', input: string): void
  (event: 'removeFile'): void
}>()

const changeValue = (e: any): void => {
  inputText.value = e.target.value
  emit('update', inputText.value)
}

const removeFile = () => {
  emit('removeFile')
}

const handleClickInputFile = () => {
  document.getElementById('hiddenInputFile')?.click()
}

const handleFileChange = (e: any) => {
  emit('fileChanged', e.target.files[0])
}
</script>