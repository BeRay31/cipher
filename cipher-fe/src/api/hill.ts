import axios from 'axios'

export const hillEncryptRequest = async(plain: string, key: string, isFile: boolean): Promise<{content: string; key: number; offset: number}> => {
  const requestBody = {
    input: {
      type: isFile ? 'FILE' : 'TEXT',
      content: plain.toUpperCase(),
    },
    key: key.toUpperCase(),
  }
  const response = await axios.post('http://localhost:1337/hill/encrypt', requestBody)
  const result = response.data

  return result
}

export const hillDecryptRequest = async(plain: string, key: string, isFile: boolean): Promise<{content: string; key: number; offset: number}> => {
  const requestBody = {
    input: {
      type: isFile ? 'FILE' : 'TEXT',
      content: plain.toUpperCase(),
    },
    key: key.toUpperCase(),
  }
  const response = await axios.post('http://localhost:1337/hill/decrypt', requestBody)
  const result = response.data

  return result
}
