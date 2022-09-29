package student

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == node {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		min := BTreeMin(root.Right)
		root.Data = min.Data
		root.Right = BTreeDeleteNode(root.Right, min)
		return root
	}
	if root.Left == node {
		root.Left = BTreeTransplant(root.Left, node, nil)
	} else {
		root.Right = BTreeTransplant(root.Right, node, nil)
	}
	return root
}