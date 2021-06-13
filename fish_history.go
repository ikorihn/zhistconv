package zhistconv

import (
	"fmt"
	"strconv"

	"gopkg.in/yaml.v2"
)

type FishHistory struct {
	Cmd     string
	When    int
	History int64    `yaml:",omitempty"`
	Paths   []string `yaml:",omitempty"`
}

func ParseFishHistory(fishHistoryBytes []byte) ([]byte, error) {
	hist := make([]FishHistory, 0)
	err := yaml.Unmarshal(fishHistoryBytes, &hist)
	if err != nil {
		return nil, err
	}
	zshHist := convertToZshHistory(hist)
	return zshHist, nil
}

func convertToZshHistory(fishHist []FishHistory) []byte {
	//: 1621066935:0;brew list
	var zshHist string
	for _, v := range fishHist {
		zshHist += fmt.Sprintf(": %s:0;%s\n", strconv.Itoa(v.When), v.Cmd)
	}
	return []byte(zshHist)

}
