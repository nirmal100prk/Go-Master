package main

import "fmt"

func main() {
	// a := []int{1, 2}
	// b := []int{}
	// func(a, b []int) {
	// 	a[0] = 3
	// }(a, b)

	// fmt.Println(a)
	////////////////////////////////

	// a := []int{1, 2}
	// b := []int{}
	// func(a, b []int) {
	// 	// b = a
	// 	for i, v := range a {
	// 		fmt.Println(v)
	// 		b[i] = v
	// 	}

	// 	fmt.Println(b)
	// }(a, b)

	// fmt.Println(b)
	///////////////////////////

	// nums1 := []int{2, 0}
	// m := 1
	// nums2 := []int{1}
	// n := 1
	// merge(nums1, m, nums2, n)
	// fmt.Println(nums1)

	maxProfit122([]int{7, 6, 5, 4, 3, 2, 1})
}

/*
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
*/
//MergeSortedArray
func merge(nums1 []int, m int, nums2 []int, n int) {

	m--
	n--

	for i := m + n + 1; i >= 0 && n >= 0; i-- {

		if m < 0 || nums1[m] < nums2[n] {
			nums1[i] = nums2[n]
			n--

		} else {
			nums1[i] = nums1[m]
			m--

		}
	}
}

// Remove element
// func removeElement(nums []int, val int) int {

// 	i := 0
// 	for _, v := range nums {
// 		if v != val {
// 			nums[i] = v
// 			i++
// 		}

//		}
//		fmt.Println(nums)
//		return counter
//	}
func removeElement(nums []int, val int) int {
	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
			fmt.Println(nums)
		}
	}
	return i
}

// 26. Remove Duplicates from Sorted Array
func removeDuplicates(nums []int) int {
	prev := nums[0]
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[l] = nums[i]
			l++
		}
		prev = nums[i]
	}
	return l
}

//169. Majority Element
func majorityElement(nums []int) int {
	m := map[int]int{}

	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

//0, 1, 2, 3, 1, 2, 3, 4, 3
func majorityElement1(nums []int) int {
	majority_element, majority_element_frequency := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if majority_element_frequency == 0 {
			majority_element, majority_element_frequency = nums[i], 1
		} else {
			if nums[i] == majority_element {
				majority_element_frequency += 1
			} else {
				majority_element_frequency -= 1
			}
		}
	}
	return majority_element
}

//189. Rotate Array
/*
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
*/
func rotate(nums []int, k int) {
	k = k % len(nums)

	if k == 0 {
		return
	}
	a := nums[len(nums)-k:]
	b := nums[:len(nums)-k]
	c := append(a, b...)
	for i, v := range c {
		nums[i] = v
	}
}

//121. Best Time to Buy and Sell Stock
func maxProfit(prices []int) int {
	maxProfit := 0
	buyableAmount := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < buyableAmount {
			buyableAmount = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-buyableAmount)
		}
	}
	return maxProfit
}

/*
122. Best Time to Buy and Sell Stock II

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.
Example 2:

Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.
Example 3:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
*/

func Accumulate(elements ...int) int {
	sum := 0
	for _, elem := range elements {
		sum += elem
	}
	return sum
}

func maxProfit122(prices []int) int {

	profit := make([]int, 0)

	for i := 0; i < len(prices)-1; i++ {

		if prices[i] < prices[i+1] {
			profit = append(profit, (prices[i+1] - prices[i]))
		}
	}

	return Accumulate(profit...)
}

/*
55. Jump Game
Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
*/

func canJump(nums []int) bool {
	return true
}
