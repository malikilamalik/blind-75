package main

// Space Complexity := O(n + m)
// Time Complexity := O(n + m)
// hashmap to count each char in str1, decrement for str2;
func isAnagram(s string, t string) bool {
	//O(n)
	sMap := make(map[byte]int)
	//O(m)
	tMap := make(map[byte]int)
	if len(s) != len(t) {
		return false
	}
	//O(n)
	for i, _ := range s {
		sMap[s[i]]++
		tMap[t[i]]++
	}
	//O(m)
	for key, element := range sMap {
		if tMap[key] != element {
			return false
		}
	}
	return true
}
