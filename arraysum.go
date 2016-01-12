package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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
