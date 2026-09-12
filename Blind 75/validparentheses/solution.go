package main

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			// If the current char is not matching with the top element of stack
			if c == ')' && top != '(' {
				return false
			}

			if c == ']' && top != '[' {
				return false
			}

			if c == '}' && top != '{' {
				return false
			}

			// Remove the top element of the stack
			stack = stack[0 : len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
