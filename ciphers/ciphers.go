package ciphers

const DefaultBlockSize int64 = 4 * 1024 * 1024

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
	EncryptBlock([]byte) error
	DecryptBlock([]byte) error
}
