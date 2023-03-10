package constant


type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}
const Cap = 2

var MapIt map[int]*Node
var Head *Node
var Tail *Node

