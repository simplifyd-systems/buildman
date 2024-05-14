package main

import (
	"os"

	"github.com/simplifyd-systems/buildman/cmd"

	"github.com/simplifyd-systems/buildman/internal/commands"
)

func main() {
	/*
		fmt.Println("buildman")

		framework, err := detect.Detect(".")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		dockerfile, err := build.Build(".", framework)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println(dockerfile)
		fmt.Println()
		fmt.Println()
		fmt.Println()

	*/

	rootCmd, err := cmd.NewBuildmanCommand()
	if err != nil {
		os.Exit(1)
	}

	ctx := commands.CreateCancellableContext()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
