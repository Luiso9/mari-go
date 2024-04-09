package commands

import "github.com/bwmarrin/discordgo"

func pingHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    s.ChannelMessageSend(m.ChannelID, "Pong!")
}