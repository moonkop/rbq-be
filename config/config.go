package config

import (
	"../model"
	"../utils"
	"encoding/json"
	"io/ioutil"
)

var config model.Config

func ReadConfig() {
	data, err := ioutil.ReadFile("./config.json")
	utils.Check(err)
	var config model.Config
	json.Unmarshal(data, &config)
}
func GetConfig() model.Config {
	return config
}
