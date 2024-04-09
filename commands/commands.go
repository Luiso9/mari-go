package commands

import (
    "github.com/bwmarrin/discordgo"
)

// Define a command with a name, description, and a handler function.
type Command struct {
    Name        string
    Description string
    Handler     func(s *discordgo.Session, m *discordgo.MessageCreate)
}

// A map to store registered commands
var commandHandlers = make(map[string]Command)

func RegisterCommands(s *discordgo.Session) {
    commands := []Command{
        {"ping", "A simple ping command", pingHandler},
    }

    for _, cmd := range commands {
        commandHandlers[cmd.Name] = cmd
        s.AddHandler(messageCreateHandler)
    }
}

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Author.ID == s.State.User.ID {
        return
    }

    cmd, ok := commandHandlers[m.Content] 
    if ok {
        cmd.Handler(s, m) 
    }
}