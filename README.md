# big-number-sqrt

Utility to calculate the root of any degree from any big number.

For example, you can get `x^(1/m)` for any `x` and any `m` from big numbers

```
package main

import (
	sqrt "github.com/zZoMROT/big-number-sqrt"
	"log"
	"math/big"
)

func main() {

	result, steps := sqrt.Sqrt(big.NewInt(100000000), big.NewInt(16))
	log.Printf("Result: %v, Steps: %d", result, steps)

}
```
