// Package config contains everything related to reading config files
package config

import (
	"git.sr.ht/~hjertnes/doing/utils"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Path string `yaml:"path"`
}

func read(filename string) (*Config, error){
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	result := &Config{}

	err = yaml.Unmarshal(content, result)

	return result, err
}

func create(filename string) (*Config, error){
	conf := &Config{
		Path: "~/txt/roam/",
	}
	content, err := yaml.Marshal(conf)
	if err != nil{
		return conf, err
	}
	err = ioutil.WriteFile(filename, content, os.ModePerm)
	return conf, err
}

func Read() (*Config, error){
	filename := utils.ReplaceTilde("~/.doing.yml")
	if utils.Exist(filename){
		return read(filename)
	} else {
		return create(filename)
	}
}

