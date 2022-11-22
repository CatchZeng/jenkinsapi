package cmd

import (
	"log"

	v "github.com/go-ecosystem/utils/v2/version"
	"github.com/spf13/cobra"
)

const (
	version   = "1.0.2"
	buildTime = "2022/01/21"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "jenkinsapi version",
	Long:  `jenkinsapi version`,
	Run:   runVersionCmd,
}

func runVersionCmd(_ *cobra.Command, _ []string) {
	v := v.Stringify(version, buildTime)
	log.Println(v)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
