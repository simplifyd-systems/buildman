package commands

import (
	"fmt"
	"os"

	"github.com/simplifyd-systems/buildman/pkg/detect"
	"github.com/simplifyd-systems/buildman/pkg/plan"
	"github.com/simplifyd-systems/buildman/pkg/planner/golang"
	"github.com/simplifyd-systems/buildman/pkg/planner/next"
	"github.com/spf13/cobra"
)

type PlanFlags struct {
	Publish  bool
	Registry string
	AppPath  string
}

// Plan an image from source code
func Plan() *cobra.Command {
	var flags PlanFlags

	cmd := &cobra.Command{
		Use: "plan <image-name>",
		// Args:    cobra.ExactArgs(1),
		Short:   "Generate a plan to build app image from source code",
		Example: "buildman plan test_img --path apps/test-app --builder cnbs/sample-builder:bionic",
		Long: "Buildman Build uses Cloud Native Buildpacks to create a runnable app image from source code.\n\nPack Build " +
			"requires an image name, which will be generated from the source code. Build defaults to the current directory, " +
			"but you can use `--path` to specify another source code directory. Build requires a `builder`, which can either " +
			"be provided directly to build using `--builder`, or can be set using the `set-default-builder` command. For more " +
			"on how to use `pack build`, see: https://buildpacks.io/docs/app-developer-guide/build-an-app/.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validatePlanFlags(&flags); err != nil {
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

			buildPlan, err := plan.Plan(dir, framework)
			if err != nil {
				return err
			}

			if framework == detect.GoFramework {
				buildPlan.(golang.GoPlan).Print()
			} else if framework == detect.NextFramework {
				buildPlan.(next.NextPlan).Print()
			}

			return nil
		},
	}

	planCommandFlags(cmd, &flags)
	AddHelpFlag(cmd, "build")
	return cmd
}

func planCommandFlags(cmd *cobra.Command, planFlags *PlanFlags) {
	cmd.Flags().StringVarP(&planFlags.AppPath, "path", "p", "", "Path to app dir or zip-formatted file (defaults to current working directory)")
}

func validatePlanFlags(flags *PlanFlags) error {
	if flags.Registry != "" {
		return fmt.Errorf("")
	}

	return nil
}
