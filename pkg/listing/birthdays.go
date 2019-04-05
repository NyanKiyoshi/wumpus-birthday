package listing

import (
	"github.com/bwmarrin/discordgo"
	"time"
	"wumpus-birthday/pkg/storage"
)

func List(s *discordgo.Session, m *discordgo.MessageCreate) {
	birthdays, err := storage.Get(time.Now())
	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, string(len(birthdays)))
}
