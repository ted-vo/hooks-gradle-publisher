package main

import (
	"log"
	"os"

	gradle "github.com/go-semantic-release/hooks-gradle-publisher/pkg/hooks"
	"github.com/go-semantic-release/semantic-release/v2/pkg/hooks"
	"github.com/go-semantic-release/semantic-release/v2/pkg/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		Hooks: func() hooks.Hooks {
			return &gradle.GradlePublisher{
				Logger: log.New(os.Stderr, "", 0),
			}
		},
	})
}
