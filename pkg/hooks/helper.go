package hooks

import (
	"bytes"
	"io/ioutil"
)

var FILE_NAME = "gradle.properties"

func UpdateVersion(oldVersion, newVersion string) error {
	input, err := ioutil.ReadFile(FILE_NAME)
	if err != nil {
		return err
	}

	output := bytes.Replace(input, []byte(oldVersion), []byte(newVersion), -1)

	if err = ioutil.WriteFile(FILE_NAME, output, 0666); err != nil {
		return err
	}

	return nil
}
