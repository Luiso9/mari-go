package commands

import (
    "github.com/bwmarrin/discordgo"
    "strings"
    "github.com/Luiso9/mari-go/config"
)

type Command struct {
    Name        string
    Description string
    Handler     func(s *discordgo.Session, m *discordgo.MessageCreate)
}

var commandHandlers = make(map[string]Command)

func RegisterCommands(s *discordgo.Session) {
    commands := []Command{
        {"ping", "A simple ping command", pingHandler},
        {"avatar", "A simple avatar command", avatar},
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

    conf := config.GetConfig()
    if !strings.HasPrefix(m.Content, conf.Prefix) {
        return
    }

    commandStr := strings.TrimPrefix(m.Content, conf.Prefix)
    cmd, ok := commandHandlers[commandStr]
    if ok {
        cmd.Handler(s, m)
    }
}