package main

import (
	"encoding/json"
	"fmt"
	"os"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	hb "github.com/opsgenie/opsgenie-go-sdk/heartbeat"
)

type Config struct {
	APIKey        string `json:"apiKey"`
	HeartbeatName string `json:"heartbeatName"`
	Description   string `json:"description"`
}

func main() {
	// Load config from JSON file
	configFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	var cfg Config
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&cfg); err != nil {
		panic(err)
	}

	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(cfg.APIKey)

	hbCli, cliErr := cli.Heartbeat()
	if cliErr != nil {
		panic(cliErr)
	}

	// create the heartbeat
	enabled := true
	req := hb.AddHeartbeatRequest{
		Name:         cfg.HeartbeatName,
		IntervalUnit: "minutes",
		Enabled:      &enabled,
		Interval:     5,
		Description:  cfg.Description,
	}
	response, hbErr := hbCli.Add(req)
	if hbErr != nil {
		panic(hbErr)
	}

	fmt.Printf("Name: %s\n", response.Name)
	fmt.Printf("Status: %s\n", response.Status)
	fmt.Printf("Code: %d\n", response.Code)

	// send heart beat request
	pingRequest := hb.PingHeartbeatRequest{Name: response.Name}
	pingResponse, sendErr := hbCli.Ping(pingRequest)

	if sendErr != nil {
		panic(sendErr)
	}

	fmt.Println()
	fmt.Printf("Heartbeat request sent\n")
	fmt.Printf("----------------------\n")
	fmt.Printf("RequestId: %s\n", pingResponse.RequestID)
	fmt.Printf("Response Time: %f\n", pingResponse.ResponseTime)
}
