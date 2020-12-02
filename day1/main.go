package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	expenses := make([]int, 0)
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		expenses = append(expenses, num)
	}

	expenseMap := part1(expenses)
	part2(expenses, expenseMap)
}

func part1 (expenses []int) map[int]bool {
	m := make(map[int]bool)

	for _, num := range expenses {
		buddyNum := 2020 - num
		if m[buddyNum] {
			product := num * buddyNum
			fmt.Printf("The two expenses that add up to 2020 are %d and %d, their product is %d\n", num, buddyNum, product)
			break
		} else {
			m[num] = true
		}
	}

	return m
}

func part2 (expenses []int, m map[int]bool) {
	for i := 0; i < len(expenses); i++ {
		num1 := expenses[i]

		for j := i + 1; j < len(expenses); j++ {
			num2 := expenses[j]

			targetNum := 2020 - num1 - num2
			if _, ok := m[targetNum]; ok {
					product := num1 * num2 * targetNum
					fmt.Printf("The three expenses that add up to 2020 are %d, %d and %d, their product is %d\n", num1, num2, targetNum, product)
					break
				}
		}
	}
}
