package notifying

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func notifyAll(s *discordgo.Session) {

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
