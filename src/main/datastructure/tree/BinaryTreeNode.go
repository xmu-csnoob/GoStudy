package tree

import "errors"

type BinaryTreeNode struct {
	value int
	init  bool
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (binaryTreeNode *BinaryTreeNode) AddNode(value int) error {
	if !binaryTreeNode.left.init {
		binaryTreeNode.left = &BinaryTreeNode{
			value: value,
			init:  true,
			left:  new(BinaryTreeNode),
			right: new(BinaryTreeNode),
		}
		return nil
	}
	if !binaryTreeNode.right.init {
		binaryTreeNode.right = &BinaryTreeNode{
			value: value,
			init:  true,
			left:  new(BinaryTreeNode),
			right: new(BinaryTreeNode),
		}
		return nil
	}
	return errors.New("此节点的两个子节点已经初始化")
}
