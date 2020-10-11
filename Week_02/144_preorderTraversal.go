package Week_02

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return nil
	}
	ret = append(ret, root.Val)
	ret = append(ret, preorderTraversal(root.Left)...)
	ret = append(ret, preorderTraversal(root.Right)...)
	return ret
}

/*
1.递归
1.1 每一层先遍历根节点，然后分别遍历左子树、右子树
1.2 返回条件：空
*/
