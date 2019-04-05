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

func replyDM(s *discordgo.Session, source *discordgo.Message, message *string) error {
	channel, err := s.UserChannelCreate(source.Author.ID)

	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSend(channel.ID, *message)

	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSend(source.ChannelID, "Sent you a DM!")
	return err
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

	if err = replyDM(s, m.Message, &msg); err != nil {
		_, _ = s.ChannelMessageSend(
			m.ChannelID, fmt.Sprintf("Tried to send you a DM, but got: `%s`", err))
		return
	}
}

func List(s *discordgo.Session, m *discordgo.MessageCreate) {
	birthdays, err := birthday.GetAll()
	sendList(s, m, birthdays, err)
}

func ListToday(s *discordgo.Session, m *discordgo.MessageCreate) {
	birthdays, err := birthday.Today()
	sendList(s, m, birthdays, err)
}
