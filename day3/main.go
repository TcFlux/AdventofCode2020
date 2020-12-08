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
	part2(gridY)
}

func part1(grid [][]string) {
	fmt.Printf("The sled would hit %d trees for part 1\n", calculateCollisions(grid, 3, 1))
}

func part2(grid [][]string) {
	// index 0 = distX, index 1 = distY
	paths := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	collisionsList := make([]int, 0)

	for _, coords := range paths {
		collisions := calculateCollisions(grid, coords[0], coords[1])
		collisionsList = append(collisionsList, collisions)
		fmt.Printf("Going right %d, and down %d collides with %d trees\n", coords[0], coords[1], collisions)
	}
	fmt.Println(collisionsList)
	fmt.Printf("The product of the collisions for part 2 would be %d\n", sumArray(collisionsList))
}

func calculateCollisions(grid [][]string, distX, distY int) int {
	collisions := 0
	posX := 0

	for posY := 0; posY < len(grid); posY += distY {
		if posY != 0 {
			posX += distX
			if posX >= width {
				posX = posX % width
			}
			if grid[posY][posX] == "#" {
				collisions += 1
			}
		}
	}

	return collisions
}

func sumArray(array []int) int {
	product := 1
	for _, val := range array {
		product *= val
	}
	return product
}
