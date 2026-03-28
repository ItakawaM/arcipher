package cmd

import (
	"github.com/ItakawaM/go-cryptotool/ciphers"
	"github.com/spf13/cobra"
)

func NewVigenereCommand() *cobra.Command {
	vigenereCmd := &cobra.Command{
		Use:   "vigenere",
		Short: "Encrypt or decrypt data using the Vigenère cipher",
		Long: `The Vigenère cipher is a classical polyalphabetic substitution cipher
that uses a keyword to apply multiple Caesar shifts across the plaintext.

Each letter in the keyword determines the shift for the corresponding
letter in the message. The keyword is repeated as needed to match the
length of the plaintext.

For example, using the keyword "KEY":
- K shifts the first letter
- E shifts the second
- Y shifts the third
- then the pattern repeats

Unlike the Caesar cipher, this approach makes frequency analysis
more difficult by distributing shifts across multiple alphabets.

This command allows encryption and decryption of messages or files
using a specified keyword.
`,
	}
	vigenereCmd.AddCommand(
		newVigenereEncryptCommand(),
		newVigenereDecryptCommand(),
	)

	return vigenereCmd
}

func newVigenereEncryptCommand() *cobra.Command {
	params := &vigenereParams{}

	encryptCmd := &cobra.Command{
		Use:   "encrypt <keyword> <message | input> [output]",
		Short: "Encrypt a given message/file with a keyword",
		Args:  cobra.RangeArgs(2, 3),
		Long: `This command allows encryption of messages or files
using the Vigenère cipher with a specified keyword.

Non-alphabetic characters remain unchanged and do not
consume characters from the keyword.

Examples:

  Encrypt text:
    1. go-cryptotool vigenere encrypt KEY "AttackAtDawn"

  Encrypt a file:
    1. go-cryptotool vigenere encrypt SECRET file.txt file.enc

Notes:

  • The keyword must consist of alphabetic characters only [a-zA-Z]
  • Letter shifts are derived from keyword characters (A=0, B=1, ..., Z=25)
  • The keyword is case-insensitive (but case can be preserved in output)
  • For very large files, performance depends on CPU and SSD
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return vigenereRunE(cmd, args, params, ciphers.Encrypt)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return vigenerePreRunE(cmd, params, args)
		},
	}
	params.addFlags(encryptCmd)

	return encryptCmd
}

func newVigenereDecryptCommand() *cobra.Command {
	params := &vigenereParams{}

	decryptCmd := &cobra.Command{
		Use:   "decrypt <keyword> <message | input> [output]",
		Short: "Decrypt a given message/file with a keyword",
		Args:  cobra.RangeArgs(2, 3),
		Long: `This command allows decryption of messages or files
using the Vigenère cipher with a specified keyword.

Non-alphabetic characters remain unchanged and do not
consume characters from the keyword.

Examples:

  Decrypt text:
    1. go-cryptotool vigenere decrypt KEY "KxrkgiKxBkal"

  Decrypt a file:
    1. go-cryptotool vigenere decrypt SECRET file.enc file.txt

Notes:

  • The keyword must consist of alphabetic characters only [a-zA-Z]
  • Letter shifts are derived from keyword characters (A=0, B=1, ..., Z=25)
  • The keyword is case-insensitive (but case can be preserved in output)
  • For very large files, performance depends on CPU and SSD
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return vigenereRunE(cmd, args, params, ciphers.Decrypt)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return vigenerePreRunE(cmd, params, args)
		},
	}
	params.addFlags(decryptCmd)

	return decryptCmd
}
