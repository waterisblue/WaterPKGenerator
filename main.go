package main

import (
	"pkgenerate/config"
	"pkgenerate/log"
)

func main() {
	configMap := config.GetConfig()

	log.Info.Println(configMap["pk.cache.num"])
}
