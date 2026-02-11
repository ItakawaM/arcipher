package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	ciphers "github.com/ItakawaM/go-cryptotool/ciphers"
	"github.com/spf13/cobra"
)

var (
	message  string
	filename string
	key      int
)

// railfenceCmd represents the railfence command
var railfenceCmd = &cobra.Command{
	Use:   "railfence",
	Short: "Classic railfence cipher",
	Long:  `Description WIP`, // TODO: Description
}

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a given message with a key",
	Long:  `Description WIP`, // TODO: Description
	Run: func(cmd *cobra.Command, args []string) {
		startTime := time.Now()

		stat, _ := os.Stdin.Stat()
		if message == "" && filename == "" && (stat.Mode()&os.ModeCharDevice) == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			var builder strings.Builder
			for scanner.Scan() {
				builder.WriteString(scanner.Text())
			}

			message = builder.String()
		}

		if message == "" && filename == "" {
			fmt.Println("Please provide an input to encrypt!")
			os.Exit(1)
		}

		if key <= 0 {
			fmt.Printf("Invalid key provided: [%d] is not viable", key)
			os.Exit(1)
		}

		if message != "" {
			encryptedMessage := ciphers.RailFenceEncryptMessage(message, key)
			fmt.Println(encryptedMessage)
		} else {
			if err := ciphers.RailFenceEncryptFile(filename, key); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		elapsed := time.Since(startTime)
		fmt.Printf("Encrypting Took: %s", elapsed)
	},
}

func init() {
	rootCmd.AddCommand(railfenceCmd)
	railfenceCmd.AddCommand(encryptCmd)

	encryptCmd.Flags().StringVarP(&message, "message", "m", "", "Message to encrypt")
	encryptCmd.Flags().StringVarP(&filename, "file", "f", "", "File to encrypt")
	encryptCmd.Flags().IntVarP(&key, "key", "k", 0, "Key to use")

	encryptCmd.MarkFlagRequired("key")

	encryptCmd.MarkFlagsMutuallyExclusive("file", "message")
}
