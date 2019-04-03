package boggle

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordFinder_Visit(t *testing.T) {
	en, err := NewDictionary()
	require.NoError(t, err)
	v := NewVisitor(en)
	board := Board{
		{"G", "O"},
		{"D", "O"},
	}
	board.Traverse(v)

	require.NotEmpty(t, v.Found)

	expected := []string{"GOO", "GOOD", "GOD", "DOG", "DOO"}
	for _, found := range v.Found {
		require.Contains(t, expected, found.Word)
	}
}

// source: http://fuzzylogicinc.net/boggle/EnterBoard.aspx
const answers = `ben bens engirt ens gen gens get girn girns girt git gite gite grit neb nebs net new news newt rig rit rite rites seg sen set sew sewn sneb spiv teg ten tens tes tew tews tig tige tiges trig twp venge venges vibe vibes vibs wen wens wet`

func TestWordFinder_Visit_4x4(t *testing.T) {
	en, err := NewDictionary()
	require.NoError(t, err)
	v := NewVisitor(en)
	board := Board{
		{"I", "R", "N", "S"},
		{"T", "G", "N", "E"},
		{"E", "W", "B", "V"},
		{"N", "S", "P", "I"},
	}
	board.Traverse(v)

	require.NotEmpty(t, v.Found)

	found := make(map[string]struct{})
	for _, f := range v.Found {
		found[f.Word] = struct{}{}
	}

	var missing []string
	for _, word := range strings.Split(answers, " ") {
		_, ok := found[strings.ToUpper(word)]
		if !ok {
			missing = append(missing, word)
		}
	}

	require.Empty(t, missing)
}

func TestWordFinder_Visit_4x4_concurrent(t *testing.T) {
	en, err := NewDictionary()
	require.NoError(t, err)
	v := NewConcurrentVisitor(en)

	board := Board{
		{"I", "R", "N", "S"},
		{"T", "G", "N", "E"},
		{"E", "W", "B", "V"},
		{"N", "S", "P", "I"},
	}
	go board.TraverseConcurrent(v)

	found := make(map[string]struct{})
	for entry := range v.Entries {
		found[entry.Word] = struct{}{}
	}

	var missing []string
	for _, word := range strings.Split(answers, " ") {
		_, ok := found[strings.ToUpper(word)]
		if !ok {
			missing = append(missing, word)
		}
	}

	require.Empty(t, missing)
}

func TestWordFinder_Visit_letterQ(t *testing.T) {
	en, err := NewDictionary()
	require.NoError(t, err)
	v := NewVisitor(en)
	board := Board{
		{"Q", "I", "C", "K"},
		{"X", "X", "X", "X"},
		{"X", "X", "X", "X"},
		{"X", "X", "X", "X"},
	}
	board.Traverse(v)

	require.NotEmpty(t, v.Found)
	for _, f := range v.Found {
		if strings.ToUpper(f.Word) == "QUICK" {
			return
		}
	}
	require.Fail(t, "did not find word: QUICK")
}
