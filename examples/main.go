package main

import (
	"github.com/ciazhar/config"
	"fmt"
)


func main() {
	conf :=config.Load()
	fmt.Println(conf.Get("port").Float())
	fmt.Println(conf.Get("database").Get("name").String())
}
