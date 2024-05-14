package plan

import (
	"fmt"

	"github.com/simplifyd-systems/buildman/pkg/detect"
	"github.com/simplifyd-systems/buildman/pkg/planner/golang"
	"github.com/simplifyd-systems/buildman/pkg/planner/next"
)

func Plan(dir string, framework detect.Framework) (any, error) {
	if framework == detect.GoFramework {
		return golang.Plan(dir)
	} else if framework == detect.NextFramework {
		return next.Plan(dir)
	}

	return "", fmt.Errorf("cannot build the %s framework", framework)
}
