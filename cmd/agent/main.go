package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kubeshop/kubtest-executor-postman/pkg/runner/newman"
	"github.com/kubeshop/kubtest/pkg/api/kubtest"
)

func main() {

	args := os.Args
	if len(args) == 1 {
		fmt.Println("missing input argument")
		os.Exit(1)
	}

	script := args[1]

	e := kubtest.Execution{}
	json.Unmarshal([]byte(script), &e)
	runner := newman.NewNewmanRunner()
	result := runner.Run(e)
	fmt.Println(result)
	fmt.Printf("$$$%s$$$", e.Id)
}