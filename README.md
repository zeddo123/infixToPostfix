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
