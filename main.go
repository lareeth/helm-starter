package main

import (
	"log"
	"os"

	"github.com/lareeth/helm-starter/cmd"
)

func main() {
	err := os.MkdirAll(os.Getenv("HELM_DATA_HOME")+"/starters", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
