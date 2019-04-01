package boggle

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDictionary_Exists(t *testing.T) {
	dict, err := NewDictionary()
	require.NoError(t, err)

	for _, word := range []string{"cat", "dog", "aardvark", "UMBRELLA", "ChEEse"} {
		require.True(t, dict.Exists(word), word)
	}

	for _, word := range []string{"78393", "billyjean", "Octember", "rstudio", "locutus"} {
		require.False(t, dict.Exists(word), word)
	}
}

func BenchmarkDictionary_Exists(b *testing.B) {
	/*
	goos: darwin
	goarch: amd64
	pkg: github.com/DavidNix/boggle
	BenchmarkDictionary_Exists-12    	100000000	        13.7 ns/op
	 */

	dict, err := NewDictionary()
	if err != nil {
		b.Fail()
	}

	for i := 0; i < b.N; i++ {
		dict.Exists("dog")
		dict.Exists("not a word")
	}
}
