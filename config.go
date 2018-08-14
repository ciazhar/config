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

func (c *Config) DoMapify() (map[string]interface{}, error) {
	if m, ok := c.data.(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("can't type assert with map[string]interface{}")
}

func (c *Config) Get(key string) *Config {
	m, err := c.DoMapify()
	if err == nil {
		if val, ok := m[key]; ok {
			return &Config{val}
		}
	}
	return &Config{nil}
}

func (c *Config) Float() float64 {
	return c.data.(float64)
}

func (c *Config) String() string {
	return c.data.(string)
}

func (c *Config) Bool() bool {
	return c.data.(bool)
}