# intopostfix
This is a simple package to convert infix expressions to postfix.

## Example usage
```go
package main

import (
	"fmt"
	"github.com/zeddo123/intopostfix"
)

func main() {
	output := intopostfix.Convert("(a+b)*c")
	fmt.Println(output.String())
}
```
By using the `Convert` function, a map will be used to specify the precedence rules. To specify your own, use the function `ConvertExtended`.
```go
func main() {
	var operators = map[rune]int {
		'*': 1,
		'+': 2,
	}
	output := intopostfix.ConvertExtended("(a+b)*c", operators)
	fmt.Println(output.String())
}
```
