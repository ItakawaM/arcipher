package ciphers

const KB = 1024

type Mode int

const (
	Encrypt Mode = iota
	Decrypt
)

func ModeToString(mode Mode) string {
	if mode == Encrypt {
		return "encrypt"
	}
	return "decrypt"
}

type BlockCipher interface {
	EncryptBlock([]byte, []byte) error
	DecryptBlock([]byte, []byte) error
}
