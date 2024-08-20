package config

import (
	"bufio"
	"io"
	"os"
	"pkgenerate/log"
	"strings"
)

var config map[string]string
var path = "./pk.properties"

func init() {
	config = make(map[string]string)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Error.Println(err.Error())
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Error.Println(err.Error())
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
}

func GetConfig() map[string]string {
	return config
}
