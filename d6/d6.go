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

	scanner.Scan()
	firstLine := scanner.Text()
	firstLineNumbers := strings.Split(firstLine, " ")
	columns := [][]string{}
	for _, v := range firstLineNumbers {
		columns = append(columns, []string{v})
	}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		for i, v := range line {
			columns[i] = append(columns[i], v)
		}
	}

	for _, v := range columns {
		helpers.LogLine(v)
		acc, _ := strconv.Atoi(v[0])
		operator := v[len(v)-1]
		for i := 1; i < len(v)-1; i++ {
			helpers.LogLine("acc is now", acc)
			number, _ := strconv.Atoi(v[i])
			if operator == "*" {
				acc *= number
			} else {
				acc += number
			}
		}
		r1 += acc
		helpers.LogLine("acc was", acc)
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

type group struct {
	operator string
	numbers  []int
}

func getNumbers(lines []string, index int) string {
	res := ""
	for _, line := range lines {
		res += string(line[index])
	}
	return res
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
	helpers.LogLine(scanner)
	lines := []string{}
	for scanner.Scan() {
		helpers.LogLine("Added line")
		lines = append(lines, scanner.Text())
	}

	groups := []group{group{}}
	operatorLine := lines[len(lines)-1]
	numberLines := lines[:len(lines)-1]
	helpers.LogLine("numberlines", numberLines)
	for i := 0; i < len(lines[0]); i++ {
		numbers := getNumbers(numberLines, i)
		if strings.TrimSpace(numbers) == "" {
			groups = append(groups, group{})
			continue
		}

		group := groups[len(groups)-1]
		if string(operatorLine[i]) != " " {
			group.operator = string(operatorLine[i])
			group.numbers = []int{}
		}

		number, _ := strconv.Atoi(strings.TrimSpace(numbers))
		group.numbers = append(group.numbers, number)

		groups[len(groups)-1] = group

		helpers.LogLine(group)
		helpers.LogLine(groups)
	}

	for _, group := range groups {
		acc := group.numbers[0]
		for _, n := range group.numbers[1:] {
			if group.operator == "*" {
				acc *= n
			} else {
				acc += n
			}
		}
		r2 += acc
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
		{FileName: "sample1_simplified.txt", ExpectedResult: 4277556},
	}

	errors := helpers.RunTests(tests, Solve, true)

	if errors != 0 {
		log.Fatal("Had errors, not running main input")
		return
	}

	fmt.Printf("P1: %d\n", Solve("input_simplified.txt", true))
}

func P2() {
	tests := []helpers.Testcase{
		{FileName: "sample1.txt", ExpectedResult: 3263827},
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
