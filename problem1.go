package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
    left, right := ReadInput("day1.txt")

    sortedLeft := make([]int, len(left))
    sortedRight := make([]int, len(right))

    copy(sortedLeft, left)
    copy(sortedRight, right)

    sort.Ints(sortedLeft)
    sort.Ints(sortedRight)

    totalDistance := 0
    similarityScore := 0

    for i, num := range sortedLeft {
        totalDistance += int(math.Abs(float64(sortedLeft[i] - sortedRight[i])))
        count := 0
        for _, num2 := range sortedRight {
            if num == num2 {
                count++
            }
        }

        similarityScore += (num*count)
    }

    fmt.Println(totalDistance, similarityScore)


}
