package main

import (
	"fmt"
	"log"
	"os"

	"github.com/simplifyd-systems/buildman/pkg/build"
	"github.com/simplifyd-systems/buildman/pkg/detect"
)

func main() {
	fmt.Println("buildman")

	framework, err := detect.Detect(".")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	build.Build(".", framework)
}
