package mentions

import (
	"github.com/asaskevich/govalidator"
	"strings"
)

var mentionSuffixes = []string{"<@!", "<@"}

// IsMention returns the ID of the mention, if any.
// Returns nothing otherwise.
func Mention(mention string) string {
	var userID string
	prefixPos := len(mention) - 1

	if !strings.HasSuffix(mention, ">") {
		return ""
	}

	for _, suffix := range mentionSuffixes {
		if strings.HasPrefix(mention, suffix) {
			userID = mention[len(suffix):prefixPos]
			break
		}
	}

	// Ensure the user id is valid, if not, return nothing
	if userID != "" && !govalidator.IsNumeric(userID) {
		return ""
	}
	return userID
}
