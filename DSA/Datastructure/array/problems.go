package main

import "sort"

// insert element in a sorted array
func InsertElement(data []int, val int) []int {
	pos := FindInsertPosition(data, val)

	data = append(data, 0)
	// shift elements to the right
	for i := len(data) - 1; i > pos; i-- {
		data[i] = data[i-1]
	}
	data[pos] = val
	return data
}

func FindInsertPosition(data []int, val int) int {
	left := 0
	right := len(data) - 1
	mid := left + (right-left)/2
	for left < right {
		if data[mid] < val {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// GroupAnagrams

// Input: strs = ["act","pots","tops","cat","stop","hat"]

// Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

//o(n*klogk) - klogk - sorting ,
func SortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func GroupAnagram(words []string) [][]string {

	anagramMap := make(map[string][]string)
	for _, word := range words {
		sortedWord := SortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	result := make([][]string, 0, len(anagramMap))
	for _, v := range anagramMap {

		result = append(result, v)
	}
	return result
}


// O(nlogn)
func TopKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int, 0)

	for _, v := range nums {
		hashMap[v]++
	}

	type pair struct {
		num   int
		count int
	}
	pairs := make([]pair, 0, len(hashMap))
	for k, v := range hashMap {
		pairs = append(pairs, pair{num: k, count: v})
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].count > pairs[j].count })
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = pairs[i].num
	}
	return result
}
