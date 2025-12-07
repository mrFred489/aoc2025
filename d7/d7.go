package main

import (
	"aoc2025/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getKeysFromMap(myMap map[int]bool) []int {
	keys := make([]int, len(myMap))

	i := 0
	for k := range myMap {
		keys[i] = k
		i++
	}

	return keys
}

func addToSet(set map[int]bool, value int) map[int]bool {
	_, ok := set[value]
	if !ok {
		set[value] = true
	}

	return set
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
	scanner.Scan()
	firstLine := scanner.Text()
	indexesCurrentLine := []int{strings.Index(firstLine, "S")}

	for scanner.Scan() {
		indexesNextLine := map[int]bool{}
		line := scanner.Text()
		helpers.LogLine(line)
		helpers.LogLine(indexesCurrentLine)
		for _, index := range indexesCurrentLine {
			if string(line[index]) == "." {
				indexesNextLine = addToSet(indexesNextLine, index)
			} else {
				helpers.LogLine("Found one", index, string(line[index]))
				r1 += 1
				left := index - 1
				right := index + 1
				if left >= 0 {
					indexesNextLine = addToSet(indexesNextLine, left)
				}
				if right < len(line) {
					indexesNextLine = addToSet(indexesNextLine, right)
				}
			}
		}

		indexesCurrentLine = getKeysFromMap(indexesNextLine)
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
	scanner.Scan()
	firstLine := scanner.Text()
	initialIndex := strings.Index(firstLine, "S")
	indexesCurrentLine := []int{initialIndex}
	timelines := map[int]int{}
	timelines[initialIndex] = 1

	for scanner.Scan() {
		indexesNextLine := map[int]bool{}
		timelinesNextLine := map[int]int{}
		line := scanner.Text()
		helpers.LogLine(line)
		helpers.LogLine(indexesCurrentLine)
		helpers.LogLine(timelines)
		for _, index := range indexesCurrentLine {
			timelinesForIndex := timelines[index]
			if string(line[index]) == "." {
				indexesNextLine = addToSet(indexesNextLine, index)
				timelinesNextLine[index] = timelinesNextLine[index] + timelinesForIndex
				helpers.LogLine("Merging straight", index, timelinesNextLine[index], timelinesForIndex)
			} else {
				left := index - 1
				right := index + 1
				if left >= 0 {
					indexesNextLine = addToSet(indexesNextLine, left)
					timelinesNextLine[left] = timelinesNextLine[left] + timelinesForIndex
				}
				if right < len(line) {
					indexesNextLine = addToSet(indexesNextLine, right)
					timelinesNextLine[right] = timelinesNextLine[right] + timelinesForIndex
				}
			}
		}

		indexesCurrentLine = getKeysFromMap(indexesNextLine)
		timelines = timelinesNextLine
		helpers.LogLine("timelines", timelines)
	}

	for _, v := range timelines {
		helpers.LogLine("summing", v, timelines[v])
		r2 += v
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
		{FileName: "sample1.txt", ExpectedResult: 21},
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
		{FileName: "sample1.txt", ExpectedResult: 40},
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
