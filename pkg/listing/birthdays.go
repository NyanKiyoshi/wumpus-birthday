package listing

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
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

	var msg string

	if len(birthdays) == 0 {
		msg = "None."
	}
	for _, day := range birthdays {
		msg += fmt.Sprintf("<@!%s> on %s\n", day.UserID, formatDate(day.Date))
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, msg)
}

func List(s *discordgo.Session, m *discordgo.MessageCreate) {
	birthdays, err := birthday.GetAll()
	sendList(s, m, birthdays, err)
}

func ListToday(s *discordgo.Session, m *discordgo.MessageCreate) {
	birthdays, err := birthday.Today()
	sendList(s, m, birthdays, err)
}
