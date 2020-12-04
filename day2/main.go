package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	validPwCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		vals, pw := strings.Split(txt, ":")[0], strings.Split(txt, ":")[1]
		freq, char := strings.Split(vals, " ")[0], strings.Split(vals, " ")[1]
		minStr, maxStr := strings.Split(freq, "-")[0], strings.Split(freq, "-")[1]
		min, _ := strconv.Atoi(minStr)
		max, _ := strconv.Atoi(maxStr)
		count := strings.Count(pw, char)
		if count >= min && count <= max {
			validPwCount += 1
		}
	}

	fmt.Println("The number of valid passwords for part 1 is:", validPwCount)
}

func part2() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	validPwCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		vals, pw := strings.Split(txt, ":")[0], strings.Split(txt, ":")[1]
		positions, char := strings.Split(vals, " ")[0], strings.Split(vals, " ")[1]
		pos1Str, pos2Str := strings.Split(positions, "-")[0], strings.Split(positions, "-")[1]
		pos1, _ := strconv.Atoi(pos1Str)
		pos2, _ := strconv.Atoi(pos2Str)
		if string(pw[pos1]) != string(pw[pos2]) {
			if string(pw[pos1]) == char || string(pw[pos2]) == char {
				validPwCount += 1
			}
		}
	}

	fmt.Println("The number of valid passwords for part 2 is:", validPwCount)
}