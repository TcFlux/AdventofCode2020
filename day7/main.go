package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		bagType, contents := strings.Split(scanner.Text(), "contain ")[0], strings.Split(scanner.Text(), "contain ")[1]
		contentsList := strings.Split(strings.Trim(contents, "."), ", ")
		rules[bagType] = contentsList
	}

	part1(rules)
	// part2(rules)
}

func part1(rules map[string][]string) {
	bagMemo := make(map[string]bool)
	count := 0
	hasGold := false;

	for bagType, _ := range rules {
		hasGold, bagMemo = containsShinyGold(bagType, rules, bagMemo)
		if hasGold {
			// fmt.Println(bagType, hasGold)
			count += 1
		}
		hasGold = false
	}

	// fmt.Printf("%+v\n", bagMemo)

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
			bagMemo[innerType] = true
			bagMemo[bagType] = true
			return true, bagMemo
		} else {
			fmt.Println(bagType)
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

func parseBag(bagText string) (string, string) {
	bagData := strings.Split(bagText, " ")
	if len(bagData) == 3 {
		return "0", ""
	} else {
		quantity := bagData[0]
		bagType := fmt.Sprintf("%s %s", bagData[1], bagData[2])
		return quantity, bagType
	}
}