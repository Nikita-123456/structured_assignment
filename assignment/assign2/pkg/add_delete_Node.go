package pkg
import(
	constant "github.com/nikita/assign2/constants"
)


func Create(k int, v int) *constant.Node {
	n := constant.Node{}
	n.Key = k
	n.Value = v
	return &n
}

// add node at the front of the linked list
func AddNode(newNode *constant.Node) {
	temp := constant.Head.Next
	newNode.Next = temp
	newNode.Prev = constant.Head
	constant.Head.Next = newNode
	temp.Prev = newNode
}

// remove node from the linked list
func DeleteNode(delNode *constant.Node) {
	delPrev := delNode.Prev
	delNext := delNode.Next
	delPrev.Next = delNext
	delNext.Prev = delPrev

}
