/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (45.64%)
 * Total Accepted:    338.6K
 * Total Submissions: 741.1K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 * 
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
 * 
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	ans := &ListNode{}
	pre := &ListNode{}
    ans.Next = head
	pre = ans
	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := a.Next
		pre.Next,b.Next,a.Next  = b,a,b.Next
		pre = a
	}
	return ans.Next
}
