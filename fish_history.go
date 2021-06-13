package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type FishHistory struct {
	cmd     string
	history int64
	paths   []string
}

func parseFishHistory(filePath string) ([]FishHistory, error) {
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

	hist := []FishHistory{}
	err = yaml.Unmarshal(bytes, &hist)
	if err != nil {
		return nil, err
	}
	fmt.Printf("history: %v\n", hist)
	return hist, err
}
