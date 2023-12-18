# int32linter

int32linter is a static analyzer to check usage of int32 variable.

# Install
to get int32linter execute this :

 $ go install github.com/beyzaekici/int32linter@latest


## Example

```go
package a

import (
        "fmt"
)

func foo() {
       	var x int32 = 5
	fmt.Println(x)
}
```

../main.go:11:8: int32 type is used, consider using int or int64

