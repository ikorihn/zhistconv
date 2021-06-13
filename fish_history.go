package zhistconv

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type FishHistory struct {
	Cmd     string
	History int64 `yaml:",omitempty"`
	When    int64
	Paths   []string `yaml:",omitempty"`
}

func ParseFishHistory(filePath string) ([]FishHistory, error) {
	if filePath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		filePath = homeDir + "/.local/share/fish/fish_history"
	}

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	hist := make([]FishHistory, 0)
	err = yaml.Unmarshal(bytes, &hist)
	if err != nil {
		return nil, err
	}
	return hist, err
}
