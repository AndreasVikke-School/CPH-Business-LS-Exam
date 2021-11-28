package main

import (
	"encoding/json"
	"os"

	eh "api_service/errorhandler"
)

type Configuration struct {
	Kafka struct {
		RequestTopic string   `json:"request_topic"`
		Brokers      []string `json:"brokers"`
	} `json:"kafka"`
	Redis struct {
		Service string `json:"service"`
		Broker  string `json:"broker"`
	} `json:"redis"`
}

func getConfig(env string) Configuration {
	file, _ := os.Open("../../configs/" + env + "_conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	eh.PanicOnError(err, "Failed to Decode Env File")
	return configuration
}
