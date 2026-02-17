package ciphers

import (
	"bytes"
	"testing"
)

func TestRailFenceEncrypt(t *testing.T) {
	tests := []struct {
		name      string
		message   string
		encrypted string
		key       int
	}{
		{"Normal 1", "Canabis", "nsaaiCb", 3},
		{"Normal 2", "Hello World!!", "o!l !lWdeolHr", 5},
		{"Normal 3", "Chicken", "hceCikn", 2},
		{"Empty", "", "", 2},
		{"Big Key", "Hello World!", "!dlroW olleH", 123},
		{"Negative Key", "Negative", "Negative", -1},
		{"Key of 1", "Positive", "Positive", 1},
	}

	for _, testSubject := range tests {
		t.Run(testSubject.name, func(t *testing.T) {
			src := []byte(testSubject.message)
			expected := []byte(testSubject.encrypted)

			cipher := NewRailFenceCipher(testSubject.key, len(src), 1)
			cipher.BuildPermutationTable()

			dst := make([]byte, len(src))

			if err := cipher.EncryptBlock(dst, src); err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(dst, expected) {
				t.Fatalf("want: %s\ngot: %s", expected, dst)
			}
		})
	}
}

func TestRailFenceCipher(t *testing.T) {
	tests := []struct {
		name    string
		message string
		key     int
	}{
		{"Normal 1", "Canabis", 3},
		{"Normal 2", "Hello World!", 5},
		{"Empty", "", 2},
		{"Big Key", "Hello World!", 123},
		{"Negative Key", "Negative", -1},
		{"Key of 1", "Positive", 1},
	}

	for _, testSubject := range tests {
		t.Run(testSubject.name, func(t *testing.T) {
			cipher := NewRailFenceCipher(testSubject.key, len(testSubject.message), 1)
			cipher.BuildPermutationTable()

			expected := []byte(testSubject.message)
			src := make([]byte, len(expected))
			copy(src, expected)

			dst := make([]byte, len(src))

			if err := cipher.EncryptBlock(dst, src); err != nil {
				t.Fatal(err)
			}

			if err := cipher.DecryptBlock(src, dst); err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(src, expected) {
				t.Fatalf("want: %s, got: %s", expected, src)
			}
		})
	}
}

func FuzzRailFenceCipher(f *testing.F) {
	f.Add("Canabis", 3)
	f.Add("ABC", 5)
	f.Add("Hello World!", 2)
	f.Add("", 12)

	f.Fuzz(func(t *testing.T, message string, key int) {
		if key <= 1 {
			return
		} else if key > len(message)+10 {
			key = len(message)
		}

		cipher := NewRailFenceCipher(key, len(message), 1)
		cipher.BuildPermutationTable()

		expected := []byte(message)
		src := make([]byte, len(expected))
		copy(src, expected)

		dst := make([]byte, len(src))

		cipher.EncryptBlock(dst, src)
		cipher.DecryptBlock(src, dst)

		if !bytes.Equal(src, expected) {
			t.Fatalf("Encrypt/Decrypt mismatch")
		}
	})
}
