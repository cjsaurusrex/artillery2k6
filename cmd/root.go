package cmd

import (
	"embed"
	"github.com/spf13/cobra"
	"os"
)

//go:embed k6-script.tmpl
var K6ScriptTemplate embed.FS

var rootCmd = &cobra.Command{
	Use:   "Artillery2k6",
	Short: "Converts Artillery scripts to k6 scripts",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ArtilleryToK6.yaml)")
}
