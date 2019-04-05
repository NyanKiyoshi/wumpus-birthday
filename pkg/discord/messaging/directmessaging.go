package messaging

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func sendDM(s *discordgo.Session, userID string, message *string) error {
	channel, err := s.UserChannelCreate(userID)

	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSend(channel.ID, *message)
	return err
}

func ReplyDM(s *discordgo.Session, source *discordgo.Message, message *string) {
	if err := sendDM(s, source.Author.ID, message); err != nil {
		_, _ = s.ChannelMessageSend(
			source.ChannelID, fmt.Sprintf("Tried to send you a DM, but got: `%s`", err))
		return
	}

	_, _ = s.ChannelMessageSend(source.ChannelID, "Sent you a DM!")
}
