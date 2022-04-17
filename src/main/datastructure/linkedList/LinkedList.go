package linkedList

type LinkedList struct {
	root *linkedListNode
	len  int
}

func (linkedList *LinkedList) CreateLinkedList(rootValue int) {
	linkedList.root = &linkedListNode{
		value: rootValue,
		next:  new(linkedListNode),
	}
	linkedList.len = 1
}
func (linkedList *LinkedList) GetLength() int {
	return linkedList.len
}
func (linkedList *LinkedList) AddNode(value int) {
	var rear *linkedListNode
	rear = linkedList.root
	for i := 0; i < linkedList.len-1; i++ {
		rear = rear.next
	}
	node := &linkedListNode{
		value: value,
		next:  new(linkedListNode),
	}
	rear.next = node
	linkedList.len++
}
func (linkedList *LinkedList) Assign(index int, value int) {
	var node = linkedList.root
	for i := 0; i < index; i++ {
		node = node.next
	}
	node.value = value
}
func (linkedList *LinkedList) Traverse() {
	print(linkedList.root.value)
	var node *linkedListNode
	node = linkedList.root
	for i := 0; i < linkedList.len-1; i++ {
		node = node.next
		print("->", node.value)
	}
	println()
}
func Test() {
	var linkedList LinkedList
	linkedList.CreateLinkedList(10)
	linkedList.AddNode(20)
	linkedList.AddNode(30)
	linkedList.Traverse()
	linkedList.Assign(1, 50)
	linkedList.Traverse()
}
