package cmd

import (
	"github.com/ItakawaM/go-cryptotool/ciphers"
	"github.com/spf13/cobra"
)

func NewRailFenceCommand() *cobra.Command {
	railfenceCmd := &cobra.Command{
		Use:   "railfence",
		Short: "Encrypt or decrypt data using the Rail Fence cipher",
		Long: `The Rail Fence cipher is a classical transposition cipher that writes
plaintext in a zigzag pattern across multiple rails and then reads
it row by row to produce the ciphertext.

This command allows encryption and decryption of messages or files
using a specified number of rails (key).
`,
	}
	railfenceCmd.AddCommand(
		newRailFenceEncryptCommand(),
		newRailFenceDecryptCommand(),
	)

	return railfenceCmd
}

func newRailFenceEncryptCommand() *cobra.Command {
	params := &railfenceParams{}

	encryptCmd := &cobra.Command{
		Use:   "encrypt <key> <message | input> [output]",
		Short: "Encrypt a given message/file with a key",
		Args:  cobra.RangeArgs(2, 3),
		Long: `This command allows encryption of messages or files
using a specified number of rails (key).

A key of 1 results in no transformation.

Examples:

  Encrypt text:
    1. go-cryptotool railfence encrypt 3 "Canabis"
  
  Encrypt a file:
    1. go-cryptotool railfence encrypt 5 file.txt file.enc

Notes:

  • The key must be >= 1
  • Larger keys increase computation time
  • For very large files, performance depends on system memory
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return railfenceRunE(cmd, args, params, ciphers.Encrypt)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return railfencePreRunE(cmd, params, args)
		},
	}
	params.addFlags(encryptCmd)

	return encryptCmd
}

func newRailFenceDecryptCommand() *cobra.Command {
	params := &railfenceParams{}

	decryptCmd := &cobra.Command{
		Use:   "decrypt <key> <message | input> [output]",
		Short: "Decrypt a given message/file with a key",
		Args:  cobra.RangeArgs(2, 3),
		Long: `This command allows decryption of messages or files
using a specified number of rails (key).

Examples:

  Decrypt text:
    1. go-cryptotool railfence decrypt 3 "nsaaiCb"
  
  Decrypt a file:
    1. go-cryptotool railfence decrypt 5 file.enc file.txt

Notes:

  • The key must be >= 1
  • Larger keys increase computation time
  • For very large files, performance depends on system memory
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return railfenceRunE(cmd, args, params, ciphers.Decrypt)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return railfencePreRunE(cmd, params, args)
		},
	}
	params.addFlags(decryptCmd)

	return decryptCmd
}
