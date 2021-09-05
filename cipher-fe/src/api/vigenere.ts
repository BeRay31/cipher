import axios from 'axios'

export const vigenereEncryptRequest = async(plain: string, key: string, isFile: boolean): Promise<{content: string; key: string}> => {
  const requestBody = {
    input: {
      type: isFile ? 'FILE' : 'TEXT',
      content: plain.toUpperCase(),
    },
    key: key.toUpperCase(),
  }
  const response = await axios.post('http://localhost:1337/vigenere/standard/encrypt', requestBody)
  const result = response.data

  return result
}

export const vigenereDecryptRequest = async(plain: string, key: string, isFile: boolean): Promise<{content: string; key: string}> => {
  const requestBody = {
    input: {
      type: isFile ? 'FILE' : 'TEXT',
      content: plain.toUpperCase(),
    },
    key: key.toUpperCase(),
  }
  const response = await axios.post('http://localhost:1337/vigenere/standard/decrypt', requestBody)
  const result = response.data

  return result
}
