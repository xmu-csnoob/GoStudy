package linkedList

import "fmt"

type LinkedList struct {
	root *linkedListNode
	len  int
}
type linkedListNode struct {
	value int
	next  *linkedListNode
}

func (node *linkedListNode) New(value int) {
	node.value = value
}
func (linkedList *LinkedList) New(len int) {
	var node *linkedListNode
	linkedList.root.New(0)
	node = linkedList.root
	for i := 0; i < len; i++ {
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		node.value = input
		var newNode *linkedListNode
		node.next = newNode
		node = newNode
	}
}
