package build

import (
	"fmt"

	"github.com/simplifyd-systems/buildman/pkg/builder/golang"
	"github.com/simplifyd-systems/buildman/pkg/builder/next"
	"github.com/simplifyd-systems/buildman/pkg/builder/react"
	"github.com/simplifyd-systems/buildman/pkg/builder/vue"
	"github.com/simplifyd-systems/buildman/pkg/detect"
	golang_planner "github.com/simplifyd-systems/buildman/pkg/planner/golang"
	next_planner "github.com/simplifyd-systems/buildman/pkg/planner/next"
	react_planner "github.com/simplifyd-systems/buildman/pkg/planner/react"
	vue_planner "github.com/simplifyd-systems/buildman/pkg/planner/vue"
)

func Build(dir string, framework detect.Framework, plan any) (string, error) {
	if framework == detect.NextFramework {
		return next.Build(dir, plan.(next_planner.NextPlan))
	} else if framework == detect.ReactFramework {
		return react.Build(dir, plan.(react_planner.ReactPlan))
	} else if framework == detect.VueFramework {
		return vue.Build(dir, plan.(vue_planner.VuePlan))
	} else if framework == detect.GoFramework {
		return golang.Build(dir, plan.(golang_planner.GoPlan))
	}

	return "", fmt.Errorf("cannot build the %s framework", framework)
}
