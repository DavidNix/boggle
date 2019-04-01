package boggle

type Vertex struct {
	Parent *Vertex
	Row, Col int
}

type Visitor interface {
	Visit(letters string) bool
}

type Board [][]string

func (b Board) Traverse(v Visitor) {
	for row := 0; row < len(b); row ++ {
		for col := 0; col < len(b[row]); col ++ {
			root := Vertex{
				Row: row,
				Col: col,
			}
			b.visit(&root, v, "")
		}
	}
}

var adjCoords = [][]int {
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

func (b Board) visit(vertex *Vertex, visitor Visitor, cum string) {
	letter := b[vertex.Row][vertex.Col]
	cum = cum + letter
	stop := visitor.Visit(cum)
	if stop {
		return
	}

	for _, coord := range adjCoords {
		vtx := &Vertex{Row: vertex.Row + coord[0], Col: vertex.Col + coord[1], Parent: vertex}
		if vtx.Row >= 0 && vtx.Row < len(b) && vtx.Col >= 0 && vtx.Col < len(b) && !visited(vertex, vtx) {
			b.visit(vtx, visitor, cum)
		}
	}
}

func visited(parent, v *Vertex) bool {
	if parent == nil {
		return false
	}
	if parent.Row == v.Row && parent.Col == v.Col {
		return true
	}
	return visited(parent.Parent, v)
}

