package adding

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

// isModerator returns whether or not the sender of a message is a moderator.
func isModerator(s *discordgo.Session, message *discordgo.Message) bool {
	p, err := s.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if err == nil {
		return p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator ||
			p&discordgo.PermissionManageChannels == discordgo.PermissionManageChannels ||
			p&discordgo.PermissionManageServer == discordgo.PermissionManageServer
	}

	log.Printf(
		"failed to get user permissions for %s @ <#%s>: %s",
		message.Author.String(), message.ChannelID, err)
	return false
}

func Add(s *discordgo.Session, m *discordgo.MessageCreate) {

}
