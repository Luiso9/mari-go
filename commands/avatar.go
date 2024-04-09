package commands

import (
	"github.com/bwmarrin/discordgo"
)

func avatar(s *discordgo.Session, m *discordgo.MessageCreate) {
	userID := m.Author.ID
	user, err := s.User(userID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Terjadi kesalahan")
		return
	}

	avatarURL := "https://cdn.discordapp.com/avatars/" + userID + "/" + user.Avatar + ".png?size=256"

	s.ChannelMessageSend(m.ChannelID, avatarURL)
}
