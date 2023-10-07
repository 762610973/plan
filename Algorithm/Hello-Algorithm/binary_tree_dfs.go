package Hello_Algorithm

var nums []int

/* 前序遍历 */
func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	// 根节点
	nums = append(nums, node.Val)
	// 左子树
	preOrder(node.Left)
	// 右子树
	preOrder(node.Right)
}

/* 中序遍历 */
func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	// 左子树
	inOrder(node.Left)
	// 根节点
	nums = append(nums, node.Val)
	// 右子树
	inOrder(node.Right)
}

/* 后序遍历 */
func postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	// 左子树
	postOrder(node.Left)
	// 右子树
	postOrder(node.Right)
	// 根节点
	nums = append(nums, node.Val)
}
