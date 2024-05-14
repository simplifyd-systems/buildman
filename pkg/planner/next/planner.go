package next

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/simplifyd-systems/buildman/pkg/detect"
)

type NextPlan struct {
	Packages        []string
	InstallCmd      string
	BuildCmd        string
	StartCmd        string
	Export          bool
	Standalone      bool
	OutputDir       string
	NodeVersion     string
	ArgsPlaceholder string
}

func Plan(dir string) (NextPlan, error) {
	var nextConfig []byte

	if _, err := os.Stat(filepath.Join(dir, "next.config.js")); err == nil {
		nextConfig, err = os.ReadFile(filepath.Join(dir, "next.config.js"))
		if err != nil {
			fmt.Println(err)
			fmt.Println("cannot read next.config.js")
			return NextPlan{}, fmt.Errorf("cannot read next.config.js")
		}
	} else if _, err := os.Stat(filepath.Join(dir, "next.config.mjs")); err == nil {
		nextConfig, err = os.ReadFile(filepath.Join(dir, "next.config.mjs"))
		if err != nil {
			fmt.Println(err)
			fmt.Println("cannot read next.config.mjs")
			return NextPlan{}, fmt.Errorf("cannot read next.config.mjs")
		}
	} else {
		fmt.Println("next.config not found")
		return NextPlan{}, fmt.Errorf("next.config not found")
	}

	return processNextConfig(string(nextConfig))
}

func processNextConfig(contents string) (NextPlan, error) {
	export := false
	standalone := false
	outputDir := ""

	scanner := bufio.NewScanner(strings.NewReader(contents))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "output") && strings.Contains(line, "export") {
			export = true
			outputDir = "out"
		} else if strings.Contains(line, "output") && strings.Contains(line, "standalone") {
			standalone = true
		} else if strings.Contains(line, "distDir") {
			_, v := parseKeyValuePair(line)
			outputDir = v
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	return NextPlan{
		Export:      export,
		Standalone:  standalone,
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

func (plan NextPlan) Print() {
	fmt.Printf(`
************ Buildman Build Plan v1 ************

Detected Framework: %s
Build Command	  : %s
Install Command   : %s
Start Command     : %s

Export            : %v
Standalone        : %v
Output Directory  : %s
Node version      : %s

************           END          ************

`, detect.NextFramework, plan.BuildCmd, plan.InstallCmd, plan.StartCmd, plan.Export, plan.Standalone, plan.OutputDir, plan.NodeVersion)
}
