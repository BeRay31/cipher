import { defineStore } from 'pinia'
import { cipherTypes } from './constants'
import { Ref } from '@vue/reactivity'
export const useMainStore = defineStore('main', () => {
  const modeIndex = ref(0)
  const mode = ref(cipherTypes[0])
  const isEncrypt = ref(true)
  const inputString = ref('')
  const keyString = ref('')
  const inputFileString = ref('')
  const fileInputProperties: Ref<File | null> = ref(null)
  const resultFile: Ref<File | null> = ref(null)
  const keyFileString = ref('')
  const resultString = ref('')

  return {
    mode,
    modeIndex,
    isEncrypt,
    inputString,
    inputFileString,
    fileInputProperties,
    resultFile,
    keyString,
    keyFileString,
    resultString,
  }
})
// export const useMainStore = defineStore('main', {
//   state: () => ({
//     mode: cipherTypes[0],
//     modeIndex: 0,
//     inputString: '',
//     keyString: '',
//     inputFileString: '',
//     keyFileString: '',
//   }),
//   getters: {
//     doEncrypt() {

//     },
//     doDecrypt() {

//     },
//   },
//   actions: {
//     setMode(i: number) {
//       this.modeIndex = i
//       this.mode = cipherTypes[i]
//     },
//     setInputString(s: string) {
//       this.inputString = s
//     },
//     setKeyString(s: string) {
//       this.keyString = s
//     },
//   },
// })
