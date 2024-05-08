package next

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
	//go:embed templates/Dockerfile_nextjs_static
	nextJSBuildScript string
)

type NextJSBuildConfig struct {
	Export          bool
	OutputDir       string
	NodeVersion     string
	ArgsPlaceholder string
}

func Build(dir string) error {
	var nextConfig []byte
	// read the next.config.js or next.config.mjs file
	if _, err := os.Stat(filepath.Join(dir, "next.config.js")); err == nil {
		nextConfig, err = os.ReadFile("next.config.js")
		if err != nil {
			fmt.Println(err)
			fmt.Println("cannot read next.config.js")
			return fmt.Errorf("cannot read next.config.js")
		}
	} else if _, err := os.Stat(filepath.Join(dir, "next.config.mjs")); err == nil {
		nextConfig, err = os.ReadFile("next.config.mjs")
		if err != nil {
			fmt.Println(err)
			fmt.Println("cannot read next.config.mjs")
			return fmt.Errorf("cannot read next.config.mjs")
		}
	} else {
		fmt.Println("next.config not found")
		return fmt.Errorf("next.config not found")
	}

	config, err := processNextConfig(string(nextConfig))
	if err != nil {
		fmt.Println(err)
		fmt.Println("cannot process next.config")
		return fmt.Errorf("cannot process next.config")
	}

	log.Println(config)

	fmt.Println(generateDockerfile(config))

	return nil
}

func processNextConfig(contents string) (NextJSBuildConfig, error) {
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

	return NextJSBuildConfig{
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

func generateDockerfile(config NextJSBuildConfig) (string, error) {
	var b bytes.Buffer
	f := bufio.NewWriter(&b)

	// template Dockerfile
	tpl, err := template.New("build_script").Parse(nextJSBuildScript)
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
