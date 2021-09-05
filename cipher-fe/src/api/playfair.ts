import axios from 'axios'

export const playfairEncryptRequest = async(plain: string, key: string, isFile: boolean): Promise<{content: string; key: number; offset: number}> => {
  const requestBody = {
    input: {
      type: isFile ? 'FILE' : 'TEXT',
      content: plain.toUpperCase(),
    },
    key: key.toUpperCase(),
  }
  const response = await axios.post('http://localhost:1337/playfair/encrypt', requestBody)
  const result = response.data

  return result
}

export const playfairDecryptRequest = async(plain: string, key: string, isFile: boolean): Promise<{content: string; key: number; offset: number}> => {
  const requestBody = {
    input: {
      type: isFile ? 'FILE' : 'TEXT',
      content: plain.toUpperCase(),
    },
    key: key.toUpperCase(),
  }
  const response = await axios.post('http://localhost:1337/playfair/decrypt', requestBody)
  const result = response.data

  return result
}
