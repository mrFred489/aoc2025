package main

import (
	"aoc2025/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkTile(matrix []string, row int, column int) int {
	if row < 0 || row >= len(matrix) || column < 0 || column >= len(matrix[0]) {
		return 0
	}

	if string(matrix[row][column]) == "@" {
		return 1
	}

	return 0
}

func checkAllPositions(matrix []string, row int, column int) int {
	surroundingPapers := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}

			surroundingPapers += checkTile(matrix, row+i, column+j)
		}
	}

	if surroundingPapers < 4 {
		//helpers.LogLine("Found match at", row, column, "surrounded by", surroundingPapers)
		return 1
	}

	return 0
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
	matrix := []string{}
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[0]); column++ {
			if string(matrix[row][column]) != "@" {
				continue
			}
			r1 += checkAllPositions(matrix, row, column)
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
	matrix := []string{}
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	toBeRemoved := [][]int{}
	running := true
	for running {
		for row := 0; row < len(matrix); row++ {
			for column := 0; column < len(matrix[0]); column++ {
				if string(matrix[row][column]) != "@" {
					continue
				}
				if checkAllPositions(matrix, row, column) == 1 {
					toBeRemoved = append(toBeRemoved, []int{row, column})
				}
			}
		}

		if len(toBeRemoved) == 0 {
			break
		}

		r2 += len(toBeRemoved)
		for _, v := range toBeRemoved {
			oldString := matrix[v[0]]
			if v[1] == 0 {
				matrix[v[0]] = "X" + oldString[v[1]+1:]
				continue
			}

			if v[1] == len(matrix[0])-1 {
				matrix[v[0]] = oldString[:v[1]] + "X"
				continue
			}

			matrix[v[0]] = oldString[:v[1]] + "X" + oldString[v[1]+1:]
		}

		//helpers.LogLine("matrix updated\n", matrix)
		toBeRemoved = [][]int{}
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
		{FileName: "sample1.txt", ExpectedResult: 13},
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
		{FileName: "sample1.txt", ExpectedResult: 43},
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
