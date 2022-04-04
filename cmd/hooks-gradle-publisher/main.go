package main

import (
	"github.com/apex/log"
	hookGradle "github.com/go-semantic-release/hooks-gradle-publisher/pkg/hooks"
	"github.com/go-semantic-release/semantic-release/v2/pkg/hooks"
	"github.com/go-semantic-release/semantic-release/v2/pkg/plugin"
)

func main() {
	log.SetHandler(hookGradle.NewLogHandler())
	plugin.Serve(&plugin.ServeOpts{
		Hooks: func() hooks.Hooks {
			return &hookGradle.GradlePublisher{}
		},
	})
}
