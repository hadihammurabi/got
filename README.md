# got
Golang result building

# Examples
```go
package main

import (
	"fmt"

	"github.com/hadihammurabi/got"
)

func IsOdd(num int) got.Result {
	if num%2 == 0 {
		return got.Err(fmt.Sprintf("%d is not odd", num))
	}

	return got.Ok(true)
}

func main() {
	result := IsOdd(1)
	if result.HasErr() {
		panic(result.Err())
	}

	fmt.Println(result.Data())
}

```
