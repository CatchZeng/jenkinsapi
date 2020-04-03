package cmd

import (
	"log"

	"github.com/CatchZeng/jenkinsapi/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "jenkinsapi version",
	Long:  `jenkinsapi version`,
	Run:   runVersionCmd,
}

func runVersionCmd(_ *cobra.Command, _ []string) {
	version := version.GetVersion()
	log.Println(version)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
