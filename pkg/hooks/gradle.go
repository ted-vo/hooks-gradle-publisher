package hooks

import (
	"log"
	"os/exec"

	"github.com/go-semantic-release/semantic-release/v2/pkg/hooks"
)

var NAME = "Gradle-Publisher"
var FUVERSION = "dev"

type GradlePublisher struct {
	Logger *log.Logger
}

func (t *GradlePublisher) Init(m map[string]string) error {
	t.Logger.Printf("init: %v\n", m)
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	t.Logger.Printf("The date is %s\n", out)
	return nil
}

func (t *GradlePublisher) Name() string {
	return NAME
}

func (t *GradlePublisher) Version() string {
	return FUVERSION
}

func (t *GradlePublisher) Success(config *hooks.SuccessHookConfig) error {
	t.Logger.Println("old version: " + config.PrevRelease.Version)
	t.Logger.Println("new version: " + config.NewRelease.Version)
	t.Logger.Printf("commit count: %d\n", len(config.Commits))
	return nil
}

func (t *GradlePublisher) NoRelease(config *hooks.NoReleaseConfig) error {
	t.Logger.Println("reason: " + config.Reason.String())
	t.Logger.Println("message: " + config.Message)
	return nil
}
