<template>
  <template v-if="!fileText" class="flex justify-center items-center space-x-4">
    <div class="flex justify-center items-center space-x-4">
      <input
        v-if="!acceptAll"
        type="text"
        :value="inputText"
        :placeholder="placeholder"
        class="px-2 py-2 rounded w-full bg-gray-600 text-gray-100  font-normal focus:outline-none"
        @change="changeValue"
      >
      <div v-if="!acceptAll">
        or
      </div>
      <button
        class="btn inline-flex items-center space-x-2"
        @click="handleClickInputFile"
      >
        <ant-design:cloud-upload-outlined />
        <div>
          Upload File
        </div>
      </button>
      <input id="hiddenInputFile" type="file" class="hidden" :accept="`${acceptAll ? '*' : '.txt'}`" @change="handleFileChange">
    </div>
  </template>
  <template v-else>
    <div class="flex items-center space-x-6">
      <div>
        Text:
      </div>
      <div class="opacity-25">
        {{ fileText }}
      </div>
      <button class="btn inline-flex items-center space-x-2" @click="handleRemoveFile">
        <ant-design:caret-left-outlined />
        Remove
      </button>
    </div>
  </template>
</template>
<script lang="ts" setup>
// PROPS
const props = defineProps({
  content: {
    type: String,
    default: '',
  },
  placeholder: {
    type: String,
    default: '',
  },
  acceptAll: {
    type: Boolean,
    default: false,
  },
})

// REFS
const inputText = ref(props.content)
const fileText = ref('')

// EMITS
const emit = defineEmits(['fileChanged', 'update', 'removeFile'])

const changeValue = (e: any): void => {
  const text = e.target.value
  inputText.value = text
  emit('update', text)
}

const handleRemoveFile = () => {
  emit('removeFile')
  fileText.value = ''
}

const handleClickInputFile = () => {
  document.getElementById('hiddenInputFile')?.click()
}

const handleFileChange = (e: any) => {
  // if not vigenere extended
  if (!props.acceptAll) {
    const reader = new FileReader()
    reader.addEventListener('load', (event) => {
      const text = reader.result as string
      fileText.value = text
      emit('fileChanged', text)
    })
    reader.readAsText(e.target.files[0])
  }
  else {
    const reader = new FileReader()
    reader.addEventListener('load', () => {
      const result = reader.result as string
      fileText.value = '<file read>'
      emit('fileChanged', result)
    })
    reader.readAsDataURL(e.target.files[0])
  }
}
</script>
