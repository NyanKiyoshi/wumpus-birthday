package mentions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var mentionCases = []struct {
	in  string
	out string
}{
	{"<@123>", "123"},
	{"<@!123>", "123"},
	{"<@!!123>", ""},
	{"<@!123>>", ""},
	{"<@!hello>", ""},
	{"hello", ""},
	{"", ""},
}

func TestMention(t *testing.T) {
	for _, tt := range mentionCases {
		t.Run(tt.in, func(t *testing.T) {
			assert.Equal(t, tt.out, Mention(tt.in))
		})
	}
}
