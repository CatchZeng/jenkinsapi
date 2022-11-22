package jenkins

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/go-resty/resty/v2"
)

// API https://jenkins.xxx.com/api/#Create%20Job

// Jenkins jenkins struct
type Jenkins struct {
	Host     string
	UserName string
	APIToken string
}

// CreateItem create jenkins job with name and config
func (j *Jenkins) CreateItem(name, config string) error {
	return j.doJenkinsRequest("/createItem", name, config)
}

// CreateItemInFolder create jenkins job in folder with name and config
func (j *Jenkins) CreateItemInFolder(folder, name, config string) error {
	// reference https://gist.github.com/stuart-warren/7786892
	return j.doJenkinsRequest("job/"+folder+"/createItem", name, config)
}

func (j *Jenkins) doJenkinsRequest(path, name, config string) error {
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

type Job struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Color string `json:"color"`
}

type JobsResponse struct {
	Jobs []Job `json:"jobs"`
}

// https://jenkins.xxx.com/job/xxx/api/json?pretty=true&tree=jobs[name,url,color]{0,10}
func (j *Jenkins) GetJobsInFolder(folder string, m, n int) ([]Job, error) {
	jobs := []Job{}

	if len(folder) < 1 {
		return jobs, errors.New("folder is empty")
	}

	u, err := url.Parse(j.Host)
	if err != nil {
		return jobs, err
	}
	jobPath := fmt.Sprintf("job/%v", folder)
	tailPath := fmt.Sprintf("/api/json?tree=jobs[name,url,color]{%v,%v}", m, n)
	u.Path = path.Join(u.Path, jobPath)
	url := u.String() + tailPath

	client := resty.New()
	client.SetRetryCount(2)
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBasicAuth(j.UserName, j.APIToken).
		SetResult(&JobsResponse{}).
		Get(url)
	if err != nil {
		return jobs, err
	}
	result := resp.Result().(*JobsResponse)
	return result.Jobs, nil
}

// https://jenkins.xxx.com/job/xxx/job/xxx/api/json
func (j *Jenkins) GetMultiBranchJobs(job Job, m, n int) ([]Job, error) {
	jobs := []Job{}

	if len(job.URL) < 1 {
		return jobs, errors.New("job URL is empty")
	}

	tailPath := fmt.Sprintf("/api/json?tree=jobs[name,url,color]{%v,%v}", m, n)
	url := job.URL + tailPath

	client := resty.New()
	client.SetRetryCount(2)
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBasicAuth(j.UserName, j.APIToken).
		SetResult(&JobsResponse{}).
		Get(url)
	if err != nil {
		return jobs, err
	}
	result := resp.Result().(*JobsResponse)
	return result.Jobs, nil
}

type BuildDetail struct {
	Number     int         `json:"number"`
	Result     string      `json:"result"`
	Timestamp  int64       `json:"timestamp"`
	URL        string      `json:"url"`
	ChangeSets []ChangeSet `json:"changeSets"`
}

type ChangeSet struct {
	Items []ChangeSetItem `json:"items"`
	Kind  string          `json:"kind"`
}

type ChangeSetItem struct {
	AffectedPaths []string `json:"affectedPaths"`
	CommitId      string   `json:"commitId"`
	Timestamp     int64    `json:"timestamp"`
	AuthorEmail   string   `json:"authorEmail"`
	Comment       string   `json:"comment"`
	Date          string   `json:"date"`
}

// https://jenkins.xxx.com/job/xxx/job/xxx/job/main/lastBuild/api/json
func (j *Jenkins) GetLastBuild(job Job) (*BuildDetail, error) {
	detail := &BuildDetail{}

	if len(job.URL) < 1 {
		return detail, errors.New("job URL is empty")
	}

	tailPath := "/lastBuild/api/json"
	url := job.URL + tailPath

	client := resty.New()
	client.SetRetryCount(2)
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBasicAuth(j.UserName, j.APIToken).
		SetResult(&BuildDetail{}).
		Get(url)
	if err != nil {
		return detail, err
	}
	result := resp.Result().(*BuildDetail)
	return result, nil
}

type Build struct {
	Number int    `json:"number"`
	URL    string `json:"url"`
}

type BuildsResponse struct {
	Builds []Build `json:"builds"`
}

// https://jenkins.xxx.com/job/xxx/job/xxx/job/main/api/json
func (j *Jenkins) GetJobBuilds(job Job, m, n int) ([]Build, error) {
	builds := []Build{}

	if len(job.URL) < 1 {
		return builds, errors.New("job URL is empty")
	}

	tailPath := fmt.Sprintf("/api/json?tree=builds[number,url]{%v,%v}", m, n)
	url := job.URL + tailPath

	client := resty.New()
	client.SetRetryCount(2)
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBasicAuth(j.UserName, j.APIToken).
		SetResult(&BuildsResponse{}).
		Get(url)
	if err != nil {
		return builds, err
	}
	result := resp.Result().(*BuildsResponse)
	return result.Builds, nil
}

// https://jenkins.xxx.com/job/xxx/job/xxx/job/main/412/api/json
func (j *Jenkins) GetBuildDetail(build Build) (*BuildDetail, error) {
	detail := &BuildDetail{}

	if len(build.URL) < 1 {
		return detail, errors.New("build URL is empty")
	}

	tailPath := "/api/json"
	url := build.URL + tailPath

	client := resty.New()
	client.SetRetryCount(2)
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBasicAuth(j.UserName, j.APIToken).
		SetResult(&BuildDetail{}).
		Get(url)
	if err != nil {
		return detail, err
	}
	result := resp.Result().(*BuildDetail)
	return result, nil
}
