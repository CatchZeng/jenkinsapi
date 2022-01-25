package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/CatchZeng/jenkinsapi/config"
	jks "github.com/CatchZeng/jenkinsapi/jenkins"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "createItem on jenkins",
	Long:  `createItem on jenkins`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(_ *cobra.Command, args []string) {
		host, user, token, err := getConfig()
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		var jenkins = jks.Jenkins{
			Host:     host,
			UserName: user,
			APIToken: token,
		}

		if len(job) < 1 {
			log.Fatal("job name can not be empty")
			return
		}

		if len(folder) < 1 {
			if err := jenkins.CreateItem(job, configXML); err != nil {
				log.Fatal(err.Error())
				return
			}
			return
		}

		if err := jenkins.CreateItemInFolder(folder, job, configXML); err != nil {
			log.Fatal(err.Error())
		}
	},
}

func getConfig() (string, string, string, error) {
	const errorFmt = "jenkins %s is missing, please add in ~ /.jenkinsapi/config.yaml file, the key is '%s'"

	host, err := config.GetConfig(config.JenkinsHost)
	if err != nil || len(host) < 1 {
		err = fmt.Errorf(errorFmt, config.JenkinsHost, config.JenkinsHost)
		return host, "", "", err
	}

	user, err := config.GetConfig(config.JenkinsUserName)
	if err != nil || len(user) < 1 {
		err = fmt.Errorf(errorFmt, config.JenkinsUserName, config.JenkinsUserName)
		return host, user, "", err
	}

	token, err := config.GetConfig(config.JenkinsAPIToken)
	if err != nil || len(token) < 1 {
		err = fmt.Errorf(errorFmt, config.JenkinsAPIToken, config.JenkinsAPIToken)
		return host, user, token, err
	}

	return host, user, token, err
}

var folder, job, configXML string

func init() {
	rootCmd.AddCommand(create)
	create.Flags().StringVarP(&folder, "folder", "f", "", "jenkins folder name")
	create.Flags().StringVarP(&job, "job", "j", "", "jenkins job name")
	create.Flags().StringVarP(&configXML, "config", "c", "", "the content of jenkins job config.xml")
}
