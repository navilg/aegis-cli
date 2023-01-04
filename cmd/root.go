/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/navilg/aegis-cli/internal"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aegis-cli",
	Short: "Aegis compatible OTP generator",
	Long:  `Aegis-cli is CLI tool to generate OTP from aegis vault file.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: func(cmd *cobra.Command, args []string) error {

		homedir, _ := os.UserHomeDir()
		vaultdata, err := internal.ReadVault(homedir + "/.config/aegis-cli/aegis.json")

		// fmt.Println(*vaultdata)

		decrypteddb, err := internal.DecryptDB(*vaultdata)

		if err != nil {
			if err.Error() == "error: Unable to authenticate using provided password." {
				fmt.Println("Error: Unable to authenticate using provided password.")
				os.Exit(0)
			}
			return err
		}

		var db internal.DB
		json.Unmarshal(decrypteddb, &db)

		internal.ItemsTUI(db)

		return err
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aegis-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
