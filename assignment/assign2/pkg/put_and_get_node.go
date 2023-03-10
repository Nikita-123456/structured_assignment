package pkg
import(
	constant "github.com/nikita/assign2/constants"
)

// brings the corresponding value to the key if present else return -1
// if the key is present then bringing the node at the front of the list and storing new address in the map
func Get(k int) int {
	if _, ok := constant.MapIt[k]; ok {
		resNode := constant.MapIt[k]
		res := resNode.Value
		delete(constant.MapIt, k)
		DeleteNode(resNode)
		AddNode(resNode)
		constant.MapIt[k] = constant.Head.Next
		return res
	}
	return -1
}

// checks if the key is already present , if yes then delete it from the list and map
// if capacity is full then delete the last node
// at last storing the new node in list and map
func Put(k int, v int) {
	if _, ok := constant.MapIt[k]; ok {
		existingNode := constant.MapIt[k]
		delete(constant.MapIt, k)
		DeleteNode(existingNode)
	}
	if len(constant.MapIt) == constant.Cap {
		delete(constant.MapIt, constant.Tail.Prev.Key)
		DeleteNode(constant.Tail.Prev)
	}
	AddNode(Create(k, v))
	constant.MapIt[k] = constant.Head.Next
}
