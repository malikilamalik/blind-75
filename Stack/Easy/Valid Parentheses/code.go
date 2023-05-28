package main

// Space Complexity := O(n)
// Time Complexity := O(n)
// push opening brace on stack, pop if matching close brace, at end if stack empty, return true;
func isValid(s string) bool {
	//Space Complexity O(n)
	stack := make([]byte, 0)
	//Create close parenthese map
	parenthesesMap := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var pair byte
	var ok bool
	//Time Complexity O(n)
	for _, char := range []byte(s) {
		//Check if not close parenthese
		if pair, ok = parenthesesMap[char]; !ok {
			//add to open parenthese to stack
			stack = append(stack, char)
			continue
		}

		//return false if stack empty
		if len(stack) == 0 {
			return false
		}

		//Check if stack is close parentheses
		if stack[len(stack)-1] != pair {
			return false
		}
		//Pop from stack
		stack = stack[:len(stack)-1]
	}
	//Return true if stack empty
	return len(stack) == 0
}
