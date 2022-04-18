package tree

type BinaryTree struct {
	root *BinaryTreeNode
}

func (binaryTree *BinaryTree) CreateBinaryTree(value int) {
	var root *BinaryTreeNode
	root = &BinaryTreeNode{
		value: value,
		left:  new(BinaryTreeNode),
		right: new(BinaryTreeNode),
		init:  true,
	}
	binaryTree.root = root
}
func (binaryTree *BinaryTree) GetDepth(node *BinaryTreeNode) int {
	var left, right int
	if node == nil {
		return 0
	}
	if node.left.init {
		left = binaryTree.GetDepth(node.left)
	}
	if node.right.init {
		right = binaryTree.GetDepth(node.right)
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func (binaryTree *BinaryTree) PreOrderTraverse(node *BinaryTreeNode) {
	if !node.init {
		return
	}
	println(node.value)
	binaryTree.PreOrderTraverse(node.left)
	binaryTree.PreOrderTraverse(node.right)
}
func (binaryTree *BinaryTree) MidOrderTraverse(node *BinaryTreeNode) {
	if !node.init {
		return
	}
	binaryTree.MidOrderTraverse(node.left)
	println(node.value)
	binaryTree.MidOrderTraverse(node.right)
}
func (binaryTree *BinaryTree) PostOrderTraverse(node *BinaryTreeNode) {
	if !node.init {
		return
	}
	binaryTree.PostOrderTraverse(node.left)
	binaryTree.PostOrderTraverse(node.right)
	println(node.value)
}
