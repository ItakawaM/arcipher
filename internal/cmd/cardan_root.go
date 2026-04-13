package cmd

import (
	"runtime"

	"github.com/ItakawaM/arcipher/ciphers"
	"github.com/spf13/cobra"
)

func NewCardanCommand() *cobra.Command {
	cardanCmd := &cobra.Command{
		Use:   "cardan",
		Short: "Encrypt or decrypt data using the Cardan grille cipher",
		Long: `The Cardan grille is a classical permutation cipher that uses
a stencil (grille) with holes cut into it. The grille is placed
over a grid, and the text is written through the holes.

After writing the letters, the grille is rotated (usually by
90 degrees) several times, filling the exposed positions each
time until the grid is complete.
`,
	}
	cardanCmd.AddCommand(
		newCardanEncryptCommand(),
		newCardanDecryptCommand(),
		newCardanGenerateKeyCommand(),
	)

	return cardanCmd
}

func newCardanEncryptCommand() *cobra.Command {
	params := &cardanParams{}

	encryptCmd := &cobra.Command{
		Use:   "encrypt [key] <message | input> [output]",
		Short: "Encrypt a message or file using a Cardan grille cipher",
		Args:  cobra.RangeArgs(1, 3),
		Long: `Encrypt messages or files using the Cardan grille cipher.

If no key is provided for text encryption, a browser interface will
open to allow interactive grille selection. For file encryption,
a key must be supplied as a JSON file (see the key generation command).

Examples:

  Encrypt text with interactive browser UI and export key:
    arcipher cardan encrypt "helloworld" --export key.json

  Encrypt text with exported/generated key:
    arcipher cardan encrypt ./key.json "helloworld"

  Encrypt a file with key and 4 threads:
    arcipher cardan encrypt key.json ./example/SunPoem ./example/SunPoem.enc --threads 4 -v

Notes:

  • The key defines the grille size and hole positions
  • A valid grille must cover every grid cell exactly once across all rotations
  • Messages shorter than the grid may be padded automatically
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cardanRunE(cmd, args, params, ciphers.Encrypt)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return cardanPreRunE(cmd, args, params)
		},
	}
	// cant use params.addFlags, because of block
	encryptCmd.Flags().IntVarP(&params.numCPU, "threads", "t", runtime.NumCPU()/2, "Amount of threads to be used")
	encryptCmd.Flags().StringVarP(&params.keyOutput, "export", "o", "", "Export the selected key to a file")

	return encryptCmd
}

func newCardanDecryptCommand() *cobra.Command {
	params := &cardanParams{}

	decryptCmd := &cobra.Command{
		Use:   "decrypt [key] <message | input> [output]",
		Short: "Decrypt a message or file using a Cardan grille cipher",
		Args:  cobra.RangeArgs(1, 3),
		Long: `Decrypt messages or files that were encrypted using the Cardan grille cipher.

A key must be provided for decryption and should match the one used
during encryption. Keys are supplied as JSON files (see the key
generation command).

Examples:

  Decrypt text with interactive browser UI and export key:
    arcipher cardan encrypt "h lowd e ol  lr " --export key.json

  Decrypt text with key:
    arcipher cardan decrypt ./key.json "h lowd e ol  lr "

  Decrypt a file with key and 4 threads:
    arcipher cardan decrypt key.json ./example/SunPoem.enc ./example/SunPoem --threads 4 -v

Notes:

  • The key must match the grille used for encryption
  • The grille defines the grid size and hole positions
  • Incorrect keys will produce invalid plaintext
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cardanRunE(cmd, args, params, ciphers.Decrypt)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return cardanPreRunE(cmd, args, params)
		},
	}
	// cant use params.addFlags, because of block
	decryptCmd.Flags().IntVarP(&params.numCPU, "threads", "t", runtime.NumCPU()/2, "Amount of threads to be used")
	decryptCmd.Flags().StringVarP(&params.keyOutput, "export", "o", "", "Export the selected key to a file")

	return decryptCmd
}

func newCardanGenerateKeyCommand() *cobra.Command {
	params := &cardanParams{}

	generateCmd := &cobra.Command{
		Use:   "generate-key <size> <output>",
		Short: "Generate a valid Cardan grille key",
		Args:  cobra.ExactArgs(2),
		Long: `Generate a valid Cardan grille key for use with the Cardan cipher.

This command generates a grille where the hole positions are arranged so
that every cell of the grid is covered exactly once across all rotations.
The generated key is saved as a JSON file and can later be used with the
encrypt and decrypt commands.

Examples:

  Generate a key for a 5x5 grid:
    arcipher cardan generate-key 5 key.json

Notes:

  • The size defines the dimensions of the grille (size x size)
  • The generated grille guarantees non-overlapping rotations
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cardanGenerateKeyRunE(cmd, args, params)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return cardanGenerateKeyPreRunE(args, params)
		},
	}

	return generateCmd
}
