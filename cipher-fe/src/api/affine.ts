import axios from 'axios'

export const affineEncryptRequest = async(plain: string, key: number, offset: number, isFile: boolean): Promise<{content: string; key: number; offset: number}> => {
  const requestBody = {
    input: {
      type: isFile ? 'FILE' : 'TEXT',
      content: plain.toUpperCase(),
    },
    key,
    offset,
  }
  const response = await axios.post('http://localhost:1337/affine/encrypt', requestBody)
  const result = response.data

  return result
}

export const affineDecryptRequest = async(plain: string, key: number, offset: number, isFile: boolean): Promise<{content: string; key: number; offset: number}> => {
  const requestBody = {
    input: {
      type: isFile ? 'FILE' : 'TEXT',
      content: plain.toUpperCase(),
    },
    key,
    offset,
  }
  const response = await axios.post('http://localhost:1337/affine/decrypt', requestBody)
  const result = response.data

  return result
}
