package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

var I info

type info struct {
	Application `yaml:"application"`
	Webserver   `yaml:"webserver"`
	Mysql       `yaml:"mysql"`
}

type Application struct {
	Domain           string `yaml:"domain"`
	Short_url_length int    `yaml:"short_url_length"`
}
type Webserver struct {
	Listen string `yaml:"listen"`
}
type Mysql struct {
	Ip       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func initconfig() error {
	f := flag.String("f", "./l2s.yml", "Specifying a configuration file")
	flag.Parse()
	yamlFile, err := ioutil.ReadFile(*f)
	if err != nil {
		return fmt.Errorf("Failed to read file:%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &I)
	if err != nil {
		return fmt.Errorf("Unmarshal: %v", err)
	}
	return nil
}
