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
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Terjadi error pada loadConfig", err)
	}

	sess, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
	}

	commands.RegisterCommands(sess)

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = sess.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
	}
	defer sess.Close()

	fmt.Println("Online~")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
