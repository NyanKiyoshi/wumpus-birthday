package permissions // isModerator returns whether or not the sender of a message is a moderator.
import (
	"github.com/bwmarrin/discordgo"
	"log"
)

// IsAuthorModerator returns whether or not the given user is a moderator.
func IsUserModerator(s *discordgo.Session, userID string, channelID string) bool {
	p, err := s.UserChannelPermissions(userID, channelID)
	if err == nil {
		return p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator ||
			p&discordgo.PermissionManageChannels == discordgo.PermissionManageChannels ||
			p&discordgo.PermissionManageServer == discordgo.PermissionManageServer
	}

	log.Printf(
		"failed to get user permissions for %s @ <#%s>: %s",
		userID, channelID, err)
	return false
}

// IsAuthorModerator returns whether or not the sender of a message is a moderator.
func IsAuthorModerator(s *discordgo.Session, message *discordgo.Message) bool {
	return IsUserModerator(s, message.Author.ID, message.ChannelID)
}
