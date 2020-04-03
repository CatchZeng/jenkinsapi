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
			log.Printf(err.Error())
			return
		}

		var jenkins = jks.Jenkins{
			Host:     host,
			UserName: user,
			APIToken: token,
		}

		if len(job) < 1 {
			log.Printf("job name can not be empty")
			return
		}

		if len(folder) < 1 {
			if err := jenkins.CreateItem(job, configXML); err != nil {
				log.Printf(err.Error())
				return
			}
			return
		}

		if err := jenkins.CreateItemInFolder(folder, job, configXML); err != nil {
			log.Printf(err.Error())
			return
		}
	},
}

func getConfig() (host string, user string, token string, err error) {
	const errorFmt = "Jenkins %s is missing, please add in ~ /.jenkinsapi/config.yaml file, the key is '%s'"

	host, err = config.GetConfig(config.JenkinsHost)
	if len(host) < 1 {
		err = fmt.Errorf(errorFmt, config.JenkinsHost, config.JenkinsHost)
		return
	}

	user, err = config.GetConfig(config.JenkinsUserName)
	if len(user) < 1 {
		err = fmt.Errorf(errorFmt, config.JenkinsUserName, config.JenkinsUserName)
		return
	}

	token, err = config.GetConfig(config.JenkinsAPIToken)
	if len(token) < 1 {
		err = fmt.Errorf(errorFmt, config.JenkinsAPIToken, config.JenkinsAPIToken)
		return
	}

	return
}

var folder, job, configXML string

func init() {
	rootCmd.AddCommand(create)
	create.Flags().StringVarP(&folder, "folder", "f", "", "jenkins folder name")
	create.Flags().StringVarP(&job, "job", "j", "", "jenkins job name")
	create.Flags().StringVarP(&configXML, "config", "c", "", "the content of jenkins job config.xml")
}
