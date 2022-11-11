package systemdesign

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	return isValidBST(root, nil, nil)
}

func isValidBST(n, r, l *TreeNode) bool {
	if n == nil {
		return true
	}

	if l != nil && l.Val > n.Val {
		return false
	}

	if r != nil && r.Val < n.Val {
		return false
	}

	return isValidBST(n.Left, l, n) && isValidBST(n.Right, n, r)
}
