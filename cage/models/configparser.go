package models

import (
	"errors"
	"io/ioutil"
	"strings"
)

type ConfigParser struct {
	Path   string                       `json:"path"`
	Parsed map[string]map[string]string `json:"parsed"`
}

func (config *ConfigParser) Parse() (err error) {
	config.Parsed = make(map[string]map[string]string)

	if config.Path == "" {
		err = errors.New("No filepath received")
		return
	}
	raw, err := ioutil.ReadFile(config.Path)
	if err != nil {
		return
	}
	sections := strings.Split(strings.Replace(string(raw), "\r", "", -1), "\n\n")
	for _, v := range sections {
		lines := strings.Split(v, "\n")
		if !(strings.Contains(lines[0], "[") && strings.Contains(lines[0], "]")) {
			err = errors.New("Failed to parse format")
			return
		}
		n := strings.Replace(lines[0], "[", "", -1)
		n = strings.Replace(n, "]", "", -1)

		config.Parsed[n] = make(map[string]string)
		for _, v2 := range lines[1:] {
			line := strings.Split(v2, "=")
			key := line[0]
			val := line[1]
			config.Parsed[n][key] = val
		}
	}
	return
}
