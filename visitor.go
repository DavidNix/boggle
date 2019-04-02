package boggle

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
