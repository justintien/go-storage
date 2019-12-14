package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadConfig(config interface{}, name string) (err error) {
	filepath := "./conf/" + name + ".json"
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Printf("can't open config file: [%s] %s", filepath, err)
		return
	}

	err = json.Unmarshal(data, config)

	return
}
