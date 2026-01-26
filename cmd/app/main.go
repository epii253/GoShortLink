package main

import (
	"fmt"
	"os"
	"path"
	fileparsers "project/internal/parser"
	"project/internal/requesters"
)

func main() {

	args := os.Args

	if len(args) <= 1 || 2 < len(args) {
		return
	}

	filePath := args[1]

	if !path.IsAbs(filePath) {
		return
	}

	allInf, err := fileparsers.Parse(filePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	workersCup := 10 // TODO extract from program arguments

	requesters.SetupScan(workersCup, allInf)
}
