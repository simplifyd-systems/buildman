package vite

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "embed"
)

var (
	//go:embed templates/Dockerfile_vite
	viteBuildScript string
)

type ViteBuildConfig struct {
	Export          bool
	OutputDir       string
	NodeVersion     string
	ArgsPlaceholder string
}

func Build(dir string) (string, error) {
	var viteConfig []byte
	// read the vite.config.js or vite.config.mjs file
	if _, err := os.Stat(filepath.Join(dir, "vite.config.js")); err == nil {
		viteConfig, err = os.ReadFile(filepath.Join(dir, "vite.config.js"))
		if err != nil {
			fmt.Println(err)
			fmt.Println("cannot read vite.config.js")
			return "", fmt.Errorf("cannot read vite.config.js")
		}
	} else {
		fmt.Println("vite.config not found")
		return "", fmt.Errorf("vite.config not found")
	}

	config, err := processViteConfig(string(viteConfig))
	if err != nil {
		fmt.Println(err)
		fmt.Println("cannot process vite.config")
		return "", fmt.Errorf("cannot process vite.config")
	}

	log.Println(config)

	return generateDockerfile(config)
}

func processViteConfig(contents string) (ViteBuildConfig, error) {
	export := false
	outputDir := ""

	fmt.Println(contents)

	if strings.Contains(contents, "output") {
		fmt.Println("output line detected")
	}

	scanner := bufio.NewScanner(strings.NewReader(contents))
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if strings.Contains(line, "output") && strings.Contains(contents, "export") {
			export = true
			outputDir = "out"
		} else if strings.Contains(line, "distDir") {
			_, v := parseKeyValuePair(line)
			outputDir = v
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	return ViteBuildConfig{
		Export:      export,
		OutputDir:   outputDir,
		NodeVersion: "22",
		ArgsPlaceholder: `
{{range .Args}}
ARG {{.}}
{{end}}`,
	}, nil
}

func parseKeyValuePair(input string) (key, value string) {
	parts := strings.SplitN(input, ":", 2)
	if len(parts) != 2 {
		return "", ""
	}

	key = strings.Trim(parts[0], " '\",")
	value = strings.Trim(parts[1], " '\",")
	return key, value
}

func generateDockerfile(config ViteBuildConfig) (string, error) {
	var b bytes.Buffer
	f := bufio.NewWriter(&b)

	// template Dockerfile
	tpl, err := template.New("build_script").Parse(viteBuildScript)
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
