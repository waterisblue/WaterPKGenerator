package main

import (
	"fmt"
	"pkgenerate/config"
)

func main() {
	config.InitConfig("./pk.properties")
	configMap := config.GetConfig()

	fmt.Println(configMap["pk.cache.num"])
}
