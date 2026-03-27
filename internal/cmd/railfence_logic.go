package cmd

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/ItakawaM/go-cryptotool/ciphers"
	"github.com/ItakawaM/go-cryptotool/internal/benchmark"
	"github.com/ItakawaM/go-cryptotool/internal/engine"
	"github.com/spf13/cobra"
)

type railfenceParams struct {
	key int
	blockCipherParams
}

func railfencePreRunE(command *cobra.Command, params *railfenceParams, args []string) error {
	key, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	} else if key < 1 {
		return fmt.Errorf("key must be >= 1")
	}
	params.key = key

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

func railfenceRunE(command *cobra.Command, args []string, params *railfenceParams, mode ciphers.CipherMode) error {
	if isVerbose {
		defer benchmark.MeasurePerformance(fmt.Sprintf("railfence %s", mode))()
	}

	switch len(args[1:]) {
	case 1:
		message := args[1]

		railFenceCipher, railFenceErr := ciphers.NewRailFenceCipher(params.key, len(message))
		if railFenceErr != nil {
			return railFenceErr
		}

		src := []byte(message)
		dst := bytes.Clone(src)

		var err error
		switch mode {
		case ciphers.Encrypt:
			err = railFenceCipher.EncryptBlock(dst, src)
		case ciphers.Decrypt:
			err = railFenceCipher.DecryptBlock(dst, src)
		}
		if err != nil {
			return err
		}

		command.Printf("'%s'", string(dst))
		return nil

	case 2:
		inFilePath := args[1]
		outFilePath := args[2]

		blockSizeBytes := params.blockSize * 1024
		railFenceCipher, err := ciphers.NewRailFenceCipher(params.key, blockSizeBytes)
		if err != nil {
			return err
		}

		return engine.NewBlockEngine(blockSizeBytes, params.numCPU).ProcessFile(railFenceCipher, mode, inFilePath, outFilePath)

	default:
		return fmt.Errorf("invalid working mode")

	}
}
