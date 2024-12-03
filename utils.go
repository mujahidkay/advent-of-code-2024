package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) *os.File {
	var path = "inputs/" + filename

	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
	}

	return file
}
func ReadInput(filename string) ([]int, []int) {
	file := readFile(filename)
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

func ReadLevels(filename string) [][]int {
	file := readFile(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var readings []int
	var multiLevelReadings [][]int
	for scanner.Scan() {
		line := scanner.Text()

		levels := strings.Fields(line)

		for _, num := range levels {
			num_int, _ := strconv.Atoi(num)
			readings = append(readings, num_int)
		}
		multiLevelReadings = append(multiLevelReadings, readings)
		readings = []int{}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %v\n", err)
	}

	return multiLevelReadings
}