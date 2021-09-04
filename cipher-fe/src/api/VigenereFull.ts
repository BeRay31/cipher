import service from "./service";

export default class VigenereFull{
  static baseURL = `vigenere/full`

  static async decrypt(cipher: string, key: string) {
    const req = {
      input: {
        type: "TEXT",
        content: cipher
      },
      key: key
    }
    const resp = await service.post(`${this.baseURL}/decrypt`, req)
    return resp;
  }

  static async encrypt(plain: string, key: string) {
    const req = {
      input: {
        type: "TEXT",
        content: plain
      },
      key: key
    }
    const resp = await service.post(`${this.baseURL}/encrypt`, req)
    return resp;
  }
}