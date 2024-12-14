package main

/*
The integer cube root of a positive number  𝑛
  is the smallest number  𝑖
  such that  𝑖3≤𝑛
  but  (𝑖+1)3>𝑛
 .
*/
func CubeRoot(n int) int {
	if n < 0 {
		panic("error")
	}
	low, high := 0, n
	var ans int

	for low <= high {
		mid := low + (high-low)/2
		cube := mid * mid * mid

		if cube <= n {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}
