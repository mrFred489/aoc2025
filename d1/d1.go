package main

import (
	"aoc2025/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Solve(fileName string, first bool) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	dial := 50
	password := 0
	password2 := 0
	for scanner.Scan() {
		text := scanner.Text()
		helpers.LogLine(text)
		direction := string(text[0])
		amount, err := strconv.Atoi(string(text[1:]))

		if err != nil {
			fmt.Printf("failed to parse amount %v, err: %v\n", string(text[1:]), err)
			continue
		}

		if amount > 99 {
			// fmt.Println(amount)
			// fmt.Println(amount / 100)
			password2 += amount / 100
			helpers.LogLine(password2)
		}

		amount = amount % 100

		if direction == "R" {
			dial += amount
			if dial >= 100 {
				password2 += 1
				helpers.LogLine(password2)
			}
			dial = dial % 100
		} else {
			if dial-amount < 0 {
				if dial != 0 {
					password2 += 1
					helpers.LogLine(password2)
				}
				dial = 100 + (dial - amount)
			} else {
				dial -= amount
				if dial == 0 {
					password2 += 1
					helpers.LogLine(password2)
				}
			}
		}
		if dial == 0 {
			password += 1
		}
	}

	helpers.LogLine(password)
	helpers.LogLine(password2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if first {
		return password
	} else {
		return password2
	}
}

func P1() {
	tests := []helpers.Testcase{
		{FileName: "sample1.txt", ExpectedResult: 3},
		{FileName: "sample2.txt", ExpectedResult: 0},
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
		{FileName: "sample2.txt", ExpectedResult: 11},
	}

	errors := helpers.RunTests(tests, Solve, false)

	if errors != 0 {
		log.Fatal("Had errors, not running main input")
		return
	}

	fmt.Printf("P2: %d\n", Solve("input.txt", false))
}

func main() {
	//P1()
	P2()
}
