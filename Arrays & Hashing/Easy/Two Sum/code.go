package main

// Space Complexity := O(n)
// Time Complexity := O(n)
// use hash map to instantly check for difference value, map will add index of last occurrence of a num, don’t use same element twice;
func twoSum(nums []int, target int) []int {
	//Hashmap makes space complexity O(n), n => the lentgh of nums
	// map[listIndex]numValue
	sum := make(map[int]int)
	//O(n)
	for i, num := range nums {
		//O(1)
		//Check key index from hashmap value, using target-num value
		if key, ok := sum[target-num]; ok {
			return []int{key, i}
		} else {
			//If not exist insert into hashmap
			sum[num] = i
		}
	}
	return []int{}
}
