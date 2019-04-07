package pkg

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"wumpus-birthday/pkg/adding"
	"wumpus-birthday/pkg/listing"
	"wumpus-birthday/pkg/notifying"
	"wumpus-birthday/pkg/removing"
)

func Register(s *discordgo.Session) {
	s.AddHandler(messageCreate)
}

func messagePrefix(s *discordgo.Session, m string) (bool, string) {
	for _, prefix := range []string{
		fmt.Sprintf("<@%s>", s.State.User.ID),
		fmt.Sprintf("<@!%s>", s.State.User.ID)} {

		if strings.HasPrefix(m, prefix) {
			return true, prefix
		}
	}

	return false, ""
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// If the message is in DM then ignore the message
	if m.GuildID == "" {
		return
	}

	// Check the bot is mentioned. If not, ignore the message
	hasPrefix, prefix := messagePrefix(s, m.Content)
	if !hasPrefix {
		return
	}

	message := strings.TrimLeft(m.Content[len(prefix):], " ")
	commands := strings.Fields(message)

	switch commands[0] {
	case "remove":
		removing.Remove(s, m)
		return
	case "list":
		listing.List(s, m)
		return
	case "today":
		listing.ListToday(s, m)
		return
	case "sentences":
		if err := notifying.Dispatch(s, m.Message, commands[1:]); err != "" {
			_, _ = s.ChannelMessageSend(m.ChannelID, err)
		}
		return
	case "help":
		_, _ = s.ChannelMessageSend(m.ChannelID, helpText)
		return
	default:
		adding.Add(s, m, commands)
	}
}
