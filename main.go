package main

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/bitrise-io/go-utils/log"
)

func main() {
	configContents := os.Getenv("telegraf_conf")
	configFile, err := ioutil.TempFile("", "telegraf.conf")
	if err != nil {
		failf("Failed to create telegraf.conf file: %s", err)
	}
	defer configFile.Close()
	configFile.WriteString(configContents)

	telegrafCmd := exec.Command("telegraf", "--config", configFile.Name())
	log.Infof("Starting Telegraf agent in the background: %s", telegrafCmd.String())
	err = telegrafCmd.Start()
	if err != nil {
		failf("Failed to start command: %s", err)
	}

	os.Exit(0)
}

func failf(format string, v ...interface{}) {
	log.Errorf(format, v...)
	os.Exit(1)
}
