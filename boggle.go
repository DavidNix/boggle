package boggle

import (
	"fmt"
	"sync"
)

type Coordinate []int

func (c Coordinate) Row() int { return c[0] }
func (c Coordinate) Col() int { return c[1] }
func (c Coordinate) String() string {return fmt.Sprintf("%d-%d", c.Row(), c.Col())}

type BoardNode struct {
	Parent *BoardNode
	Row, Col int

	length int
}

func (node *BoardNode) Path() []Coordinate {
	path := make([]Coordinate, node.length)
	node.buildPath(node.length-1, path)
	return path
}

func (node *BoardNode) buildPath(idx int, path []Coordinate) {
	if node == nil {
		return
	}
	path[idx] = Coordinate{node.Row, node.Col}
	idx--
	node.Parent.buildPath(idx, path)
}

type Visitor interface {
	Visit(node *BoardNode, letters string) bool
}

type ConcurrentVisitor interface {
	Visitor
	Done()
}

type Board [][]string

func (b Board) Traverse(v Visitor) {
	for row := 0; row < len(b); row ++ {
		for col := 0; col < len(b[row]); col ++ {
				root := BoardNode{
					Row: row,
					Col: col,
					length: 1,
				}
				b.visit(&root, v, "")
		}
	}
}

func (b Board) TraverseConcurrent(v ConcurrentVisitor) {
	var wg sync.WaitGroup
	defer v.Done()
	for row := 0; row < len(b); row ++ {
		for col := 0; col < len(b[row]); col ++ {
			wg.Add(1)
			go func(row, col int) {
				root := BoardNode{
					Row: row,
					Col: col,
					length: 1,
				}
				b.visit(&root, v, "")
				wg.Done()
			}(row, col)
		}
	}
	wg.Wait()
}

var adjCoords = []Coordinate {
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

func (b Board) visit(node *BoardNode, visitor Visitor, cum string) {
	letter := b[node.Row][node.Col]
	cum = cum + letter
	stop := visitor.Visit(node, cum)
	if stop {
		return
	}

	for _, coord := range adjCoords {
		child := &BoardNode{Row: node.Row + coord.Row(), Col: node.Col + coord.Col(), Parent: node, length: node.length+1}
		if child.Row >= 0 && child.Row < len(b) && child.Col >= 0 && child.Col < len(b) && !visited(node, child) {
			b.visit(child, visitor, cum)
		}
	}
}

func visited(parent, node *BoardNode) bool {
	if parent == nil {
		return false
	}
	if parent.Row == node.Row && parent.Col == node.Col {
		return true
	}
	return visited(parent.Parent, node)
}

