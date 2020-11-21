package infixtopostfix

import "strings"

func top(stack []rune) rune {
	return stack[len(stack)-1]
}

func swap(stack []rune, char rune) []rune {
	stack[len(stack)-1] = char
	return stack
}

func pop(stack []rune) []rune {
	return stack[:len(stack)-1]
}

func parenthesisPop(stack []rune, postfix *strings.Builder) []rune {
	for top(stack) != '(' {
		postfix.WriteRune(top(stack))
		stack = pop(stack)
	}
	stack = pop(stack)
	return stack
}

func popuntilempty(stack []rune, postfix *strings.Builder) []rune {
	for len(stack) > 0 {
		postfix.WriteRune(top(stack))
		stack = pop(stack)
	}
	return stack
}
