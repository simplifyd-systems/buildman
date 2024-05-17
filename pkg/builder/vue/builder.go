package vue

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"

	_ "embed"

	"github.com/simplifyd-systems/buildman/pkg/planner/vue"
)

var (
	//go:embed templates/Dockerfile_vuejs
	vueJSBuildScript string
)

func Build(dir string, plan vue.VuePlan) (string, error) {
	return generateDockerfile(plan)
}

func generateDockerfile(plan vue.VuePlan) (string, error) {
	var b bytes.Buffer
	f := bufio.NewWriter(&b)

	// template Dockerfile
	tpl, err := template.New("build_script").Parse(vueJSBuildScript)
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
