package golang

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"log"

	"github.com/simplifyd-systems/buildman/pkg/planner/golang"
)

var (
	//go:embed templates/Dockerfile_go
	goBuildScript string
)

func Build(dir string, plan golang.GoPlan) (string, error) {
	return generateDockerfile(plan)
}

func generateDockerfile(plan golang.GoPlan) (string, error) {
	var b bytes.Buffer
	f := bufio.NewWriter(&b)

	var tpl *template.Template
	var err error

	// template Dockerfile
	tpl, err = template.New("build_script").Parse(goBuildScript)
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
