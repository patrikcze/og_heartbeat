package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/opsgenie/opsgenie-go-sdk-v2/heartbeat"
)

type config struct {
	APIKey        string `json:"apiKey"`
	HeartbeatName string `json:"heartbeatName"`
}

func main() {
	// Read the config file
	configBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config file: ", err)
		return
	}

	var cfg config
	err = json.Unmarshal(configBytes, &cfg)
	if err != nil {
		fmt.Println("Error parsing config file: ", err)
		return
	}

	for {
		hbCli, err := heartbeat.NewClient(cfg.APIKey)
		if err != nil {
			fmt.Println("Error creating heartbeat client: ", err)
			return
		}

		_, err = hbCli.Ping(cfg.HeartbeatName)
		if err != nil {
			fmt.Println("Error generating heartbeat: ", err)
			return
		}

		fmt.Println("Generated OpsGenie heartbeat.")
		time.Sleep(5 * time.Minute)
	}
}
