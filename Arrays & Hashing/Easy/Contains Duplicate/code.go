package main

// Space Complexity := O(n)
// Time Complexity := O(n)
// hash to get unique values in array, to check for duplicates easily
func containsDuplicate(nums []int) bool {
	//Hashmap makes space complexity O(n), n => the lentgh of nums
	duplicateMap := make(map[int]int)
	//O(n)
	for _, num := range nums {
		//Increasing map value with key from num
		duplicateMap[num]++
		//Check if map has value more than 1, if true nums contains duplicate
		if duplicateMap[num] > 1 {
			return true
		}
	}
	return false
}
