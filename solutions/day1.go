package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("inputs/day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	var measurements = strings.Split(string(dat), "\n")
	fmt.Println("Answer part 1: " + strconv.Itoa(part1(measurements)))
	fmt.Println("Answer part 2: " + strconv.Itoa(part2(measurements)))
}

func part1(measurements []string) int {
	var up int = 0

	for i := 0; i < len(measurements)-1; i++ {
		var current, err1 = strconv.Atoi(measurements[i])
		var next, err2 = strconv.Atoi(measurements[i+1])
		if err1 != nil {
			panic(err1)
		}
		if err2 != nil {
			panic(err2)
		}
		if next > current {
			up++
		}
	}

	return up
}

func part2(measurements []string) int {
	var up int = 0
	var previous int = 0

	for i := 0; i < len(measurements)-2; i++ {
		var x1, err1 = strconv.Atoi(measurements[i])
		var x2, err2 = strconv.Atoi(measurements[i+1])
		var x3, err3 = strconv.Atoi(measurements[i+2])
		if err1 != nil {
			panic(err1)
		}
		if err2 != nil {
			panic(err2)
		}
		if err3 != nil {
			panic(err3)
		}
		var sum = x1 + x2 + x3
		if previous != 0 && sum > previous {
			up++
		}
		previous = sum
	}

	return up
}
