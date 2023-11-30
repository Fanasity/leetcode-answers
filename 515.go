/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	leaves := make([]*TreeNode, 0)
	leaves = append(leaves, root)

	for {
		if len(leaves) == 0 {
			break
		}
		res = append(res, getChildrenMaxNum(leaves))
		leaves = getLeavesChildren(leaves)
	}

	return res
}
func getLeavesChildren(leaves []*TreeNode) []*TreeNode {
	var res []*TreeNode
	for i := range leaves {
		if leaves[i].Left != nil {
			res = append(res, leaves[i].Left)
		}
		if leaves[i].Right != nil {
			res = append(res, leaves[i].Right)
		}
	}
	return res
}

func getChildrenMaxNum(leaves []*TreeNode) int {
	var max int
	for i := range leaves {
		if i == 0 {
			max = leaves[i].Val
		}
		if leaves[i].Val > max {
			max = leaves[i].Val
		}
	}
	return max
}