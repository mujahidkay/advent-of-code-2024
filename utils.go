package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func multiply(groups []string) int{

	pattern := `\d+`
	result := 0
	re := regexp.MustCompile(pattern)
	shouldMult := true
	var first, second int
	for _, group := range groups {
		if group == "do()" {
			shouldMult = true
		} else if group == "don't()" {
			shouldMult = false
		} else {
			nums := re.FindAllString(group, -1)
			first, _ = strconv.Atoi(nums[0])
			second, _ = strconv.Atoi(nums[1])
			if shouldMult {
				result += (first * second)
			}
		}
	}

	return result
}

func ReadCorruptedData(filename string) int {
	file := readFile(filename)

	defer file.Close()

	pattern := `(?m)do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(file)
	var matches []string
	for scanner.Scan() {
		line := scanner.Text()

		matches = append(matches, re.FindAllString(line, -1)...)
	}

	return multiply(matches)

}