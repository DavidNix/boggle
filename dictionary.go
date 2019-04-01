package boggle

import (
	"bufio"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type Dictionary map[string] struct{}

func NewDictionary() (Dictionary, error) {
	f, err := os.Open("en.txt")
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
	_, ok := dict[strings.ToLower(word)]
	return ok
}
