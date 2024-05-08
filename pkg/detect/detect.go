package detect

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Framework string

const ReactFramework Framework = "React"
const VueFramework Framework = "Vue"
const NextFramework Framework = "Next"

const NotDetectedFramework Framework = "NotDetectedFramework"

func Detect(dir string) (Framework, error) {
	if _, err := os.Stat(filepath.Join(dir, "package.json")); err == nil {
		// read in package.json
		packageJSONContents, err := os.ReadFile(filepath.Join(dir, "package.json"))
		if err != nil {
			fmt.Println(err)
			return NotDetectedFramework, err
		}
		if strings.Contains(string(packageJSONContents), "vue-cli-service") {
			fmt.Println("VueJS application detected")
			return VueFramework, nil
		} else if strings.Contains(string(packageJSONContents), "\"next\"") {
			fmt.Println("NextJS application detected")
			return NextFramework, nil
		} else if strings.Contains(string(packageJSONContents), "\"react\"") {
			fmt.Println("ReactJS application detected")
			return ReactFramework, nil
		}
	}

	return NotDetectedFramework, nil
}
