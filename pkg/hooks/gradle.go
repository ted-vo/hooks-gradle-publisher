package hooks

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/apex/log"
	"github.com/ted-vo/semantic-release/v3/pkg/hooks"
)

var NAME = "Gradle Publisher"
var FUVERSION = "dev"

type GradlePublisher struct {
	CMD string
}

func (gp *GradlePublisher) Init(m map[string]string) error {
	log.Infof("Init %v", m)
	gp.CMD = m["cmd"]

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
	if err := gp.gradlePublish(); err != nil {
		return err
	}
	return nil
}

func (gp *GradlePublisher) NoRelease(config *hooks.NoReleaseConfig) error {
	log.Infof("reason: " + config.Reason.String())
	log.Infof("message: " + config.Message)
	return nil
}

func (gp *GradlePublisher) gradlePublish() error {
	log.Infof("Start gradle publish...")

	cmd := gp.CMD
	cmdArgs := strings.Fields(cmd)

	if len(gp.CMD) == 0 {
		cmdArgs = append(cmdArgs, "./gradlew", "publish")
	}

	cmdPipe := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	log.Infof("Command: %s %s", cmdArgs[0], strings.Join(cmdArgs[1:], " "))

	stdout, err := cmdPipe.StdoutPipe()
	if err != nil {
		log.Infof("error oucring when publishing. Detail: %s", err.Error())
		return err
	}
	if err := cmdPipe.Start(); err != nil {
		log.Infof("error oucring when publishing. Detail: %s", err.Error())
		return err
	}

	// print the output of the subprocess
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		log.Infof(m)
	}

	if err := cmdPipe.Wait(); err != nil {
		log.Infof("error oucring when publishing. Detail: %s", err.Error())
		return err
	}
	return nil
}
