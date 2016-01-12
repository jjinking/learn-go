package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInt(reader *bufio.Reader) int {
	line, _ := reader.ReadString('\n')
	x, _ := strconv.Atoi(strings.Trim(line, "\n"))
	return x
}

func readStr(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.Trim(line, "\n")
}

func splitInt(line string) (rowInt []int) {
	row := strings.Split(line, " ")
	rowInt = make([]int, len(row))
	for i, v := range row {
		x, _ := strconv.Atoi(v)
		rowInt[i] = x
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
	reader := bufio.NewReader(os.Stdin)

	// Read in first line
	line, _ := reader.ReadString('\n')
	line = strings.Trim(line, "\n")
	n, _ := strconv.Atoi(line)

	var d1Sum, d2Sum int
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		row := splitInt(strings.Trim(line, "\n"))
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
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(line)
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
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.Trim(line, "\n"))

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
	reader := bufio.NewReader(os.Stdin)
	t := readInt(reader)

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

	reader := bufio.NewReader(os.Stdin)
	t := readInt(reader)

	for i := 0; i < t; i++ {
		h := 1
		ncycles := readInt(reader)
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
	reader := bufio.NewReader(os.Stdin)
	t := readInt(reader)

	for i := 0; i < t; i++ {
		nStr := readStr(reader)
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

func main() {
	finddigits()
}
