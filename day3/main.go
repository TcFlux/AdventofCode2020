package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var width int

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gridY := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gridX := strings.Split(scanner.Text(), "")
		gridY = append(gridY, gridX)
	}
	
	width = len(gridY[0])

	part1(gridY)
	// part2(gridY)
}

func calculateCollisions(grid [][]string, distX, distY int) int {
	collisions := 0
	posX := 0

	for posY, gridX := range grid {
		if posY == 0 {
			width = len(gridX)
		} else {
			posX += distX
			if posX >= width {
				posX = posX % width
			}
			if gridX[posX] == "#" {
				collisions += 1
			}
		}
		posY += distY
	}

	return collisions
}

func part1(grid [][]string) {
	collisions := calculateCollisions(grid, 3, 1)

	fmt.Printf("The sled would hit %d trees\n", collisions)
}

// func part2(grid [][]string) {
// 	collisions := 0
// 	posX := 0
// 	width := 0

// 	for posY, gridX := range grid {
// 		if posY == 0 {
// 			width = len(gridX)
// 		} else {
// 			posX += 3
// 			if posX >= width {
// 				posX = posX % width
// 			}
// 			if gridX[posX] == "#" {
// 				collisions += 1
// 			}
// 		}
// 		posY += 1
// 	}

// 	fmt.Printf("The sled would hit %d trees\n", collisions)
// }
