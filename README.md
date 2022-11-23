# jenkinsapi

> jenkinsapi is the go implementation of jenkins api. Support Docker, Jenkinsfile, command line, go module.

## Feature

- [x] Support [Docker](#Docker)
- [x] Support [Jenkinsfile](#Jenkinsfile)
- [x] Support [module](#module)
- [x] CreateItem
      ![demo](/images/demo.png)

  ![demo-folder](/images/demo-folder.png)

## Install

## Docker pull

```shell
docker pull catchzeng/jenkinsapi
```

## go install

```shell
# Go 1.16+
go install github.com/CatchZeng/jenkinsapi@v1.1.0

# Go version < 1.16
go get -u github.com/CatchZeng/jenkinsapi@v1.1.0
```

## Usage

### Config.yaml

create `config.yaml` in `$HOME/.jenkinsapi` folder.

```yaml
# $HOME/.jenkinsapi/config.yaml
host: "https://jenkins.catchzeng.com/"
username: "admin"
token: "xazcasdasdasdadfsfsaefew"
```

### Docker

```shell
docker run -v $HOME/.jenkinsapi/:/root/.jenkinsapi catchzeng/jenkinsapi jenkinsapi create -j "job name" -f "folder name" -c "the content of jenkins job config.xml"
```

### Jenkinsfile

```groovy
pipeline {
    agent {
        docker {
            image 'catchzeng/jenkinsapi'
            args '-v $HOME/.jenkinsapi/:/root/.jenkinsapi -u root'
        }
    }
    stages {
        stage('Test') {
            steps {
                sh 'jenkinsapi create -j "job name" -f "folder name" -c "the content of jenkins job config.xml"'
            }
        }
    }
}
```

### module

```go
package main

import (
	"fmt"

	jks "github.com/CatchZeng/jenkinsapi/jenkins"
)

var jenkins = jks.Jenkins{
	Host:     "https://jenkins.catchzeng.com/",
	UserName: "admin",
	APIToken: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
}

func main() {
  const folder = "abc-folder"
	const name = "abc"
	const gitURL = "https://github.com/CatchZeng/abc.git"
  xml := fmt.Sprintf(config, name, gitURL)

  // create job with name and xml config
	if err := jenkins.CreateItem(name, xml); err != nil {
		log.Printf(err.Error())
	}

  // create job in folder with name and xml config
	if err := jenkins.CreateItemInFolder(folder, name, xml); err != nil {
		log.Printf(err.Error())
	}
}

// get from jenkins master node
const config = `<?xml version='1.1' encoding='UTF-8'?>
<org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject plugin="workflow-multibranch@2.21">
  <actions/>
  <description></description>
......
  <triggers>
    <com.igalg.jenkins.plugins.mswt.trigger.ComputedFolderWebHookTrigger plugin="multibranch-scan-webhook-trigger@1.0.5">
      <spec></spec>
      <token>%s</token>
    </com.igalg.jenkins.plugins.mswt.trigger.ComputedFolderWebHookTrigger>
  </triggers>
  <disabled>false</disabled>
  <sources class="jenkins.branch.MultiBranchProject$BranchSourceList" plugin="branch-api@2.5.5">
    <data>
      <jenkins.branch.BranchSource>
        <source class="jenkins.plugins.git.GitSCMSource" plugin="git@4.0.0">
          <id>blueocean</id>
          <remote>%s</remote>
          <credentialsId>jenkins-generated-ssh-key</credentialsId>
......
</org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject>`
```
