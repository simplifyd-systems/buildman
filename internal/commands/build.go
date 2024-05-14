package commands

import (
	"fmt"
	"os"

	"github.com/simplifyd-systems/buildman/pkg/build"
	"github.com/simplifyd-systems/buildman/pkg/detect"
	"github.com/spf13/cobra"
)

type BuildFlags struct {
	Publish  bool
	Registry string
	AppPath  string
}

// Build an image from source code
func Build() *cobra.Command {
	var flags BuildFlags

	cmd := &cobra.Command{
		Use: "build <image-name>",
		// Args:    cobra.ExactArgs(1),
		Short:   "Generate app image from source code",
		Example: "buildman build test_img --path apps/test-app --builder cnbs/sample-builder:bionic",
		Long: "Buildman Build uses Cloud Native Buildpacks to create a runnable app image from source code.\n\nPack Build " +
			"requires an image name, which will be generated from the source code. Build defaults to the current directory, " +
			"but you can use `--path` to specify another source code directory. Build requires a `builder`, which can either " +
			"be provided directly to build using `--builder`, or can be set using the `set-default-builder` command. For more " +
			"on how to use `pack build`, see: https://buildpacks.io/docs/app-developer-guide/build-an-app/.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateBuildFlags(&flags); err != nil {
				return err
			}

			dir, err := os.Getwd()
			if err != nil {
				return err
			}

			framework, err := detect.Detect(dir)
			if err != nil {
				return err
			}

			fmt.Printf("\n%s framework detected\n", framework)

			dockerfile, err := build.Build(dir, framework, "")
			if err != nil {
				return err
			}

			fmt.Printf(`
************ Builman Dockerfile v1 ************
%s
************     		END		   ************

`, dockerfile)

			return nil
		},
	}

	buildCommandFlags(cmd, &flags)
	AddHelpFlag(cmd, "build")
	return cmd
}

func buildCommandFlags(cmd *cobra.Command, buildFlags *BuildFlags) {
	cmd.Flags().StringVarP(&buildFlags.AppPath, "path", "p", "", "Path to app dir or zip-formatted file (defaults to current working directory)")
}

func validateBuildFlags(flags *BuildFlags) error {
	if flags.Registry != "" {
		return fmt.Errorf("")
	}

	return nil
}
