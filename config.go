package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	Port int `yaml:"port"`
	RedisConfig struct {
		Address 	string 	`yaml:"address"`
		Port 		int 	`yaml:"port"`
		Password 	string 	`yaml:"password"`
		DB 			int 	`yaml:"db"`
	} `yaml:"redis"`
}

func GetConfig() *Config {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var c = &Config{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}