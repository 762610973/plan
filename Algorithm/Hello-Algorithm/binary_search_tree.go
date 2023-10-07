package Hello_Algorithm

func (bst *TreeNode) search(num int) *TreeNode {
	node := bst
	// 循环查找,越过叶节点后跳出
	for node != nil {
		if node.Val < num {
			// 目标节点在 cur 的右子树中
			node = node.Right
		} else if node.Val > num {
			// 目标节点在 cur 的左子树中
			node = node.Left
		} else {
			// 找到目标节点, 跳出循环
			break
		}
	}
	// 返回目标节点
	return node
}
