package main

import (
	"aoc2025/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkInclusion(ranges [][]int, number int) bool {
	for _, v := range ranges {
		if v[0] <= number && number <= v[1] {
			return true
		}
	}
	return false
}

func checkInclusion2(ranges [][]int, number int) int {
	for i, v := range ranges {
		if v[0] <= number && number <= v[1] {
			return i
		}
	}
	return -1
}

func checkIfRangesCompletelyIncluded(ranges [][]int, start int, end int) []int {
	includedRanges := []int{}
	for i, v := range ranges {
		if start <= v[0] && v[0] <= end && start <= v[1] && v[1] <= end {
			includedRanges = append(includedRanges, i)
		}
	}
	return includedRanges
}

func lookupIfExists(ranges [][]int, index int) []int {
	if index == -1 {
		return []int{}
	}

	return ranges[index]
}

func Solve(fileName string, first bool) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r1 := 0
	r2 := 0
	ranges := [][]int{}
	for scanner.Scan() {
		freshRange := scanner.Text()
		if freshRange == "" {
			break
		}
		startEnd := strings.Split(freshRange, "-")
		start, _ := strconv.Atoi(startEnd[0])
		end, _ := strconv.Atoi(startEnd[1])
		ranges = append(ranges, []int{start, end})
	}

	helpers.LogLine("Checking ids, size of set is", len(ranges))

	for scanner.Scan() {
		id, _ := strconv.Atoi(scanner.Text())
		ok := checkInclusion(ranges, id)
		if ok {
			r1 += 1
		}
	}

	helpers.LogLine(r1)
	helpers.LogLine(r2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if first {
		return r1
	} else {
		return r2
	}
}

func Solve2(fileName string, first bool) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r1 := 0
	r2 := 0
	ranges := [][]int{}
	for scanner.Scan() {
		freshRange := scanner.Text()
		if freshRange == "" {
			break
		}
		startEnd := strings.Split(freshRange, "-")
		start, _ := strconv.Atoi(startEnd[0])
		end, _ := strconv.Atoi(startEnd[1])

		startRangeIndex := checkInclusion2(ranges, start)
		endRangeIndex := checkInclusion2(ranges, end)
		includedRanges := checkIfRangesCompletelyIncluded(ranges, start, end)

		helpers.LogLine("range", start, end)
		helpers.LogLine("\tstart included in", startRangeIndex, lookupIfExists(ranges, startRangeIndex))
		helpers.LogLine("\tend included in", endRangeIndex, lookupIfExists(ranges, endRangeIndex))
		helpers.LogLine("\tincludes ranges", includedRanges)

		if startRangeIndex != -1 {
			if endRangeIndex != -1 {
				if startRangeIndex == endRangeIndex {
					// included in same range
					helpers.LogLine("-- Included in same range")
					continue
				} else {
					// combining two ranges
					newRange := []int{ranges[startRangeIndex][0], ranges[endRangeIndex][1]}
					if startRangeIndex < endRangeIndex {
						// start range is before end range
						if startRangeIndex == endRangeIndex-1 {
							helpers.LogLine("-- combining start before end, no middle section")
							ranges = append(ranges[:startRangeIndex], ranges[endRangeIndex+1:]...)
						} else {
							helpers.LogLine("-- combining start before end, including middle section")
							ranges = append(append(ranges[:startRangeIndex], ranges[startRangeIndex+1:endRangeIndex]...), ranges[endRangeIndex+1:]...)
						}

						ranges = append(ranges, newRange)
					} else {
						// end range is before start range
						if startRangeIndex-1 == endRangeIndex {
							helpers.LogLine("-- combining end before start, no middle section")
							ranges = append(ranges[:endRangeIndex], ranges[startRangeIndex+1:]...)
						} else {
							helpers.LogLine("-- combining end before start, including middle section")
							ranges = append(append(ranges[:endRangeIndex], ranges[endRangeIndex+1:startRangeIndex]...), ranges[startRangeIndex+1:]...)
						}

						ranges = append(ranges, newRange)
					}
				}
			} else {
				// extending end of a range
				helpers.LogLine("-- Extending end of range")
				newRange := []int{ranges[startRangeIndex][0], end}
				ranges = append(append(ranges[:startRangeIndex], newRange), ranges[startRangeIndex+1:]...)
			}
		} else if endRangeIndex != -1 {
			// extending start of a range
			helpers.LogLine("-- Extending start of range")
			newRange := []int{start, ranges[endRangeIndex][0]}
			ranges = append(append(ranges[:endRangeIndex], newRange), ranges[endRangeIndex+1:]...)
		} else {
			if len(includedRanges) > 0 {
				for i := len(includedRanges) - 1; i >= 0; i-- {
					helpers.LogLine("-- Removing included range", i, includedRanges[i])
					ranges = append(ranges[:includedRanges[i]], ranges[includedRanges[i]+1:]...)
				}
			}

			ranges = append(ranges, []int{start, end})
		}

		helpers.LogLine("-- now ranges is", len(ranges), ranges[len(ranges)-1])
	}

	for _, v := range ranges {
		r2 += v[1] - v[0] + 1
	}

	helpers.LogLine(r1)
	helpers.LogLine(r2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if first {
		return r1
	} else {
		return r2
	}
}

func P1() {
	tests := []helpers.Testcase{
		{FileName: "sample1.txt", ExpectedResult: 3},
	}

	errors := helpers.RunTests(tests, Solve, true)

	if errors != 0 {
		log.Fatal("Had errors, not running main input")
		return
	}

	fmt.Printf("P1: %d\n", Solve("input.txt", true))
}

func P2() {
	tests := []helpers.Testcase{
		{FileName: "sample1.txt", ExpectedResult: 14},
	}

	errors := helpers.RunTests(tests, Solve2, false)

	if errors != 0 {
		log.Fatal("Had errors, not running main input")
		return
	}

	fmt.Printf("P2: %d\n", Solve2("input.txt", false))
}

func main() {
	// P1()
	P2()
}
