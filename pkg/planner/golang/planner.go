package golang

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/simplifyd-systems/buildman/pkg/detect"
)

type GoPlan struct {
	GoVersion       string
	Packages        []string
	InstallCmd      string
	BuildCmd        string
	StartCmd        string
	ArgsPlaceholder string
}

func Plan(dir string) (GoPlan, error) {
	var goModConfig []byte

	// read the go.mod file
	if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
		goModConfig, err = os.ReadFile(filepath.Join(dir, "go.mod"))
		if err != nil {
			fmt.Println(err)
			fmt.Println("cannot read go.mod")
			return GoPlan{}, fmt.Errorf("cannot read go.mod")
		}
	} else {
		fmt.Println("go.mod not found")
		return GoPlan{}, fmt.Errorf("go.mod not found")
	}

	return processGoModConfig(string(goModConfig))
}

func processGoModConfig(contents string) (GoPlan, error) {
	goVersion := "1.22"

	scanner := bufio.NewScanner(strings.NewReader(contents))
	for scanner.Scan() {
		line := scanner.Text()

		// get go version
		if strings.Contains(line, "go 1.") {
			parts := strings.SplitN(line, " ", 2)
			if len(parts) != 2 {
				return GoPlan{}, fmt.Errorf("cannot read go version from go.mod file")
			}

			goVersion = strings.Trim(parts[1], " '\",")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	return GoPlan{
		GoVersion: goVersion,
		StartCmd:  "/go/bin/main",
		BuildCmd:  "go build -o main .",
		ArgsPlaceholder: `
{{range .Args}}
ARG {{.}}
{{end}}`,
	}, nil
}

func (plan GoPlan) Print() {
	fmt.Printf(`
************ Buildman Build Plan v1 ************

Detected Framework: %s
Build Command	  : %s
Install Command   : %s
Start Command     : %s

************           END          ************

`, detect.GoFramework, plan.BuildCmd, plan.InstallCmd, plan.StartCmd)
}
