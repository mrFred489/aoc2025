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

func FindLargestInString(bank string) int {
	for i := 9; i >= 0; i-- {
		index := strings.Index(bank, strconv.Itoa(i))
		if index != -1 {
			return index
		}
	}

	return -1
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
	for scanner.Scan() {
		bank := scanner.Text()
		first := FindLargestInString(bank[:len(bank)-1])
		second := FindLargestInString(bank[first+1:]) + first + 1
		helpers.LogLine(first, string(bank[first]), second, string(bank[second]))
		res, _ := strconv.Atoi(string(bank[first]) + string(bank[second]))
		r1 += res
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
	for scanner.Scan() {
		bank := scanner.Text()
		helpers.LogLine("Now evaluating", bank)
		prev := FindLargestInString(bank[:len(bank)-11])
		totalValue := string(bank[prev])
		for i := 1; i < 12; i++ {
			index := FindLargestInString(bank[prev+1:len(bank)-(11-i)]) + prev + 1
			prev = index
			totalValue += string(bank[index])
			helpers.LogLine("Found", string(bank[index]), "at index", index)
		}
		helpers.LogLine("Final number is", totalValue)
		res, _ := strconv.Atoi(totalValue)
		r2 += res
	}

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
		{FileName: "sample1.txt", ExpectedResult: 357},
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
		{FileName: "sample1.txt", ExpectedResult: 3121910778619},
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
