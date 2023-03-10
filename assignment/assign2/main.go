// ------Implement LRU cache-----
package main

import (
	"fmt"

	constant "github.com/nikita/assign2/constants"
	pkg "github.com/nikita/assign2/pkg"
)

func main() {

	constant.MapIt = make(map[int]*constant.Node)

	//-------------------------------------------------
	//creating 2 nodes head and tail and making them point to each other
	//nil<---head<====>tail--->nil

	constant.Head = pkg.Create(-1, -1)
	constant.Tail = pkg.Create(-1, -1)

	constant.Head.Next = constant.Tail
	constant.Head.Prev = nil

	constant.Tail.Prev = constant.Head
	constant.Tail.Next = nil
	//-------------------------------------------------

	pkg.Put(1, 10)          // nil, linked list: [1:10]
	pkg.Put(2, 20)          // nil, linked list: [2:20,1:10]
	fmt.Println(pkg.Get(1)) // 10, linked list: [1:10,2:20]
	pkg.Put(3, 30)          // nil, linked list: [3:30,1:10]
	fmt.Println(pkg.Get(2)) // -1, linked list: [3:30,1:10]
	pkg.Put(4, 40)          // nil, linked list: [4:40,3:30]
	fmt.Println(pkg.Get(1)) // -1, linked list: [4:40,3:30]
	fmt.Println(pkg.Get(3)) // 30, linked list: [3:30,4:40]

}

