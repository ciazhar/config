package config

import (
	"os"
	"fmt"
	"encoding/json"
	"errors"
)

type Config struct {
	data interface{}
}

func Load()	*Config {
	var c Config

	if len(os.Args)<=1{
		fmt.Println("Please Set Config File In Command Line")
		os.Exit(1)
	}

	configFile, err := os.Open(os.Args[1]+".json")
	if err!=nil {
		fmt.Println("Error No Such File In Directory")
		os.Exit(1)
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&c.data)
	if err!=nil {
		fmt.Println("Error Parse Json")
		os.Exit(1)
	}
	return &c
}

func (c *Config) doMapify() (map[string]interface{}, error) {
	if m, ok := c.data.(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("can't type assert with map[string]interface{}")
}

func (c *Config) Get(key string) *Config {
	m, err := c.doMapify()
	if err == nil {
		if val, ok := m[key]; ok {
			return &Config{val}
		}
	}
	return &Config{nil}
}

func (c *Config) Float() float64 {
	if s,ok :=c.data.(float64); ok {
		return s
	}
	fmt.Println("Error Conversion, Field Is Not Float")
	os.Exit(1)
	return 0
}

func (c *Config) String() string {

	if s,ok :=c.data.(string); ok {
		return s
	}
	fmt.Println("Error Conversion, Field Is Not String")
	os.Exit(1)
	return ""
}

func (c *Config) Bool() bool {
	if s,ok :=c.data.(bool); ok {
		return s
	}
	fmt.Println("Error Conversion, Field Is Not Boolean")
	os.Exit(1)
	return false
}