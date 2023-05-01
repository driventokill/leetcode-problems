package longestvalidparentheses

// 1. If we see `(` we push index on stack
// 2. If we see `)` we pop item from stack
// 3. If stack is empty
// 	 a. Empty stack means the end of valid subsequence
// 	 b. Not empty stack means subsequence is valid. We need find maximum of two numbers
func longestValidParentheses(s string) int {
	var stack []int
	stack = append(stack, -1)

	max := 0
	for i,e:=range s{
		if e == '(' {
			stack = append(stack, i)
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			if l := i - stack[len(stack)-1]; l > max {
				max = l
			}
		}
	}

	return max
}
