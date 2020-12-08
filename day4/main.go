package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
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
	passports = append(passports, *currentPassport)
	
	part1(passports)
}

func part1(passports []passport) {
	validPassports := 0
	for i, passport := range passports {
		if passport.isValid() {
			validPassports += 1
		}
	}

	fmt.Printf("There are %d valid passports for part 1\n", validPassports)
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

func (p *passport) isValid() bool {
	return p.byr != "" &&
				 p.iyr != "" &&
				 p.eyr != "" &&
				 p.hgt != "" &&
				 p.hcl != "" &&
				 p.ecl != "" &&
				 p.pid != ""
}