package cli

import (
	"fmt"

	"github.com/puglao/cli/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of CLI",
	Long:  `All software has versions. This is CLI's`,
	Run:   versionCmdRun,
}

func versionCmdRun(cmd *cobra.Command, args []string) {
	fmt.Print(version.Info())
}
