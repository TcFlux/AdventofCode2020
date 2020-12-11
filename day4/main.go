package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

var eyeColors = []string{
	"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	passports := make([]passport, 0)
	currentPassport := new(passport)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			passports = append(passports, *currentPassport)
			currentPassport = new(passport)
		} else {
			data := strings.Split(scanner.Text(), " ")
			for _, fieldDataStr := range data {
				fieldData := strings.Split(fieldDataStr, ":")
				field, val := fieldData[0], fieldData[1]
				currentPassport.setPassportField(field, val)
			}
		}
	}
	// input file doesn't end in a newline
	passports = append(passports, *currentPassport)
	
	part1(passports)
	part2(passports)
}

func part1(passports []passport) {
	validPassports := 0
	for _, passport := range passports {
		if passport.isLooselyValid() {
			validPassports += 1
		}
	}

	fmt.Printf("There are %d valid passports for part 1\n", validPassports)
}

func part2(passports []passport) {
	validPassports := 0
	for _, passport := range passports {
		if passport.isValid() {
			validPassports += 1
		}
	}

	fmt.Printf("There are %d valid passports for part 2\n", validPassports)
}


func (p *passport) setPassportField(field, value string) {
	switch field{
	case "byr":
		p.byr = value
	case "iyr":
		p.iyr = value
	case "eyr":
		p.eyr = value
	case "hgt":
		p.hgt = value
	case "hcl":
		p.hcl = value
	case "ecl":
		p.ecl = value
	case "pid":
		p.pid = value
	case "cid":
		p.cid = value
	}
}

func (p *passport) isLooselyValid() bool {
	return p.byr != "" &&
				 p.iyr != "" &&
				 p.eyr != "" &&
				 p.hgt != "" &&
				 p.hcl != "" &&
				 p.ecl != "" &&
				 p.pid != ""
}

func (p *passport) isValid() bool {
	return p.byrValid() &&
				 p.iyrValid() &&
				 p.eyrValid() &&
				 p.hgtValid() &&
				 p.hclValid() &&
				 p.eclValid() &&
				 p.pidValid()
}

func (p *passport) byrValid() bool {
	if p.byr == "" {
		return false
	}

	if len(p.byr) != 4 {
		return false
	}

	year, err := strconv.Atoi(p.byr)
	if err != nil {
		return false
	}

	if year < 1920 || year > 2002 {
		return false
	}

	return true
}

func (p *passport) iyrValid() bool {
	if p.iyr == "" {
		return false
	}

	if len(p.iyr) != 4 {
		return false
	}

	year, err := strconv.Atoi(p.iyr)
	if err != nil {
		return false
	}

	if year < 2010 || year > 2020 {
		return false
	}

	return true
}

func (p *passport) eyrValid() bool {
	if p.eyr == "" {
		return false
	}

	if len(p.eyr) != 4 {
		return false
	}

	year, err := strconv.Atoi(p.eyr)
	if err != nil {
		return false
	}

	if year < 2020 || year > 2030 {
		return false
	}

	return true
}

func (p *passport) hgtValid() bool {
	if p.hgt == "" {
		return false
	}

	unit := p.hgt[len(p.hgt)-2:]
	measurement, err := strconv.Atoi(p.hgt[:len(p.hgt)-2])
	if err != nil {

		return false
	}

	if unit == "cm" {
		if measurement < 150 || measurement > 193 {
			return false
		}
	} else if unit == "in" {
			if measurement < 59 || measurement > 76 {
				return false
			}
	} else {
		return false
	}
	return true
}

func (p *passport) hclValid() bool {
	if p.hcl == "" {
		return false
	}

	if p.hcl[:1] != "#" {
		return false
	}

	_, err := strconv.ParseUint(p.hcl[1:], 16, 64)
	if err != nil {
		return false
	}
	return true
}

func (p *passport) eclValid() bool {
	if p.ecl == "" {
		return false
	}

	for _, color := range eyeColors {
		if p.ecl == color {
			return true
		}
	}
	return false
}

func (p *passport) pidValid() bool {
	if p.pid == "" {
		return false
	}

	if len(p.pid) != 9 {
		return false
	}

	_, err := strconv.Atoi(p.pid)
	if err != nil {
		return false
	}

	return true
}
