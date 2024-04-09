package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
    Token string `json:"token"` 
}

func loadConfig() (Config, error) {
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

func main () {
	config, err := loadConfig()
    if err != nil {
        fmt.Println("Failed to load config:", err)
        os.Exit(1) // Exit with an error
    }

	sess, err := discordgo.New("Bot " + config.Token)
	if err !=  nil {	
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "hello" {
			s.ChannelMessageSend(m.ChannelID, "World!")
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("Online~")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

