package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const yamlFile = "./config.yaml"

type WorkFile struct {
	PathWorkFile   string `yaml:"path_file_process"`
	PathAnswerFile string `yaml:"path_file_count"`
	PathCountFile  string `yaml:"path_file_answer"`
	DNSImageLoader string `yaml:"bd_connect_image"`
	DNSFileLoader  string `yaml:"bd_connect_file"`
}

func NenConfig() (*WorkFile, error) {
	var confRead WorkFile

	yamlFile, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &confRead)
	if err != nil {
		return nil, err
	}

	return &confRead, nil
}
