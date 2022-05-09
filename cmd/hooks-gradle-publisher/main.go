package main

import (
	"github.com/apex/log"
	hookGradle "github.com/ted-vo/hooks-gradle-publisher/pkg/hooks"
	"github.com/ted-vo/semantic-release/v3/pkg/hooks"
	"github.com/ted-vo/semantic-release/v3/pkg/plugin"
)

func main() {
	log.SetHandler(hookGradle.NewLogHandler())
	plugin.Serve(&plugin.ServeOpts{
		Hooks: func() hooks.Hooks {
			return &hookGradle.GradlePublisher{}
		},
	})
}
