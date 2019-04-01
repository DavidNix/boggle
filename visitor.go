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

func (wf *WordFinder) Visit(node *Node, letters string) bool {
	if len(letters) < 3 {
		return false
	}
	if wf.dict.Exists(letters) {
		wf.Found = append(wf.Found, Entry{Word: letters, Path: node.Path()})
	}
	return false
}



