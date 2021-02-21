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
	return ConvertExtended(infix, operators)
}

// ConvertExtended converts a infix expression to postfix, and adds the ability to give a operators map
func ConvertExtended(infix string, operators map[rune]int) (postfix strings.Builder) {
	var stack []rune
	for _, i := range infix {
		if unicode.IsLetter(i) || unicode.IsNumber(i) {
			postfix.WriteRune(i)
		} else if unicode.IsSpace(i) {
			continue
		} else {
			stack = process(i, stack, &postfix, operators)
		}
	}
	stack = popuntilempty(stack, &postfix)
	return
}

func process(char rune, stack []rune, postfix *strings.Builder, operators map[rune]int) []rune {
	switch char {
	case '(':
		stack = append(stack, char)
	case ')':
		stack = parenthesisPop(stack, postfix)
	default:
		value := operators[char]
		if len(stack) == 0 || value > operators[top(stack)] {
			stack = append(stack, char)
		} else if value < operators[top(stack)] {
			postfix.WriteRune(top(stack))
			stack = pop(stack)
			stack = process(char, stack, postfix, operators)
		} else {
			postfix.WriteRune(top(stack))
			stack = swap(stack, char)
		}
	}
	return stack
}
