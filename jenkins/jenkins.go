package jenkins

import (
	"bytes"
	"fmt"
	"net/http"
)

// API https://jenkins.xxx.com/api/#Create%20Job

// Jenkins jenkins struct
type Jenkins struct {
	Host     string
	UserName string
	APIToken string
}

// CreateItem create jenkins job with name and config
func (j *Jenkins) CreateItem(name string, config string) error {
	return j.doJenkinsRequest("/createItem", name, config)
}

// CreateItemInFolder create jenkins job in folder with name and config
func (j *Jenkins) CreateItemInFolder(folder string, name string, config string) error {
	// reference https://gist.github.com/stuart-warren/7786892
	return j.doJenkinsRequest("job/"+folder+"/createItem", name, config)
}

func (j *Jenkins) doJenkinsRequest(path string, name string, config string) error {
	request := NewRequest(j.Host).
		SetHeader("Content-Type", "application/xml").
		SetParam("name", name).
		SetBasicAuth(j.UserName, j.APIToken)

	buffer := []byte(config)
	response, err := request.Post(path, bytes.NewBuffer(buffer), nil)
	if err != nil {
		return err
	}
	if response.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf("respose: %v", response)
}
