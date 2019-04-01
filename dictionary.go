package boggle

import (
	"bufio"
	"github.com/pkg/errors"
	"os"
)

type Dictionary map[string] struct{}

func NewDictionary() (Dictionary, error) {
	f, err := os.Open("words_preprocessed.txt")
	if err != nil {
		return nil, errors.Wrap(err, "NewDictionary")
	}
	defer f.Close()

	dict := make(Dictionary)
	sn := bufio.NewScanner(f)
	for sn.Scan() {
		dict[sn.Text()] = struct{}{}
	}
	return dict, errors.Wrap(sn.Err(), "NewDictionary")
}

func (dict Dictionary) Exists(word string) bool {
	_, ok := dict[word]
	return ok
}
