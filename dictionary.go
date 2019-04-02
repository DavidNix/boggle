package boggle

import (
	"bufio"
	"github.com/DavidNix/boggle/trie"
	"github.com/pkg/errors"
	"os"
	"strings"
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
		trie.Insert(sn.Text())
	}
	if sn.Err() != nil {
		return Dictionary{}, errors.Wrap(sn.Err(), "NewDictionary scan")
	}

	return Dictionary{Node: trie}, nil
}

func (d Dictionary) Exists(word string) (prefixExists, wordExists bool) {
	return d.Node.Exists(strings.ToLower(word))
}
