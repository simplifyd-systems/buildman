package build

import (
	"fmt"

	"github.com/simplifyd-systems/buildman/pkg/builder/next"
	"github.com/simplifyd-systems/buildman/pkg/builder/react"
	"github.com/simplifyd-systems/buildman/pkg/detect"
)

func Build(dir string, framework detect.Framework) (string, error) {
	if framework == detect.NextFramework {
		return next.Build(dir)
	} else if framework == detect.ReactFramework {
		return react.Build(dir)
	}

	return "", fmt.Errorf("cannot build the %s framework", framework)
}
