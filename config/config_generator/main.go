// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package main

import (
	"fmt"
	"os"

	"github.com/mattermost/mattermost-server/config/config_generator/generator"
)

func main() {
	outputFile := os.Getenv("OUTPUT_CONFIG")
	if outputFile == "" {
		fmt.Println("Output file name is missing. Please set OUTPUT_CONFIG env variable to absolute path")
		os.Exit(2)
	}
	if _, err := os.Stat(outputFile); !os.IsNotExist(err) {
		_, _ = fmt.Fprintf(os.Stderr, "File %s already exists. Not overwriting!\n", outputFile)
		os.Exit(2)
	}
	fmt.Println(outputFile)

	if file, err := os.Create("C://Users/ajw0409/.babun/cygwin/home/ajw0409/go/src/github.com/mattermost/mattermost-server/dist/mattermost/config/config.json"); err == nil {
		err = generator.GenerateDefaultConfig(file)
		//fmt.Println("여기서는?")
		_ = file.Close()
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}

}
