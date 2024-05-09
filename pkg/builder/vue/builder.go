package vue

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"

	_ "embed"
)

var (
	//go:embed templates/Dockerfile_vuejs
	vueJSBuildScript string
)

type VueJSBuildConfig struct {
	Export          bool
	OutputDir       string
	NodeVersion     string
	ArgsPlaceholder string
}

func Build(dir string) (string, error) {
	return generateDockerfile(VueJSBuildConfig{
		Export:      true,
		OutputDir:   "build",
		NodeVersion: "22",
		ArgsPlaceholder: `
{{range .Args}}
ARG {{.}}
{{end}}`,
	})
}

func generateDockerfile(config VueJSBuildConfig) (string, error) {
	var b bytes.Buffer
	f := bufio.NewWriter(&b)

	// template Dockerfile
	tpl, err := template.New("build_script").Parse(vueJSBuildScript)
	if err != nil {
		fmt.Printf("generateDockerfile failure: %s\n", err)
		return "", err
	}

	err = tpl.Execute(f, config)
	if err != nil {
		log.Print("Template execute: ", err)
		return "", err
	}

	f.Flush()

	return string(b.String()), nil
}
