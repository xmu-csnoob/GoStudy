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
	}
	binaryTree.root = root
}
