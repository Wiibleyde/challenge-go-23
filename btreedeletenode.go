package student

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node.Left == nil {
		return BTreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		return BTreeTransplant(root, node, node.Left)
	} else {
		min := BTreeMin(node.Right)
		if min.Parent != node {
			BTreeTransplant(root, min, min.Right)
			min.Right = node.Right
			min.Right.Parent = min
		}
		BTreeTransplant(root, node, min)
		min.Left = node.Left
		min.Left.Parent = min
		return root
	}
}
