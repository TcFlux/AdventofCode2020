package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	seats := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seats = append(seats, scanner.Text())
	}

	part1(seats)
	part2(seats)
}

func part1(seats []string) {
	var maxID int

	for _, seatText := range seats {
		_, _, id := getSeatDetails(seatText)
		if id > maxID {
			maxID = id
		}
	}

	fmt.Println("The greatest seat ID for part 1 is:", maxID)
}

func part2(seats []string) {
	rows := make([][]int, 127)
	fullRowfound := false

	for _, seatText := range seats {
		row, seat, _ := getSeatDetails(seatText)
		rows[row] = append(rows[row], seat)
	}

	for row, seatsList := range rows {
		if fullRowfound {
			if len(seatsList) < 8 {
				sort.Ints(seatsList)
				for i, seat := range seatsList {
					if i != seat {
						fmt.Println("The missing seat ID for part 2 is:", row * 8 + i)
						break
					}
				}
			}
		} else {
			if len(seatsList) == 8 {
				fullRowfound = true
			} else {
				continue;
			}
		}
	}
}

func getSeatDetails(seatText string) (int, int, int) {
	rowData, seatData := seatText[:7], seatText[7:]
	row, seat := getRow(rowData), getSeat(seatData)

	return row, seat, (row * 8) + seat
}

func getRow(rowData string) int {
	min := 0.0
	max := 127.0

	for i := 0; i < len(rowData); i++ {
		half := rowData[i:i+1]
		if half == "F" {
			max = max - math.Ceil((max - min) / 2)
		} else {
			min = min + math.Ceil((max - min) / 2)
		}
	}

	if min != max {
		log.Fatal("you done goofed the row: min = %d, max = %d", min, max)
	}

	return int(min)
}

func getSeat(seatData string) int {
	min := 0.0
	max := 7.0
	for i := 0; i < len(seatData); i++ {
		half := seatData[i:i+1]
		if half == "L" {
			max = max - math.Ceil((max - min) / 2)
		} else {
			min = min + math.Ceil((max - min) / 2)
		}
	}

	if min != max {
		log.Fatal("you done goofed the seat: min = %d, max = %d", min, max)
	}

	return int(min)
}