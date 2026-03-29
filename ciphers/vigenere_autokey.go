package ciphers

import "fmt"

type VigenereAutoKeyCipher struct {
	VigenereCipher
}

func NewVigenereAutoKeyCipher(key []byte) (*VigenereAutoKeyCipher, error) {
	normalizedKey, err := NormalizeVigenereKey(key)
	if err != nil {
		return nil, err
	}
	return &VigenereAutoKeyCipher{VigenereCipher{Key: normalizedKey}}, nil
}

func NewVigenereAutoKeyCipherNormalized(normalizedKey []byte) *VigenereAutoKeyCipher {
	return &VigenereAutoKeyCipher{VigenereCipher{Key: normalizedKey}}
}

func (vCipher *VigenereAutoKeyCipher) IsInPlace() bool {
	return false
}

func (vCipher *VigenereAutoKeyCipher) EncryptBlock(dst []byte, src []byte) error {
	if len(dst) != len(src) {
		return fmt.Errorf("block size mismatch src=%d dst=%d", len(src), len(dst))
	}

	keyIndex := 0
	keyCycle := len(vCipher.Key)
	letterIndex := 0
	for index, char := range src {
		switch {
		case char >= 'a' && char <= 'z':
			if keyIndex < keyCycle {
				dst[index] = (char-'a'+(vCipher.Key[keyIndex]))%26 + 'a'
			} else {
				for !isASCIILetter(src[letterIndex]) {
					letterIndex += 1
				}
				dst[index] = (char-'a'+getShift(src[letterIndex]))%26 + 'a'
				letterIndex += 1
			}
			keyIndex += 1

		case char >= 'A' && char <= 'Z':
			if keyIndex < keyCycle {
				dst[index] = (char-'A'+(vCipher.Key[keyIndex]))%26 + 'A'
			} else {
				for !isASCIILetter(src[letterIndex]) {
					letterIndex += 1
				}
				dst[index] = (char-'A'+getShift(src[letterIndex]))%26 + 'A'
				letterIndex += 1
			}
			keyIndex += 1

		default:
			dst[index] = char
		}
	}

	return nil
}

func (vCipher *VigenereAutoKeyCipher) DecryptBlock(dst []byte, src []byte) error {
	if len(dst) != len(src) {
		return fmt.Errorf("block size mismatch src=%d dst=%d", len(src), len(dst))
	}

	keyIndex := 0
	keyCycle := len(vCipher.Key)
	letterIndex := 0
	for index, char := range src {
		switch {
		case char >= 'a' && char <= 'z':
			if keyIndex < keyCycle {
				dst[index] = (char-'a'-(vCipher.Key[keyIndex])+26)%26 + 'a'
			} else {
				for !isASCIILetter(src[letterIndex]) {
					letterIndex += 1
				}
				dst[index] = (char-'a'-getShift(dst[letterIndex])+26)%26 + 'a'
				letterIndex += 1
			}
			keyIndex += 1

		case char >= 'A' && char <= 'Z':
			if keyIndex < keyCycle {
				dst[index] = (char-'A'-(vCipher.Key[keyIndex])+26)%26 + 'A'
			} else {
				for !isASCIILetter(src[letterIndex]) {
					letterIndex += 1
				}
				dst[index] = (char-'A'-getShift(dst[letterIndex])+26)%26 + 'A'
				letterIndex += 1
			}
			keyIndex += 1

		default:
			dst[index] = char
		}
	}

	return nil
}
