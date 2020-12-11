package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	resps := make([][]string, 0)
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resps = append(resps, strings.Split(scanner.Text(), ""))
	}

	part1(resps)
	part2(resps)
}

func part1(resps [][]string) {
	answerMap := make(map[string]bool)
	counts := make([]int, 0)
	
	for _, answers := range resps {
		if len(answers) == 0 {
			counts = append(counts, len(answerMap))
			answerMap = make(map[string]bool)
		} else {
			for _, answer := range answers {
				if _, ok := answerMap[answer]; !ok {
					answerMap[answer] = true
				}
			}
		}
	}

	// input file doesn't end in a newline
	counts = append(counts, len(answerMap))

	sum := 0
	for _, count := range counts {
		sum += count
	}

	fmt.Println("The sum of distinct yes answers per group for part 1 is:", sum)
}

func part2(resps [][]string) {
	answerMap := make(map[string]bool)
	counts := make([]int, 0)
	groupFirst := true
	
	for _, answers := range resps {
		if len(answers) == 0 {
			counts = append(counts, len(answerMap))
			answerMap = make(map[string]bool)
			groupFirst = true
		} else {
			// Don't worry about deleting keys for the first set of answers per group
			if groupFirst {
				for _, answer := range answers {
					if _, ok := answerMap[answer]; !ok {
						answerMap[answer] = false
					}
				}
			} else {
				// For subsequent group members set overlapping yes responses to true
				for _, answer := range answers {
					if _, ok := answerMap[answer]; ok {
						answerMap[answer] = true
					}
				}
				// Delete not trues and set trues to false for next round of comparisons
				for k, v := range answerMap {
					if v == false {
						delete(answerMap, k)
					} else {
						answerMap[k] = false
					}
				}
			}

			if groupFirst {
				groupFirst = false
			}
		}
	}

	counts = append(counts, len(answerMap))

	sum := 0
	for _, count := range counts {
		sum += count
	}

	fmt.Println("The sum of answers that the entire group answered yes for in part 2 is:", sum)
}