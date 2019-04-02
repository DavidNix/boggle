package trie

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNode_InsertExists(t *testing.T) {
	trie := New()
	trie.Insert([]rune("cabbage"))
	trie.Insert([]rune("⛷"))

	for _, tt := range []struct {
		Word           string
		Prefix, Exists bool
	}{
		{"cabbage", true, true},
		{"cabx", false, false},
		{"cabb", true, false},
		{"cabbages", false, false},
		{"cabbag", true, false},
		{"", false, false},
		{"c", true, false},
		{"xoxo", false, false},
		{"⛷", true, true},
	} {
		prefix, ok := trie.Exists([]rune(tt.Word))
		require.Equal(t, tt.Prefix, prefix, "prefix for %s", tt.Word)
		require.Equal(t, tt.Exists, ok, "exists for %s", tt.Word)
	}
}
