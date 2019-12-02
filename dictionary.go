package boggle

import (
	"bufio"
	"os"
	"strings"

	"github.com/DavidNix/boggle/trie"
	"github.com/pkg/errors"
)

type Lang int

const (
	EnLang Lang = iota
)

type Dictionary struct {
	*trie.Node
	Language Lang
}

func NewDictionary() (Dictionary, error) {
	f, err := os.Open("en.txt")
	if err != nil {
		return Dictionary{}, errors.Wrap(err, "NewDictionary file")
	}
	defer f.Close()

	node := trie.New()

	sn := bufio.NewScanner(f)
	for sn.Scan() {
		node.Insert([]rune(sn.Text()))
	}
	if sn.Err() != nil {
		return Dictionary{}, errors.Wrap(sn.Err(), "NewDictionary scan")
	}

	return Dictionary{Node: node, Language: EnLang}, nil
}

func (d Dictionary) Exists(word []rune) (prefixExists, wordExists bool) {
	normalized := strings.ToLower(string(word))
	return d.Node.Exists([]rune(normalized))
}
