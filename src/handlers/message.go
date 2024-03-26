package handlers

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func message(s *discordgo.Session) string {
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		msg := strings.ToLower(m.Content)
		if strings.Contains(msg, "aide") {
			s.ChannelMessageSend(m.ChannelID, "import (\"Cherche toi même\")")
		}
		return string(m.Content)
	})
	return ""
}
