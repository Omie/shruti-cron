package dockeron

type Environment struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Job struct {
	Name        string        `json:"name"`
	Image       string        `json:"image"`
	Command     string        `json:"command"`
	Environment []Environment `json:"environment"`
	Interval    int64         `json:"interval"`
	Links       []string      `json:"links"`
	Binds       []string      `json:"binds"`
}

type Jobs []*Job
