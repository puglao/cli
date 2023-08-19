package cli

import (
	"fmt"
	"os"

	"github.com/puglao/cli/config"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(config.Initiate)
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVarP(&config.File, "config", "c", "", "config file (default is ./config.yaml)")
}

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Base CLI application for Go",
	Long:  `Base CLI application for Go, can be used as a template for other CLI applications`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
