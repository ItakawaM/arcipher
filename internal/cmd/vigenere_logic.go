package cmd

import (
	"fmt"

	"github.com/ItakawaM/go-cryptotool/ciphers"
	"github.com/ItakawaM/go-cryptotool/internal/benchmark"
	"github.com/ItakawaM/go-cryptotool/internal/engine"
	"github.com/spf13/cobra"
)

type vigenereParams struct {
	key []byte
	blockCipherParams
}

func vigenerePreRunE(command *cobra.Command, params *vigenereParams, args []string) error {
	params.key = []byte(args[0])

	switch len(args[1:]) {
	case 1:
		return params.parseSourceMessageParams(command)

	case 2:
		if !fileExists(args[1]) {
			return fmt.Errorf("provided input file does not exist: %s", args[1])
		}
		return params.parseSourceFileParams()

	default:
		return fmt.Errorf("invalid working mode")
	}
}

func vigenereRunE(command *cobra.Command, args []string, params *vigenereParams, mode ciphers.CipherMode) error {
	if isVerbose {
		defer benchmark.MeasurePerformance(fmt.Sprintf("caesar %s", mode))()
	}

	switch len(args[1:]) {
	case 1:
		vigenereCipher, vigenereErr := ciphers.NewVigenereCipher(params.key)
		if vigenereErr != nil {
			return vigenereErr
		}

		message := args[1]
		buffer := []byte(message)

		var err error
		switch mode {
		case ciphers.Encrypt:
			err = vigenereCipher.EncryptBlock(buffer, buffer)
		case ciphers.Decrypt:
			err = vigenereCipher.DecryptBlock(buffer, buffer)
		}
		if err != nil {
			return err
		}

		command.Printf("'%s'", string(buffer))
	case 2:
		inFilePath := args[1]
		outFilePath := args[2]

		blockSizeBytes := params.blockSize * 1024
		vigenereCipher, caesarErr := ciphers.NewVigenereCipher(params.key)
		if caesarErr != nil {
			return caesarErr
		}

		return engine.NewBlockEngine(blockSizeBytes, params.numCPU).ProcessFile(vigenereCipher, mode, inFilePath, outFilePath)
	}

	return nil
}
