package notifying

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

const helpText = `
Invalid command. Usage: @bot sentences add|remove sentence.
Example: @bot sentences add Happy birthday, %s!!`

func Dispatch(
	s *discordgo.Session, m *discordgo.Message,
	commands []string) string {

	if len(commands) < 1 {
		return helpText
	}

	switch commands[1] {
	case "add":
		addSentence(s, m, strings.Join(commands[1:], " "))
		break
	case "remove":
		removeSentence(s, m, strings.Join(commands[1:], " "))
		break
	default:
		return helpText
	}

	return ""
}
