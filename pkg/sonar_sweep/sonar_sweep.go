package main

import (
	"fmt"
	"strconv"

	get_lines "github.com/vogt4nick/advent-of-code-2021/pkg/utils"
)

const path = "pkg/sonar_sweep/data/sonar-data.txt"

func readSonarData(path string) []int {
	lines := get_lines.GetLines(path)
	data := make([]int, len(lines))

	for i, line := range get_lines.GetLines(path) {
		datum, _ := strconv.Atoi(line)
		data[i] = datum
	}

	return data
}

func getDiffs(data []int) []int {
	diffs := make([]int, len(data)-1)
	for i, x2 := range data {
		if i == 0 {
			continue
		}
		x1 := data[i-1]
		diffs = append(diffs, x2-x1)
	}
	return diffs
}

func countIncreasesAndDecreases(diffs []int) (uint, uint) {
	var count_increases uint
	var count_decreases uint
	for _, d := range diffs {
		if d > 0 {
			count_increases++
		} else if d < 0 {
			count_decreases++
		}
	}
	return count_increases, count_decreases
}

func main() {
	data := readSonarData(path)
	diffs := getDiffs(data)
	count_increases, count_decreases := countIncreasesAndDecreases(diffs)
	fmt.Println("count_increases:", count_increases)
	fmt.Println("count_decreases:", count_decreases)
}
