package next

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"

	_ "embed"

	"github.com/simplifyd-systems/buildman/pkg/planner/next"
)

var (
	//go:embed templates/Dockerfile_nextjs_static
	nextJSStaticBuildScript string

	//go:embed templates/Dockerfile_nextjs_non_static
	nextJSNonStaticBuildScript string
)

func Build(dir string, plan next.NextPlan) (string, error) {
	return generateDockerfile(plan)
}

func generateDockerfile(plan next.NextPlan) (string, error) {
	var b bytes.Buffer
	f := bufio.NewWriter(&b)

	var tpl *template.Template
	var err error

	// template Dockerfile
	if plan.Export || plan.Standalone {
		tpl, err = template.New("build_script").Parse(nextJSStaticBuildScript)
		if err != nil {
			fmt.Printf("generateDockerfile failure: %s\n", err)
			return "", err
		}
	} else {
		tpl, err = template.New("build_script").Parse(nextJSNonStaticBuildScript)
		if err != nil {
			fmt.Printf("generateDockerfile failure: %s\n", err)
			return "", err
		}
	}

	err = tpl.Execute(f, plan)
	if err != nil {
		log.Print("Template execute: ", err)
		return "", err
	}

	f.Flush()

	return string(b.String()), nil
}
