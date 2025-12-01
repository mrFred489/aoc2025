package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
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
		fmt.Println(text)
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
			fmt.Println(password2)
		}

		amount = amount % 100

		if direction == "R" {
			dial += amount
			if dial >= 100 {
				password2 += 1
				fmt.Println(password2)
			}
			dial = dial % 100
		} else {
			if dial-amount < 0 {
				if dial != 0 {
					password2 += 1
					fmt.Println(password2)
				}
				dial = 100 + (dial - amount)
			} else {
				dial -= amount
				if dial == 0 {
					password2 += 1
					fmt.Println(password2)
				}
			}
		}
		// 	fmt.Println(direction)
		// 	fmt.Println(amount)
		// 	fmt.Println(dial)
		if dial == 0 {
			password += 1
		}
	}

	fmt.Println(password)
	fmt.Println(password2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
