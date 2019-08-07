/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (44.36%)
 * Total Accepted:    2M
 * Total Submissions: 4.5M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 *
 */

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for k, v := range nums {

		if _,ok := m[v];ok{
			if target == 2*v {
				return []int{m[target-v],k}
			}
			return []int{}
		}
		m[v] = k		
		if _, ok := m[target-v]; ok && m[target-v] != k {
			return []int{m[target-v], k}
		}
	}
	return []int{}
}

