package boggle

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

type visitor struct {
	Visited []string
	Stop func(s string) bool
}

func (v *visitor) Visit(letters string) bool {
	v.Visited = append(v.Visited, letters)
	if v.Stop == nil {
		return false
	}
	return v.Stop(letters)
}

func TestBoard_Traverse(t *testing.T) {
	board := Board{
		{"A", "B"},
		{"C", "D"},
	}
	v := &visitor{}

	board.Traverse(v)

	expected := "A AB ABD ABDC ABC ABCD AD ADB ADBC ADC ADCB AC ACB ACBD ACD ACDB"

	for _, val := range strings.Split(expected, " "){
		require.Contains(t, v.Visited, val)
	}

	require.Len(t, v.Visited, 64)

	for i := range v.Visited {
		letters := v.Visited[i]

		require.NotEmpty(t, letters)
		require.True(t, len(letters) <= 4)
	}
}

func TestBoard_Traverse_stop(t *testing.T) {
	board := Board{
		{"A", "B"},
		{"C", "D"},
	}
	v := &visitor{Stop: func(s string) bool {
		return len(s) == 1
	}}
	board.Traverse(v)

	require.Equal(t, "ABCD", strings.Join(v.Visited, ""))
}
