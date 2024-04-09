package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/Luiso9/mari-go/config"
    "github.com/Luiso9/mari-go/commands"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	conf := config.GetConfig()
	if conf.Token == "" {
		log.Fatal("Token is missing in configuration")
	}

	sess, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		log.Fatal(err)
	}

	commands.RegisterCommands(sess)

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = sess.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}
	defer sess.Close()

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
