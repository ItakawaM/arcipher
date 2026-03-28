package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cryptotool",
	Short: "Classic encryption ciphers in Go!",
	Long: `ItakawaM
	
	Work In Progress`,
	// TODO: Create a nice Long description for the cli tool itself
	// TODO: Replace current placeholder Long descriptions
	// TODO: Create shared RunE and PreRunE for simple ciphers: Vigenere, Railfence and Caesar

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&isVerbose, "verbose", "v", false, "Displays additional info")

	rootCmd.AddCommand(
		NewRailFenceCommand(),
		NewCaesarCommand(),
		NewCardanCommand(),
		NewVigenereCommand(),
	)
}
