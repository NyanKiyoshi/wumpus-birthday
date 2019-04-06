package listing

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
	"wumpus-birthday/pkg/discord/messaging"
	"wumpus-birthday/pkg/storage/birthday"
)

func formatDate(t time.Time) string {
	suffix := "th"

	switch t.Day() {
	case 1, 21, 31:
		suffix = "st"
	case 2, 22:
		suffix = "nd"
	case 3, 23:
		suffix = "rd"
	}

	return t.Format("Monday 2" + suffix + " January 2006")
}

func sendList(
	s *discordgo.Session, m *discordgo.MessageCreate,
	birthdays []birthday.Birthday, err error) {

	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	if len(birthdays) == 0 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "None.")
		return
	}

	var msg = fmt.Sprintf("As of %s\n", formatDate(time.Now()))

	for _, day := range birthdays {
		msg += fmt.Sprintf("<@!%s> on %s\n", day.UserID, formatDate(day.Date))
	}

	messaging.ReplyDM(s, m.Message, &msg)
}

func List(s *discordgo.Session, m *discordgo.MessageCreate) {
	birthdays, err := birthday.GetAll(m.GuildID)
	sendList(s, m, birthdays, err)
}

func ListToday(s *discordgo.Session, m *discordgo.MessageCreate) {
	birthdays, err := birthday.Today(m.GuildID)
	sendList(s, m, birthdays, err)
}
