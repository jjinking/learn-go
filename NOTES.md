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

- `if` and `for` require braces, but the conditional expressions do not require parens

- Only `for` loops, and init and post statements are optional, which basically turns it into **while** loop

- Infinite loops can be done with `for {}`

- `if` can have **short** statements, like **init** statements in `for`, whose namespace i available in the `else` block

- `defer` argument evaluated immediately, but function call not executed until surrounding function returns
  - Multiple `defer`s are pushed onto a stack, and executed in LIFO order

``` go
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		if zNew := z - (z * z - x) / (2.0 * z); math.Abs(zNew - z) > .0000001 {
			z = zNew
		} else {
			return z
		}
	}
}

func main() {
	total := 0
	for i := 0; i < 100; i++ {
		total += i
	}
	fmt.Println(total)
}
```

### More types: structs, slices, and maps

- Pointers but no pointer arithmetic

- `struct` is a collection of fields, declared with `type`, kind of like a **class**

``` go
package main

import "fmt"

type Point2D struct {
	X int
	Y int
}

func main() {
	p := Point2D{3, 4}
	q := &p
	q.X = 100
	p.Y = 200
	fmt.Println(v.X)
}
```

- Arrays cannot be resized - length is part of the type

- Arrays are inconvenient to use, so mostly use `slice`

	- **Slice** objects point to an array, with length and capacity
	
	- Slice similar to array, but leave out # of elments, i.e. `letters := []string{"a", "b", "c", "d"}`

	- `make` function creates a slice object, with specified length and capacity (optional)

``` go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}
```

- Exercise: Slices

``` go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy) // One row per unit of y.
	for x := range picture {
		picture[x] = make([]uint8, dx)
		for y := 0; y < dx; y++ {
			picture[x][y] = uint8((x + y) / 2)
		}
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
```

- Maps exercise

``` go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		m[w] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
```

- Functions are first-class citizens

- Function closures - functions bound to a variable outside its block

- Exercise - Fibonacci closure

``` go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() (c int) {
		c = a + b
		a = b
		b = c
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

