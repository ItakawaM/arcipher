package cmd

import (
	"fmt"
	"os"

	ciphers "github.com/ItakawaM/go-cryptotool/ciphers"
	"github.com/spf13/cobra"
)

var (
	message string
	key     int
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
		if message == "" {
			fmt.Println("Please provide a message to encrypt!")
			os.Exit(1)
		}

		if key <= 0 {
			fmt.Printf("Invalid key provided: [%d] is not viable", key)
			os.Exit(1)
		}

		encryptedMessage := ciphers.RailFenceEncryptMessage(message, key)
		fmt.Println(encryptedMessage)
	},
}

func init() {
	rootCmd.AddCommand(railfenceCmd)
	railfenceCmd.AddCommand(encryptCmd)

	encryptCmd.Flags().StringVarP(&message, "message", "m", "", "Message to encrypt")
	encryptCmd.Flags().IntVarP(&key, "key", "k", 0, "Key to use")

	encryptCmd.MarkFlagRequired("cipher")
	encryptCmd.MarkFlagRequired("key")
}
