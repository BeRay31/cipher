import axios from 'axios'

export const vigenereEncryptRequest = async(plain: string, key: string, isFile: boolean, style?: string): Promise<{content: string; key: string}> => {
  const requestBody = {
    input: {
      type: isFile ? 'FILE' : 'TEXT',
      content: plain.toUpperCase(),
    },
    key: key.toUpperCase(),
  }
  const response = await axios.post(`http://localhost:1337/vigenere/${style || 'standard'}/encrypt`, requestBody)
  const result = response.data

  return result
}

export const vigenereDecryptRequest = async(plain: string, key: string, isFile: boolean, style: string): Promise<{content: string; key: string}> => {
  const requestBody = {
    input: {
      type: isFile ? 'FILE' : 'TEXT',
      content: plain.toUpperCase(),
    },
    key: key.toUpperCase(),
  }
  const response = await axios.post(`http://localhost:1337/vigenere/${style || 'standard'}/decrypt`, requestBody)
  const result = response.data

  return result
}

export const vigenereExtendedEncryptRequest = async(plain: string, key: string): Promise<{content: string; key: string}> => {
  const requestBody = {
    content: plain,
    key: key.toUpperCase(),
  }
  const response = await axios.post('http://localhost:1337/vigenere/extended/encrypt', requestBody)
  const result = response.data

  return result
}

export const vigenereExtendedDecryptRequest = async(plain: string, key: string): Promise<{content: string; key: string}> => {
  const requestBody = {
    content: plain,
    key: key.toUpperCase(),
  }
  const response = await axios.post('http://localhost:1337/vigenere/extended/decrypt', requestBody)
  const result = response.data

  return result
}
