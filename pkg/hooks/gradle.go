package hooks

import (
	"os/exec"

	"github.com/apex/log"
	"github.com/ted-vo/semantic-release/v3/pkg/hooks"
)

var NAME = "Gradle Publisher"
var FUVERSION = "dev"

type GradlePublisher struct{}

func (gp *GradlePublisher) Init(m map[string]string) error {
	log.Infof("Init %v", m)
	return nil
}

func (gp *GradlePublisher) Name() string {
	return NAME
}

func (gp *GradlePublisher) Version() string {
	return FUVERSION
}

func (gp *GradlePublisher) Success(config *hooks.SuccessHookConfig) error {
	oldVersion := config.PrevRelease.Version
	newVersion := config.NewRelease.Version
	log.Infof("old version: " + oldVersion)
	log.Infof("new version: " + newVersion)
	if err := gradlePublish(); err != nil {
		return err
	}
	return nil
}

func (gp *GradlePublisher) NoRelease(config *hooks.NoReleaseConfig) error {
	log.Infof("reason: " + config.Reason.String())
	log.Infof("message: " + config.Message)
	return nil
}

func gradlePublish() error {
	log.Infof("Start gradle publish...")

	out, err := exec.Command("./gradlew", "publish").Output()
	if err != nil {
		return err
	}
	log.Infof("Result %s", out)
	return nil
}
