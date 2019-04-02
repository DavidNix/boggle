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
	if len(letters) < 3 {
		return false
	}
	prefix, isWord := wf.dict.Exists(letters)
	if !prefix {
		return true
	}
	if isWord {
		wf.Found = append(wf.Found, Entry{Word: letters, Path: node.Path()})
	}
	return false
}

type ConcurrentFinder struct {
	dict Dictionary

	Entries chan Entry
}

func NewConcurrentVisitor(dict Dictionary) *ConcurrentFinder {
	return &ConcurrentFinder{
		dict: dict,
		Entries: make(chan Entry, 1000),
	}
}

func (cf *ConcurrentFinder) Visit(node *BoardNode, letters string) bool {
	if len(letters) < 3 {
		return false
	}
	prefix, isWord := cf.dict.Exists(letters)
	if !prefix {
		return true
	}
	if isWord {
		cf.Entries <- Entry{Word: letters, Path: node.Path()}
	}
	return false
}

func (cf *ConcurrentFinder) Done() {
	close(cf.Entries)
}
