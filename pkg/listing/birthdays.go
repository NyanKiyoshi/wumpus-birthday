package listing

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func List(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSend(m.ChannelID, "Listing..."); err != nil {
		log.Printf("Failed to send listing: %s", err)
	}
}
