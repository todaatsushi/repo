package repoconf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadConfig() Config {
	configPath, exists := GetLibPath()
	if exists == false {
		log.Fatal("Config doesn't exist")
	}

	file, _ := ioutil.ReadFile(configPath)
	data := Config{}

	_ = json.Unmarshal([]byte(file), &data)
	return data
}
