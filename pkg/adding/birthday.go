package adding

import (
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/bwmarrin/discordgo"
	"strings"
	"wumpus-birthday/pkg/discord/mentions"
	"wumpus-birthday/pkg/discord/permissions"
)

func addDate(userID string, date string) string {
	parsedDate, err := dateparse.ParseAny(date)

	if err == nil {
		return fmt.Sprint("Added: ", parsedDate.String())
	}

	return err.Error()
}

func addDateFromMention(
	s *discordgo.Session, m *discordgo.Message, mention string, date string) string {

		// Attempt to convert the mention to a user id
		userID := mentions.Mention(mention)

		// If the user id is invalid, do nothing
		if userID == "" {
			return ""
		}

		// If the user id was passed, ensure the user is a moderator
		if permissions.IsAuthorModerator(s, m) {
			return addDate(userID, date)
		}
		return "Not enough privileges to add someone's birthday."
}

func parseMessage(
	s *discordgo.Session, m *discordgo.Message, mention string, date string) string {

		if mention != "" {
			userID := mentions.Mention(mention)
			if userID != "" {
				return addDateFromMention(s, m, mention, date)
			}
		}

		// Continue parsing
		return addDate(m.Author.ID, mention +" "+date)
}

func Add(s *discordgo.Session, m *discordgo.MessageCreate, commands []string) {
	var msg string
	argc := len(commands)

	if argc == 0 {
		msg = "Please provide a date."
	} else if argc > 1 {
		msg = parseMessage(s, m.Message, commands[0], strings.Join(commands[1:], " "))
	} else {
		msg = parseMessage(s, m.Message, "", commands[0])
	}

	if msg != "" {
		_, _ = s.ChannelMessageSend(m.ChannelID, msg)
	}
}
