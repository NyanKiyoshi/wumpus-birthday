package removing

import (
	"github.com/bwmarrin/discordgo"
	"wumpus-birthday/pkg/storage/birthday"
)

func Remove(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := birthday.Remove(m.Author.ID)

	if err == nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Your birthday has been removed!")
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
}
