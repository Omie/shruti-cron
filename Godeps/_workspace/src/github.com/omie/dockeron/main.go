package dockeron

import (
	"os"

	"github.com/omie/shruti-cron/Godeps/_workspace/src/github.com/mgutz/logxi/v1"
	"github.com/omie/shruti-cron/Godeps/_workspace/src/github.com/samalba/dockerclient"
)

var (
	dk       *Dockeron
	dkLogger log.Logger
)

func init() {
	dkLogger = log.NewLogger(log.NewConcurrentWriter(os.Stdout), "[dockeron]")
	dkLogger.SetLevel(log.LevelAll)
	// try to connect to docker socket,
	// if it fails, makes no sense to go any further, just panic
	docker, err := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	if err != nil {
		dkLogger.Error(err.Error())
		panic(err)
	}
	dk = &Dockeron{client: docker}
}

func StartJobs(jobs []*Job) {
	dkLogger.Debug("--- starting jobs")
	for _, j := range jobs {
		go worker(j)
	}
}
