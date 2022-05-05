package repoconf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func WriteConfig(data Config) error {
	configPath, exists := GetLibPath()
	filename, _ := GetLibPath()
	if exists == false {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("write: couldn't create %s", filename)
		}
		defer file.Close()
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
