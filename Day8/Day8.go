package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func getData() (lines []string) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func part1() {

	data := getData()
	sum := 0
	for _, line := range data {
		digits := strings.Split(strings.Split(line, " | ")[1], " ")
		for _, digit := range digits {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				sum++
			}
		}
		fmt.Println(sum)
	}
}

func part2() {
	data := getData()
	sum := 0
	for _, line := range data {
		patterns := strings.Split(strings.Split(line, " | ")[0], " ")
		digits := strings.Split(strings.Split(line, " | ")[1], " ")
		m := make(map[string]int)
		var one, four, six string
		var deriveFive, deriveSix []string
		for _, pattern := range patterns {
			val := 100
			switch len(pattern) {
			case 2:
				val = 1
				one = pattern
			case 3:
				val = 7
			case 4:
				val = 4
				four = pattern
			case 5:
				// 2, 3, 5
				deriveFive = append(deriveFive, pattern)
			case 6:
				// 0, 6, 9
				deriveSix = append(deriveSix, pattern)
			case 7:
				val = 8
			}
			m[pattern] = val
		}

		for _, pattern := range deriveFive {
			for x, char := range one {
				if !strings.Contains(pattern, string(char)) {
					break
				}
				if x == 1 {
					m[pattern] = 3
				}
			}
		}

		for p, pattern := range deriveSix {
			for x, char := range four {
				if !strings.Contains(pattern, string(char)) {
					break
				}
				if x == 3 {
					fmt.Println(pattern)
					deriveSix = append(deriveSix[:p], deriveSix[p+1:]...)
					m[pattern] = 9
				}
			}
		}

		for _, pattern := range deriveSix {
			for x, char := range one {
				if !strings.Contains(pattern, string(char)) {
					break
				}
				if x == 1 {
					fmt.Println(pattern)
					m[pattern] = 0
				}
			}
		}

		for _, pattern := range deriveSix {

			if m[pattern] == 9 || m[pattern] == 0 {
				continue
			} else {
				m[pattern] = 6
				six = pattern
				break
			}
		}

		for _, pattern := range deriveFive {
			for x, char := range pattern {
				if !strings.Contains(six, string(char)) {
					break
				}
				if x == 4 {
					m[pattern] = 5
				}
			}
		}

		for _, pattern := range patterns {
			if m[pattern] == 100 {
				m[pattern] = 2
			}
		}

		var z []int
		var zStr string
		for _, digit := range digits {
			for _, pattern := range patterns {
				if len(digit) == len(pattern) {
					for x, char := range digit {
						if !strings.Contains(pattern, string(char)) {
							break
						}
						if x == len(digit)-1 {
							z = append(z, m[pattern])
						}
					}
				}
			}

			for _, item := range z {
				t := strconv.Itoa(item)
				zStr += t
			}
			z = make([]int, 0)
		}
		fmt.Println(m)
		fmt.Println(zStr)
		a, _ := strconv.Atoi(zStr)
		sum += a
		//fmt.Println(sum)
	}
	fmt.Println(sum)
}
