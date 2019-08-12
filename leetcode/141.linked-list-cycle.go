/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 *
 * https://leetcode.com/problems/linked-list-cycle/description/
 *
 * algorithms
 * Easy (37.60%)
 * Total Accepted:    438.1K
 * Total Submissions: 1.2M
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * Given a linked list, determine if it has a cycle in it.
 * 
 * To represent a cycle in the given linked list, we use an integer pos which
 * represents the position (0-indexed) in the linked list where tail connects
 * to. If pos is -1, then there is no cycle in the linked list.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: head = [3,2,0,-4], pos = 1
 * Output: true
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * second node.
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: head = [1,2], pos = 0
 * Output: true
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * first node.
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: head = [1], pos = -1
 * Output: false
 * Explanation: There is no cycle in the linked list.
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * Follow up:
 * 
 * Can you solve it using O(1) (i.e. constant) memory?
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func hasCycle(head *ListNode) bool {
	fast := &ListNode{}
    slow := &ListNode{}
	fast = head
	slow = head

	for fast != nil && slow != nil  && fast.Next != nil{

		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
		
	}
	return false	
}


/*
func hasCycle(head *ListNode) bool {
	ans := &ListNode{}
    cur := &ListNode{}
	ans.Next = head
	m := make(map[*ListNode]int)
	cur = ans
	for cur.Next != nil{
		if _,ok := m[cur.Next];ok{
			return true
		}
		m[cur.Next] = 1
		cur = cur.Next
	}
	return false
}
*/
