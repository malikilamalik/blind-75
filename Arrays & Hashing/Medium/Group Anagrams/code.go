package main

// // Space Complexity := O(26)
// Time Complexity := O(n + 26)
// count each char in str1, decrement for str2;
func groupAnagrams(strs []string) [][]string {
	valid := make(map[[26]int][]string)
	for _, str := range strs {
		key := [26]int{}
		for i := 0; i < len(str); i++ {
			key[str[i]-'a']++
		}
		valid[key] = append(valid[key], str)
	}
	var res [][]string
	for _, v := range valid {
		res = append(res, v)
	}
	return res
}
