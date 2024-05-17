package vue

import (
	"fmt"

	"github.com/simplifyd-systems/buildman/pkg/detect"
)

type VuePlan struct {
	Packages        []string
	InstallCmd      string
	BuildCmd        string
	StartCmd        string
	OutputDir       string
	NodeVersion     string
	ArgsPlaceholder string
}

func Plan(dir string) (VuePlan, error) {
	return processConfig()
}

func processConfig() (VuePlan, error) {
	return VuePlan{
		OutputDir:   "build",
		NodeVersion: "22",
		ArgsPlaceholder: `
{{range .Args}}
ARG {{.}}
{{end}}`,
	}, nil
}

func (plan VuePlan) Print() {
	fmt.Printf(`
************ Buildman Build Plan v1 ************

Detected Framework: %s
Build Command	  : %s
Install Command   : %s
Start Command     : %s

Output Directory  : %s
Node version      : %s

************           END          ************

`, detect.ReactFramework, plan.BuildCmd, plan.InstallCmd, plan.StartCmd, plan.OutputDir, plan.NodeVersion)
}
