package main // import "github.com/omie/shruti-cron"

import (
	"os"

	"github.com/mgutz/logxi/v1"
	"github.com/omie/dockeron"
)

const (
	CONF_DIR = "conf"
	CONF_EXT = ".conf"
)

var (
	cronLogger log.Logger
)

func main() {
	cronLogger = log.NewLogger(log.NewConcurrentWriter(os.Stdout), "[shruti-cron]")
	cronLogger.SetLevel(log.LevelAll)

	files, err := getConfFiles(CONF_DIR, CONF_EXT)
	if err != nil {
		cronLogger.Error("err getting conf files: ", err)
		return
	}

	jobs, err := getParsedConfig(CONF_DIR, files)
	if err != nil {
		cronLogger.Error("err parsing conf files: ", err)
		return
	}
	cronLogger.Debug("parsed config:", jobs)

	dockeron.StartJobs(jobs)

	cronLogger.Info("Attempting to start HTTP server")
	err = StartHTTPServer("127.0.0.1", "9577")
	if err != nil {
		cronLogger.Error("Error starting server", err)
	}
}
