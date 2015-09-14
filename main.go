package main // import "github.com/omie/shruti-cron"

import (
	"os"
	"time"

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

	host := os.Getenv("SHRUTI_CRON_HOST")
	port := os.Getenv("SHRUTI_CRON_PORT")
	if host == "" || port == "" {
		cronLogger.Error("main: host or port not set")
		return
	}

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

	go func() {
		// wait until other containers are up
		// there should be a much better solution,
		// hack prevails
		time.Sleep(1 * time.Minute)
		dockeron.StartJobs(jobs)
	}()

	cronLogger.Info("Attempting to start HTTP server")
	err = StartHTTPServer(host, port)
	if err != nil {
		cronLogger.Error("Error starting server", err)
	}
}
