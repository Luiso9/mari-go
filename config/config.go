package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
    Token string `json:"token"` 
    Prefix string `json:"prefix"`
}

func LoadConfig() (Config, error) {
    // Read the config file
    configFile, err := os.Open("config.json")
    if err != nil {
        return Config{}, fmt.Errorf("error opening config.json: %v", err)
    }
    defer configFile.Close()

    jsonBytes, err := ioutil.ReadAll(configFile)
    if err != nil {
        return Config{}, fmt.Errorf("error reading config.json: %v", err)
    }

    // Parse JSON into the Config struct
    var config Config
    err = json.Unmarshal(jsonBytes, &config)
    if err != nil {
        return Config{}, fmt.Errorf("error parsing config.json: %v", err)
    }

    return config, nil
}