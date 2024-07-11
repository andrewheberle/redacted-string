# redacted-string

This is a very simple module that can be used to print a string as all '*' characters.

## Example

```go
package main

import (
    "fmt"
    
    "github.com/andrewheberle/redacted-string"
)

func main() {
	s := "this is a string"
	fmt.Println(redacted.Redacted(s))
}
```
