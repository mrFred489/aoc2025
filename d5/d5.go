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

func Solve(fileName string, first bool) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r1 := 0
	r2 := 0
	set := map[int]bool{}
	for scanner.Scan() {
		freshRange := scanner.Text()
		if freshRange == "" {
			break
		}
		startEnd := strings.Split(freshRange, "-")
		start, _ := strconv.Atoi(startEnd[0])
		end, _ := strconv.Atoi(startEnd[1])
		for i := start; i <= end; i++ {
			_, ok := set[i]
			if !ok {
				set[i] = true
			}
		}
	}

	helpers.LogLine("Checking ids, size of set is", len(set))

	for scanner.Scan() {
		id, _ := strconv.Atoi(scanner.Text())
		_, ok := set[id]
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
		{FileName: "sample1.txt", ExpectedResult: 6},
	}

	errors := helpers.RunTests(tests, Solve, false)

	if errors != 0 {
		log.Fatal("Had errors, not running main input")
		return
	}

	fmt.Printf("P2: %d\n", Solve("input.txt", false))
}

func main() {
	P1()
	// P2()
}
