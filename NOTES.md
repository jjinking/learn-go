## Notes on "Tour of Go" [link](https://tour.golang.org/list)

### Basics [link](https://tour.golang.org/basics)

#### Packages, variables, and functions [link](https://tour.golang.org/basics/1)

- Every go program made of packages

- Names that begin with capital letters are exported

- Strongly typed - type comes after variable name

- Named return values

  - return variables defined at the top of function

  - allows "naked" returns, which are `return` statemenst w/o arguments, resulting in bad readability for long functions

- Declare a list of variables using keyword `var`

- Short variable declarations using `:=` inside functions - no `var` necessary, and data type is implicit

- Variables declared w/o initial values are defaulted to corresponding zero values for the data type

- Constants declared with `const` keyword, and cannot be declared using `:=`

``` go
package main

import (
	"fmt"
	"math"
)

const (
	Big   = 1 << 100 // High precision numeric constant
	Small = Big >> 99
)

// Variable declarations in block
var (
	YesOrNo bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

var a, b, c int
var u, v string = "foo", "bar"

func bSum(x, y int) int {
	return x + y
}

// Return multiple values
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
func fourSix() (x, y int) {
	x = 2 + 2
	y = 3 + 3
	return
}

func main() {
	fmt.Println(math.Pi)
	fmt.Println(bSum(1, 3))
	a, b := swap("foo", "bar")
	fmt.Println(a, b)
	
	// inferring type from initial values
	var i int = 1
	var x, y, z = 1, "hello", i
	fmt.Println(x, y)
	
	const format = "%T(%v)\n"
	fmt.Printf(format, YesOrNo, YesOrNo)
	fmt.Printf(format, MaxInt, MaxInt)
	fmt.Printf(format, z, z)
}
```

### Flowcontrol [link](https://tour.golang.org/flowcontrol/1)
