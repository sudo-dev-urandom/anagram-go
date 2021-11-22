package refactoringproblem

import "strings"

func findFirstStringInBracket(s string) string {
	// get start pointer and last pointer
	x := strings.IndexByte(s, '(')
	y := strings.IndexByte(s[x:], ')')
	// throw error
	if x < 0 || y < 0 {
		return ""
	}
	// return only the content
	return s[x+1 : x+y]
}
