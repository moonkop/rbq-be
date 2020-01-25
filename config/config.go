package config

import (
	"encoding/json"
	"io/ioutil"
	"rbq-be/model"
	"rbq-be/utils"
)

var config model.Config

func ReadConfig() model.Config {
	data, err := ioutil.ReadFile("./config.json")
	utils.Check(err)
	json.Unmarshal(data, &config)
	return config
}
func GetConfig() model.Config {
	return config
}
