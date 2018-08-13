package config

import (
	"os"
	"fmt"
	"encoding/json"
)
var config map[string]interface{}

func Load() {

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
	err = jsonParser.Decode(&config)
}

func GetString(s string) string {
	if config[s]!= nil{
		return config[s].(string)
	}else {
		fmt.Println("field "+s+" not exist")
		os.Exit(1)
		return ""
	}
}

func GetInt(s string) int {
	if config[s]!= nil {
		return config[s].(int)
	}else {
		fmt.Println("field not exist")
		os.Exit(1)
		return 0
	}
}

func GetBool(s string) bool {
	if config[s]!= nil{
		return config[s].(bool)
	}else {
		fmt.Println("field not exist")
		os.Exit(1)
		return false
	}
}