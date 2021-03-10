package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Host        string
		Port        string
		Username    string
		Password    string
		Name        string
		Connections int
	} `yaml:"database"`
	Server struct {
		Host string
		Port string
		Log  string
	} `yaml:"server"`
	Email struct {
		Host string
		Port string
		Name string
	} `yaml: "email"`
}

func main() {
	yamlFile, err := os.Open("sample.yaml")
	if err != nil {
		fmt.Println(err)
	}

	var cfg Config
	yaml.NewDecoder(yamlFile).Decode(&cfg)

	fmt.Println(cfg)
}
