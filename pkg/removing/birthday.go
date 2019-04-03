package removing

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func Remove(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSend(m.ChannelID, "Removing..."); err != nil {
		log.Printf("Failed to send removal: %s", err)
	}
}
