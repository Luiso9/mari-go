package config
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Config struct {
    Token  string `json:"token"`
    Prefix string `json:"prefix"`
}

var loadedConfig *Config

func LoadConfig() error {
    configFile, err := os.Open("config.json")
    if err != nil {
        return fmt.Errorf("error opening config.json: %v", err)
    }
    defer configFile.Close()

    jsonBytes, err := ioutil.ReadAll(configFile)
    if err != nil {
        return fmt.Errorf("error reading config.json: %v", err)
    }

    var config Config
    err = json.Unmarshal(jsonBytes, &config)
    if err != nil {
        return fmt.Errorf("error parsing config.json: %v", err)
    }

    loadedConfig = &config 
    return nil
}

func GetConfig() *Config {
    return loadedConfig
}