package Hello_Algorithm

import "container/list"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 层序遍历
func levelOrder(root *TreeNode) []any {
	// 初始化队列，加入根节点
	queue := list.New()
	queue.PushBack(root)
	// 初始化一个切片，用于保存遍历序列
	nums := make([]any, 0)
	for queue.Len() > 0 {
		// 队列出队
		node := queue.Remove(queue.Front()).(*TreeNode)
		// 保存节点值
		nums = append(nums, node.Val)
		// 将此节点的左右子节点加入队列中
		if node.Left != nil {
			// 左子节点入队
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			// 右子节点入队
			queue.PushBack(node.Right)
		}
	}
	return nums
}
