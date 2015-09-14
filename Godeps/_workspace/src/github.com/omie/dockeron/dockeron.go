package dockeron

import (
	"fmt"

	"github.com/omie/shruti-cron/Godeps/_workspace/src/github.com/samalba/dockerclient"
)

type Dockeron struct {
	client *dockerclient.DockerClient
}

func (d *Dockeron) MakeContainer(job *Job) (string, error) {
	env := make([]string, 0)
	for _, e := range job.Environment {
		env = append(env, fmt.Sprintf("%s=%s", e.Key, e.Value))
	}

	containerConfig := &dockerclient.ContainerConfig{
		Image: job.Image,
		Env:   env,
	}
	containerId, err := d.client.CreateContainer(containerConfig, "")
	if err != nil {
		return "", err
	}

	return containerId, nil
}

func (d *Dockeron) StartContainer(containerId string, job *Job) error {
	hostConfig := &dockerclient.HostConfig{}
	hostConfig.Links = job.Links
	hostConfig.Binds = job.Binds

	err := d.client.StartContainer(containerId, hostConfig)
	if err != nil {
		return err
	}
	return nil
}
