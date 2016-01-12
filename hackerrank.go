package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {

}
