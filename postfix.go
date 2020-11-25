package intopostfix

import (
	"strings"
	"unicode"
)

var operators = map[rune]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
	'^': 3,
}

// Convert : converts a infix expression to postfix
func Convert(infix string) (postfix strings.Builder) {
	var stack []rune
	for _, i := range infix {
		if unicode.IsLetter(i) {
			postfix.WriteRune(i)
		} else {
			stack = process(i, stack, &postfix)
		}
	}
	stack = popuntilempty(stack, &postfix)
	return
}

func process(char rune, stack []rune, postfix *strings.Builder) []rune {
	switch char {
	case '(':
		stack = append(stack, char)
	case ')':
		stack = parenthesisPop(stack, postfix)
	default:
		value := operators[char]
		if len(stack) == 0 || value > operators[top(stack)] {
			// if the stack is empty or the new operator has more precedent, push it
			stack = append(stack, char)
		} else if value < operators[top(stack)] {
			stack = pop(stack)
			stack = process(char, stack, postfix)
		} else {
			postfix.WriteRune(top(stack))
			stack = swap(stack, char)
		}
	}
	return stack
}
