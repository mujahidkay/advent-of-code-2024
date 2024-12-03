package main

import (
	"fmt"
)

func generateNewData(readings []int) [][]int{

	multiLevelNewReadings := [][]int{}

	index := 0
	subReadings := []int{}
	for index < len(readings) {
		for i, num := range readings {
			if i != index {
				subReadings = append(subReadings, num)
			}
		}
		index++
		multiLevelNewReadings = append(multiLevelNewReadings, subReadings)
		subReadings = []int{}
	}
	return multiLevelNewReadings
}

func naiveLevels(data [][]int, breakOnFirst bool) (int, [][][]int) {
    
	count := 0
	additionalData := [][][]int{}
	for _, readings := range data {
		index := 0
		size := len(readings)
		all := 0
		for (index < size-1) {
			if (readings[index] < readings[index+1] && (readings[index+1] - readings[index] <= 3 && readings[index+1] - readings[index] >= 1)) {
				all++
			}
			index++
		}
		// fmt.Println("all", all, readings)
		index = 0
		all2 := 0
		for (index < size-1) {
			if (readings[index] > readings[index+1] && (readings[index] - readings[index+1] <= 3 && readings[index] - readings[index+1] >= 1)) {
				all2++
			}
			index++
		}
		// fmt.Println("all2", all2, readings)
		if (all==size-1 && !(all2==size-1)) || (all2==size-1 && !(all==size-1)) {
			count++
			if breakOnFirst {
				return count, additionalData
			}
		} else if all == size-2 || all2 == size-2 {
			additionalData = append(additionalData, generateNewData(readings))
		}
	}
	return count, additionalData

}

func main() {
	data := ReadLevels("day2.txt")

	count, additionalData := naiveLevels(data, false)
	// fmt.Println("additional data", additionalData, count)
	countAdditional := 0
	for _, additionalMultiLevelReadings := range additionalData {
		countAdditional, _ = naiveLevels(additionalMultiLevelReadings, true)
		count += countAdditional
	}
	fmt.Println(count)
}
