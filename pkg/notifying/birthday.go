package notifying

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
	"wumpus-birthday/pkg/storage/birthday"
)

const BirthdayChannelName = "birthdays"

func getNotificationChannelFromServerID(
	s *discordgo.Session, cache map[string]string, serverID string) (string, error) {

	channels, err := s.GuildChannels(serverID)
	for _, ch := range channels {
		if strings.ToLower(ch.Name) == BirthdayChannelName {
			cache[serverID] = ch.ID
		}
	}
	return "", err
}

func notifyAll(s *discordgo.Session) error {
	birthdayChannels := make(map[string]string)
	birthdays, err := birthday.TodayAllServers()

	if err != nil {
		return err
	}

	for _, date := range birthdays {
		channelID, errGet := getNotificationChannelFromServerID(
			s, birthdayChannels, date.ServerID)

		if errGet != nil {
			log.Printf(
				"Failed to get birthday channel of %s: %s", date.ServerID, errGet)
			continue
		}

		if channelID == "" {
			continue
		}

		_, _ = s.ChannelMessageSend(
			channelID, fmt.Sprintf("Happy birthday! <@!%s>", date.UserID))
	}

	return nil
}

func notifyServer(s *discordgo.Session) {

}

func waitOnce() <-chan time.Time {
	now := time.Now()
	tomorrowDiff := now.Sub(time.Date(
		now.Year(), now.Month(), now.Day(), 0, 1, 0, 0, now.Location()))
	return time.After(tomorrowDiff)
}

// WaitForEver waits forever for birthdays to notify.
func WaitForEver(s *discordgo.Session, stop chan bool) {
	for {
		select {
		case <-waitOnce():
			notifyAll(s)
			break
		case <-stop:
			return
		}
	}
}
