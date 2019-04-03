package pkg

import (
	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

const botUserID = "123"

var mentionCases = []struct {
	matches bool
	message string
	prefix  string
}{
	// Mention two users, including the bot, but in wrong order
	{
		matches: false,
		message: "<@!231> <@123>",
		prefix:  "",
	},

	// Mention only the bot
	{
		matches: true,
		message: "<@!123> hello",
		prefix:  "<@!123>",
	},

	// Mention only the bot (without `!`)
	{
		matches: true,
		message: "<@123>",
		prefix:  "<@123>",
	},

	// Mention nobody
	{
		matches: false,
		message: "<#123>",
		prefix:  "",
	},
}

func Test_hasCommandPrefix(t *testing.T) {
	s := &discordgo.Session{
		StateEnabled: true,
		State:        discordgo.NewState(),
	}

	s.State.User = &discordgo.User{
		ID: botUserID,
	}

	for _, tt := range mentionCases {
		t.Run(tt.message, func(t *testing.T) {
			isMatch, prefix := messagePrefix(s, tt.message)

			assert.Equal(
				t, tt.matches, isMatch,
				"It did not correctly detected the mention, expected to match? %t",
				tt.matches,
			)
			assert.Equal(
				t, tt.prefix, prefix,
				"Returned an invalid prefix",
			)
		})
	}
}

func TestRegister(t *testing.T) {
}
