export enum MODE {
  ENCRYPT = 'ENCRYPT',
  DECRYPT = 'DECRYPT'
}

export enum TYPE {
  STANDARD_VIGENERE = 'STANDARD VIGENERE',
  FULL_VIGENERE = 'FULL VIGENERE',
  EXTENDED_VIGENERE = 'EXTENDED VIGENERE',
  AUTO_KEY_VIGENERE = 'AUTO KEY VIGENERE',
  AFFINE = 'AFFINE',
  PLAYFAIR = 'PLAYFAIR',
  HILL = 'HILL',
}

export const cipherTypes = ['vigenerestd', 'vigenerefull', 'vigenereext', 'vigenereauto', 'affine', 'playfair', 'hill']
export const displayModes = ['nospace', '5charblock']

export enum RESULT_STRING_TYPE {
  NO_SPACE = 'NO SPACE',
  GROUPED_5_LETTERS = 'GROUPED 5 LETTERS'
}

export enum RESULT {
  NO_SPACE = 'NO SPACE',
  GROUPED = 'GROUPED',
  FILE = 'FILE'
}
