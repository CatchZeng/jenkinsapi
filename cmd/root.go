package cmd

import (
	"fmt"
	"os"

	"github.com/CatchZeng/jenkinsapi/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jenkinsapi",
	Short: "jenkinsapi is a command line tool that talks with Jenkins API.",
	Long:  "jenkinsapi is a command line tool that talks with Jenkins API.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
