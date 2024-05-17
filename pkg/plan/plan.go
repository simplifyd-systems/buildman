package plan

import (
	"fmt"

	"github.com/simplifyd-systems/buildman/pkg/detect"
	"github.com/simplifyd-systems/buildman/pkg/planner/golang"
	"github.com/simplifyd-systems/buildman/pkg/planner/next"
	"github.com/simplifyd-systems/buildman/pkg/planner/react"
	"github.com/simplifyd-systems/buildman/pkg/planner/vue"
)

func Plan(dir string, framework detect.Framework) (any, error) {
	if framework == detect.GoFramework {
		return golang.Plan(dir)
	} else if framework == detect.NextFramework {
		return next.Plan(dir)
	} else if framework == detect.ReactFramework {
		return react.Plan(dir)
	} else if framework == detect.VueFramework {
		return vue.Plan(dir)
	}

	return "", fmt.Errorf("cannot build the %s framework", framework)
}
