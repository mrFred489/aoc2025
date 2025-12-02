package main

import (
	"aoc2025/helpers"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(fileName string, first bool) int {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)
	r1 := 0
	r2 := 0
	ranges := strings.Split(input, ",")
	helpers.LogLine(input)
	helpers.LogLine(ranges)
	helpers.LogLine(r1)
	helpers.LogLine(r2)
	for _, v := range ranges {
		values := strings.Split(v, "-")
		start, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal("Failed to parse start value", err, values)
			return 0
		}
		end, err := strconv.Atoi(strings.TrimSpace(values[1]))
		if err != nil {
			log.Fatal("Failed to parse end value", err, values)
			return 0
		}

		for i := start; i <= end; i++ {
			numberAsString := strconv.Itoa(i)
			stringLength := len(numberAsString)

			if stringLength%2 == 0 {
				// p1
				part1 := numberAsString[0 : stringLength/2]
				part2 := numberAsString[stringLength/2:]
				if part1 == part2 {
					r1 += i
				}
			}

			testing := false

			if numberAsString == "824824824" {
				helpers.LogLine("Evaluating 824824824")
				testing = true
			}

			// p2
			for j := 0; j < stringLength/2; j++ {
				if testing {
					helpers.LogLine(stringLength % (j + 1))
				}
				if stringLength%(j+1) != 0 {
					continue
				}

				occurrences := strings.Count(numberAsString, numberAsString[0:j+1])
				if testing {
					helpers.LogLine("testing r2 number", i, "\n\toccurrences", occurrences, "\n\tstringlength", stringLength, "\n\texpected occurrences", stringLength/(j+1), "\n\trecurring number", numberAsString[0:j+1])
				}
				if occurrences == stringLength/(j+1) {
					helpers.LogLine("found r2 number", i, "\n\toccurrences", occurrences, "\n\tstringlength", stringLength, "\n\texpected occurrences", stringLength/(j+1), "\n\trecurring number", numberAsString[0:j+1])
					r2 += i
					break
				}
			}
		}
	}

	if first {
		return r1
	} else {
		return r2
	}
}

func P1() {
	tests := []helpers.Testcase{
		{FileName: "sample1.txt", ExpectedResult: 1227775554},
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
		{FileName: "sample1.txt", ExpectedResult: 4174379265},
	}

	errors := helpers.RunTests(tests, Solve, false)

	if errors != 0 {
		log.Fatal("Had errors, not running main input")
		return
	}

	fmt.Printf("P2: %d\n", Solve("input.txt", false))
}

func main() {
	// P1()
	P2()
}
