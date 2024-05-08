package build

import (
	"github.com/simplifyd-systems/buildman/pkg/builder/next"
	"github.com/simplifyd-systems/buildman/pkg/detect"
)

func Build(dir string, framework detect.Framework) {
	if framework == detect.NextFramework {
		next.Build(dir)
	}
}
