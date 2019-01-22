package main

import (
	"fmt"

	ch "./bin/configHandler"
)

func main() {
	configs, err := ch.GetConfigs()
	fmt.Println(configs, err)
}
