package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Read first line of text containing a single integer from standard input
func readFirstIntLine() (int, *bufio.Reader) {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	x, _ := strconv.Atoi(strings.Trim(line, "\n"))
	return x, reader
}

// Read next line of text containing a single integer
func readNextIntLine(reader *bufio.Reader) int {
	line, _ := reader.ReadString('\n')
	x, _ := strconv.Atoi(strings.Trim(line, "\n"))
	return x
}

// Read first line of text and return as string
func readNextStrLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.Trim(line, "\n")
}

// Read next line of text containing numbers and return list of integers
func readNextIntsLine(reader *bufio.Reader) (rowInt []int) {
	line, _ := reader.ReadString('\n')
	return splitInt(strings.Trim(line, "\n"))
}

// Split a string containing space-separated list of integers
func splitInt(line string) (rowInt []int) {
	row := strings.Split(line, " ")
	rowInt = make([]int, len(row))
	for i, v := range row {
		x, _ := strconv.Atoi(v)
		rowInt[i] = x
	}
	return
}

// Minimum of a slice of ints
func minInt(xs []int) (minVal int, minIdx int) {
	minVal = xs[minIdx]
	for i := 1; i < len(xs); i++ {
		if xs[i] < minVal {
			minVal = xs[i]
			minIdx = i
		}
	}
	return
}

func arraysum() {
	reader := bufio.NewReader(os.Stdin)

	// Read in first line, which is useless
	reader.ReadString('\n')

	// Read in numbers
	line, _ := reader.ReadString('\n')

	// Trim line
	line = strings.Trim(line, "\n")
	sum := 0
	for _, v := range strings.Split(line, " ") {
		x, _ := strconv.Atoi(v)
		sum += x
	}
	fmt.Println(sum)
}

func diagdiff() {
	// Read in first line
	n, reader := readFirstIntLine()

	var d1Sum, d2Sum int
	for i := 0; i < n; i++ {
		row := readNextIntsLine(reader)
		d1Sum += row[i]
		d2Sum += row[n-i-1]
	}
	var diff int
	if d1Sum > d2Sum {
		diff = d1Sum - d2Sum
	} else {
		diff = d2Sum - d1Sum
	}
	fmt.Println(diff)
}

// https://www.hackerrank.com/challenges/plus-minus
func plusminus() {
	reader := bufio.NewReader(os.Stdin)
	// Read in first line
	line, _ := reader.ReadString('\n')
	n, _ := strconv.ParseFloat(strings.Trim(line, "\n"), 64)

	line, _ = reader.ReadString('\n')
	var nPos, nNeg, nZeros float64
	for _, v := range splitInt(strings.Trim(line, "\n")) {
		if v > 0 {
			nPos++
		} else if v < 0 {
			nNeg++
		} else {
			nZeros++
		}
	}
	fmt.Println(nPos / n)
	fmt.Println(nNeg / n)
	fmt.Println(nZeros / n)
}

// https://www.hackerrank.com/challenges/staircase
func staircase() {
	n, _ := readFirstIntLine()
	for i := 0; i < n; i++ {
		fmt.Printf("%s%s\n", strings.Repeat(" ", n-i-1), strings.Repeat("#", i+1))
	}
}

// https://www.hackerrank.com/challenges/time-conversion
func timeconv() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.Trim(line, "\n")
	arr := strings.Split(line, ":")
	hour := arr[0]
	ampm := arr[2][2:]
	arr[2] = arr[2][:2]
	if hour == "12" && ampm == "AM" {
		arr[0] = "00"
	} else if hour != "12" && ampm == "PM" {
		t, _ := strconv.Atoi(hour)
		arr[0] = strconv.Itoa(t + 12)
	}
	fmt.Println(strings.Join(arr, ":"))
}

// https://www.hackerrank.com/challenges/angry-professor
func angryprof() {
	t, reader := readFirstIntLine()

	for i := 0; i < t; i++ {
		// Read in N and K
		line, _ := reader.ReadString('\n')
		row := strings.Split(strings.Trim(line, "\n"), " ")
		strconv.Atoi(row[0])
		k, _ := strconv.Atoi(row[1])

		// Read in the students
		line, _ = reader.ReadString('\n')
		var nArrived int
		for _, s := range splitInt(strings.Trim(line, "\n")) {
			if s <= 0 {
				nArrived++
			}
		}

		if nArrived >= k {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}

// https://www.hackerrank.com/challenges/sherlock-and-the-beast
func sherlockbeast() {
	t, reader := readFirstIntLine()

	for i := 0; i < t; i++ {
		line, _ := reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.Trim(line, "\n"))

		var quot int
		for n%3 != 0 {
			n -= 5
			quot += 5
		}
		if n < 0 {
			fmt.Println(-1)
		} else {
			fmt.Printf(
				"%s%s\n",
				strings.Repeat("5", n),
				strings.Repeat("3", quot))
		}
	}
}

// https://www.hackerrank.com/challenges/utopian-tree
func utopiantree() {
	cycle := func(x int, spring bool) int {
		if spring {
			return 2 * x
		}
		return x + 1
	}

	t, reader := readFirstIntLine()

	for i := 0; i < t; i++ {
		h := 1
		ncycles := readNextIntLine(reader)
		spring := true
		for j := 0; j < ncycles; j++ {
			h = cycle(h, spring)
			spring = !spring
		}
		fmt.Println(h)
	}
}

// https://www.hackerrank.com/challenges/find-digits
func finddigits() {
	t, reader := readFirstIntLine()

	for i := 0; i < t; i++ {
		nStr := readNextStrLine(reader)
		n, _ := strconv.Atoi(nStr)
		var count int
		for _, c := range nStr {
			d := int(c - '0')
			if d != 0 && n%d == 0 {
				count++
			}
		}
		fmt.Println(count)
	}
}

// https://www.hackerrank.com/challenges/sherlock-and-squares
func sherlocksquares() {
	t, reader := readFirstIntLine()

	for i := 0; i < t; i++ {
		row := readNextIntsLine(reader)
		a, b := row[0], row[1]
		start := int(math.Sqrt(float64(a)))
		end := int(math.Sqrt(float64(b)))
		var counter int
		for x := start; x <= end; x++ {
			xSq := x * x
			if a <= xSq && xSq <= b {
				counter++
			}
		}
		fmt.Println(counter)
	}
}

func servicelane() {
	reader := bufio.NewReader(os.Stdin)
	row := readNextIntsLine(reader)
	_, t := row[0], row[1]

	widths := readNextIntsLine(reader)

	var i, j int
	for k := 0; k < t; k++ {
		row = readNextIntsLine(reader)
		i, j = row[0], row[1]
		minVal, _ := minInt(widths[i : j+1])
		fmt.Println(minVal)
	}
}

func main() {
	servicelane()
}
