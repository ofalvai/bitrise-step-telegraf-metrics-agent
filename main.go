package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/bitrise-io/go-utils/log"
)

const linuxPackageUrl = "https://dl.influxdata.com/telegraf/releases/telegraf_1.18.0-1_amd64.deb"

func main() {
	if runtime.GOOS == "linux" {
		installOnLinux()
	}

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

func installOnLinux() {
	log.Infof("Downloading Telegraf deb package...")

	debFile, err := ioutil.TempFile("", "telegraf.deb")
	defer debFile.Close()
	if err != nil {
		failf("Failed to create telegraf.deb file: %s", err)
	}

	resp, err := http.Get(linuxPackageUrl)
	if err != nil {
		failf("Failed to download Telegraf package from %s: %s", linuxPackageUrl, err)
	}
	defer resp.Body.Close()
	log.Infof("Download successful")

	_, err = io.Copy(debFile, resp.Body)
	if err != nil {
		failf("Failed to save Telegraf package to disk: %s", err)
	}

	log.Infof("Installing package...")
	_, err = exec.Command("sudo", "dpkg", "-i", debFile.Name()).CombinedOutput()
	if err != nil {
		failf("Failed to install package: %s", err)
	}
	log.Infof("Telegraf successfully installed")
}
