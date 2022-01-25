package main

import (
	"log"
	"os"

	"github.com/CatchZeng/jenkinsapi/cmd"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	cmd.Execute()
}
