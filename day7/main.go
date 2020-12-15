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
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rules := make(map[string][]string)
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bagType, contents := strings.Split(scanner.Text(), " bags contain ")[0], strings.Split(scanner.Text(), " bags contain ")[1]
		contentsList := strings.Split(strings.Trim(contents, "."), ", ")
		rules[bagType] = contentsList
	}

	part1(rules)
	part2(rules)
}

func part1(rules map[string][]string) {
	bagMemo := make(map[string]bool)
	count := 0
	hasGold := false;

	for bagType, _ := range rules {
		hasGold, bagMemo = containsShinyGold(bagType, rules, bagMemo)
		if hasGold {
			count += 1
		}
		hasGold = false
	}

	fmt.Printf("There are %d bag colors that contain atleast one shiny gold bag in part 1\n", count)
}

func containsShinyGold(bagType string, rules map[string][]string, bagMemo map[string]bool) (bool, map[string]bool) {
	if val, ok := bagMemo[bagType]; ok {
		return val, bagMemo
	}
	for _, innerBag := range rules[bagType] {
		_, innerType := parseBag(innerBag)
		if innerType == "" {
			bagMemo[innerType] = false
		} else if innerType == "shiny gold" {
			bagMemo[bagType] = true
			return true, bagMemo
		} else {
			containsGold := false
			containsGold, bagMemo = containsShinyGold(innerType, rules, bagMemo)
			bagMemo[innerType] = containsGold
			if containsGold {
				return containsGold, bagMemo
			}
		}
	}
	bagMemo[bagType] = false
	return false, bagMemo
}

func part2(rules map[string][]string) {
	count := countInnerBags("shiny gold", rules)

	fmt.Printf("A shiny gold bag contains %d individual bags in part 2\n", count)
}

func countInnerBags(bagType string, rules map[string][]string) int {
	count := 0
	for _, innerBag := range rules[bagType] {
		quantity, innerType := parseBag(innerBag)
		innerCount := countInnerBags(innerType, rules)
		if innerCount > 0 {
			count += quantity * innerCount + quantity
		} else {
			count += quantity
		}
		// fmt.Println(int(count))
	}

	return count
}

func parseBag(bagText string) (int, string) {
	bagData := strings.Split(bagText, " ")
	if len(bagData) == 3 {
		return 0, ""
	} else {
		quantity, err := strconv.Atoi(bagData[0])
		if err != nil {
			log.Fatal(err)
		}
		bagType := fmt.Sprintf("%s %s", bagData[1], bagData[2])
		return quantity, bagType
	}
}