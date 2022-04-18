package tree

func Test() {
	var tree BinaryTree
	tree.CreateBinaryTree(4)
	err := tree.root.AddNode(5)
	if err != nil {
		return
	}
	err = tree.root.AddNode(10)
	if err != nil {
		return
	}
	err = tree.root.left.AddNode(20)
	if err != nil {
		return
	}
	println(tree.GetDepth(tree.root))
	tree.PreOrderTraverse(tree.root)
	tree.MidOrderTraverse(tree.root)
	tree.PostOrderTraverse(tree.root)
}
