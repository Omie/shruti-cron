package dockeron

import (
	"time"
)

func worker(job *Job) {
	containerId, err := dk.MakeContainer(job)
	if err != nil {
		dkLogger.Error("Error creating container:", job.Name, err.Error())
		return
	}

	for {
		err = dk.StartContainer(containerId, job)
		if err != nil {
			dkLogger.Error("Error starting container:", job.Name, err.Error())
			return
		}
		time.Sleep(time.Duration(job.Interval) * time.Minute)
	}

}
