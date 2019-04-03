package boggle

import (
	"unicode"
)

type Entry struct {
	Word string
	Path []Coordinate
}

type WordFinder struct {
	dict Dictionary

	Found []Entry
}

func NewVisitor(dict Dictionary) *WordFinder {
	return &WordFinder{dict: dict}
}

func (wf *WordFinder) Visit(node *BoardNode, letters string) bool {
	entry, stop := findWord(node, wf.dict, []rune(letters))
	if entry.Word != "" {
		wf.Found = append(wf.Found, entry)
	}
	return stop
}

type ConcurrentFinder struct {
	dict Dictionary

	Entries chan Entry
}

func NewConcurrentVisitor(dict Dictionary) *ConcurrentFinder {
	return &ConcurrentFinder{
		dict:    dict,
		Entries: make(chan Entry, 1000),
	}
}

func (cf *ConcurrentFinder) Visit(node *BoardNode, letters string) bool {
	entry, stop := findWord(node, cf.dict, []rune(letters))
	if entry.Word != "" {
		cf.Entries <- entry
	}
	return stop
}

func (cf *ConcurrentFinder) Done() {
	close(cf.Entries)
}

func findWord(node *BoardNode, dict Dictionary, letters []rune) (Entry, bool) {
	letters = languageMods(dict, letters)
	if len(letters) < 3 {
		return Entry{}, false
	}
	prefix, isWord := dict.Exists(letters)
	if !prefix {
		return Entry{}, true
	}
	if isWord {
		return Entry{Word: string(letters), Path: node.Path()}, false
	}
	return Entry{}, false
}

func languageMods(dict Dictionary, letters []rune) []rune {
	switch dict.Language {
	case EnLang:
		q := []rune("q")[0]
		const offset = 1
		for i := offset; i < len(letters); i++ {
			prev := letters[i-1]
			if unicode.ToLower(prev) == q {
				newLetters := append(letters[:i], append([]rune("u"), letters[i:]...)...)
				return newLetters
			}
		}
		return letters
	default:
		return letters
	}
}
