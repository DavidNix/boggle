package boggle

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDictionary_Exists(t *testing.T) {
	dict, err := NewDictionary()
	require.NoError(t, err)

	for _, pref := range []string{"catt", "aardv", "UMBR", "ChEES"} {
		prefix, exists := dict.Exists(pref)
		require.True(t, prefix, pref)
		require.False(t, exists, pref)
	}

	for _, word := range []string{"cat", "dog", "aardvark", "UMBRELLA", "ChEEse"} {
		prefix, exists := dict.Exists(word)
		require.True(t, prefix, word)
		require.True(t, exists, word)
	}

	for _, word := range []string{"78393", "billyjean", "Octember", "rstudio", "locutus"} {
		prefix, exists := dict.Exists(word)
		require.False(t, prefix, word)
		require.False(t, exists, word)
	}
}

func BenchmarkDictionary_Exists(b *testing.B) {
	/*
	Builtin Map
	goos: darwin
	goarch: amd64
	pkg: github.com/DavidNix/boggle
	BenchmarkDictionary_Exists-12    	100000000	        13.7 ns/op

	Trie
	goos: darwin
	goarch: amd64
	pkg: github.com/DavidNix/boggle
	BenchmarkDictionary_Exists-12    	10000000	       112 ns/op
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
