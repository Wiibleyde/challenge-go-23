package student

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == node {
		return rplc
	}
	if root.Left == node {
		root.Left = rplc
	} else {
		root.Right = rplc
	}
	return root
}
