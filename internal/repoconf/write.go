package repoconf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func WriteConfig(data Config) error {
	configPath, exists := GetLibPath()
	if exists == false {
		log.Fatal("Write: config doesn't exist")
	}

	content, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Write: couldn't load data to json")
	}
	err = ioutil.WriteFile(configPath, content, 0644)
	if err != nil {
		log.Fatal("Write: couldn't write data to json")
	}
	return err
}
