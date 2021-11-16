/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

package isPalindrome

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var ints []int

	for x > 0 {
		ints = append(ints, x%10)
		x = x / 10
	}

	i, j := 0, len(ints)-1

	for i < j {
		if ints[i] != ints[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// @lc code=end
