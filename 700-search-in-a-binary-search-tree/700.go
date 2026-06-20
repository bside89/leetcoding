// Link: https://leetcode.com/problems/search-in-a-binary-search-tree/
package searchinabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Complexity: O(log n)
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	} else if val > root.Val {
		return searchBST(root.Right, val)
	} else {
		return root
	}
}
