package boggle

import (
	"bufio"
	"os"
	"strings"

	"github.com/DavidNix/boggle/trie"
	"github.com/pkg/errors"
)

type Dictionary struct {
	*trie.Node
}

func NewDictionary() (Dictionary, error) {
	f, err := os.Open("en.txt")
	if err != nil {
		return Dictionary{}, errors.Wrap(err, "NewDictionary file")
	}
	defer f.Close()

	trie := trie.New()

	sn := bufio.NewScanner(f)
	for sn.Scan() {
		trie.Insert([]rune(sn.Text()))
	}
	if sn.Err() != nil {
		return Dictionary{}, errors.Wrap(sn.Err(), "NewDictionary scan")
	}

	return Dictionary{Node: trie}, nil
}

func (d Dictionary) Exists(word []rune) (prefixExists, wordExists bool) {
	normalized := strings.ToLower(string(word))
	return d.Node.Exists([]rune(normalized))
}
