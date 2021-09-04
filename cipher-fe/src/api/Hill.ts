import service from "./service";

export default class VigenereStandard {
  static baseURL = `hill`

  static async Decrypt(cipher: string, key: string) {
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

  static async Encrypt(plain: string, key: string) {
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