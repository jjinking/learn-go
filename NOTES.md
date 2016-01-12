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

#### Flowcontrol [link](https://tour.golang.org/flowcontrol/1)

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

#### More types: structs, slices, and maps

- Pointers but no pointer arithmetic

- `struct` is a collection of fields, declared with `type`, kind of like a **class**

``` go
package main

import "fmt"

type Point2D struct {
	X int
	Y int
	Z int
}

func main() {
	p := Point2D{3, 4, 5}
	q := &p
	r := p
	p.X = 100
	q.Y = 200
	r.Z = 300
	fmt.Println(p)
	fmt.Println(r)
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

### Methods and Interfaces

- No classes in Go

- Implement methods on types using special *receiver* argument
  
  - Only allowed for types that are defined in same package as the method
  
  - Pointer receivers (receivers passed in with `*`) can modify value of original receiver (object), wherease value receivers are just copies of original
  
  - Methods automatically convert receivers if the method signature requires pointer receiver, but if using just simple functions that require pointer param, then user must pass in the argument with `&`

- Interfaces - set of methods decoupled from the implementation packages

	- Declare a variable of the interface type, and later assign a value to it where the value is an instance of a type that implements the methods described in the interface

``` go
package main

import (
	"fmt"
	"math"
)

type Point2D struct {
	X, Y float64
}

func (p Point2D) Dist(q Point2D) float64 {
	return math.Sqrt(math.Pow(p.X - q.X, 2) + math.Pow(p.Y - q.Y, 2))
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Interface
type Abser interface {
	Abs() float64
}

func main() {
	p1 := Point2D{0, 100}
	p2 := Point2D{100, 100}
	fmt.Println(p1.Dist(p2))
	
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	
	var a Abser
	a = f  // Allowed since MyFloat implements Abser
}
```

- Exercise: Stringers

``` go
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip *IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, &a)
	}
}
```

- Exercise: Errors

``` go
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {return x, ErrNegativeSqrt(x)}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
```

- Exercise: Readers

``` go
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int, err error) {
	for i := 0; i < 8; i++ {
		b[i] = 'A'
	}
	return 8, nil
}


func main() {
	reader.Validate(MyReader{})
}
```

- Exercise: rot13Reader

``` go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	for i := 0; i < n; i++ {
		if ('a' <= b[i] && b[i] <= 'm') || ('A' <= b[i] && b[i] <= 'M') {
			b[i] = b[i] + 13
		} else if ('n' <= b[i] && b[i] <= 'z') || ('N' <= b[i] && b[i] <= 'Z') {
			b[i] = b[i] - 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
```

- Exercise: HTTP Handlers

``` go
package main

import (
	"log"
	"net/http"
	"fmt"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (h String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, string(h))
}

func (h *Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, h)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
```

- Exercise: Images

``` go
package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (im *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (im *Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 128, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(&m)
}
```
