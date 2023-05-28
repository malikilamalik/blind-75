package main

// // Space Complexity := O(26)
// Time Complexity := O(n + 26)
// count each char in str1, decrement for str2;
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	//O(26)
	var freq [26]int

	//O(n)
	for idx := 0; idx < len(s); idx++ {
		freq[s[idx]-'a']++
		freq[t[idx]-'a']--
	}

	//O(26)
	for idx := 0; idx < len(freq); idx++ {
		if freq[idx] != 0 {
			return false
		}
	}

	return true
}
