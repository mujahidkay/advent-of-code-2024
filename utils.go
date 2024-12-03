package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filename string) ([]int, []int) {
	var path = "inputs/" + filename

	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left []int
	var right []int
	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Fields(line)
		
		a, b := tokens[0], tokens[1]

		a_int, _ := strconv.Atoi(a)
		b_int, _ := strconv.Atoi(b)
		left = append(left, a_int)
		right = append(right, b_int)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %v\n", err)
	}

	return left, right
}
