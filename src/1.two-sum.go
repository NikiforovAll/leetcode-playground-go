/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

package twoSum

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[target-v] = i
	}

	for i, v := range nums {
		if k, contains := m[v]; contains && i != k {
			return []int{i, k}
		}
	}
	return []int{0, 0}
}

// @lc code=end
