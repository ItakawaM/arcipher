package ciphers

type Mode int

const (
	Encrypt Mode = iota
	Decrypt
)

func (mode Mode) String() string {
	if mode == Encrypt {
		return "encrypt"
	}
	return "decrypt"
}

type BlockCipher interface {
	IsInPlace() bool
	EncryptBlock(dst []byte, src []byte) error
	DecryptBlock(dst []byte, src []byte) error
}
