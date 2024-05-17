package react

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"

	_ "embed"

	"github.com/simplifyd-systems/buildman/pkg/planner/react"
)

var (
	//go:embed templates/Dockerfile_reactjs
	reactJSBuildScript string
)

func Build(dir string, plan react.ReactPlan) (string, error) {
	return generateDockerfile(plan)
}

func generateDockerfile(plan react.ReactPlan) (string, error) {
	var b bytes.Buffer
	f := bufio.NewWriter(&b)

	// template Dockerfile
	tpl, err := template.New("build_script").Parse(reactJSBuildScript)
	if err != nil {
		fmt.Printf("generateDockerfile failure: %s\n", err)
		return "", err
	}

	err = tpl.Execute(f, plan)
	if err != nil {
		log.Print("Template execute: ", err)
		return "", err
	}

	f.Flush()

	return string(b.String()), nil
}
