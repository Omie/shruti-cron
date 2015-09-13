package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/omie/dockeron"
)

func getConfFiles(dir, ext string) (confs []os.FileInfo, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	confs = make([]os.FileInfo, 0)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ext) {
			confs = append(confs, f)
		}
	}
	if len(confs) == 0 {
		return nil, errors.New("No conf files found")
	}
	return confs, nil
}

func getParsedConfig(basePath string, files []os.FileInfo) (jobs dockeron.Jobs, err error) {
	allJobs := dockeron.Jobs{}

	for _, c := range files {

		b, err := ioutil.ReadFile(filepath.Join(basePath, c.Name()))
		if err != nil {
			return nil, err
		}

		jobsInFile := dockeron.Jobs{}
		err = json.Unmarshal(b, &jobsInFile)
		if err != nil {
			return nil, err
		}

		for _, job := range jobsInFile {
			allJobs = append(allJobs, job)
		}
	}
	return allJobs, nil
}
