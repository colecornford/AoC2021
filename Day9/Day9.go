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

func getData() (grid [][]int) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var intItems [][]int
	lines := strings.Split(string(content), "\n")
	for _, l := range lines {
		var positions []int
		for _, i := range l {
			p, _ := strconv.Atoi(string(i))
			positions = append(positions, p)
		}
		intItems = append(intItems, positions)
	}
	return intItems
}

func part1() {

	data := getData()
	riskTotal := 0

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			checkAbove := y > 0
			checkBelow := y < len(data)-1
			checkLeft := x > 0
			checkRight := x < len(data[y])-1

			if checkAbove {
				if data[y][x] >= data[y-1][x] {
					continue
				}
			}
			if checkBelow {
				if data[y][x] >= data[y+1][x] {
					continue
				}
			}
			if checkLeft {
				if data[y][x] >= data[y][x-1] {
					continue
				}
			}
			if checkRight {
				if data[y][x] >= data[y][x+1] {
					continue
				}
			}
			fmt.Println(data[y][x] + 1)
			riskTotal += data[y][x] + 1
		}
	}

	fmt.Println(riskTotal)

}

func part2() {

	data := getData()

	first, second, third, z := 0, 0, 0, 0

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] >= 9 {
				continue
			}
			z += 11
			data = check(y, x, z, data)
		}
	}
	for _, item := range data {
		for _, it := range item {
			if it < 10 {
				fmt.Print(" ")
			}
			fmt.Print(it)
			fmt.Print("  ")
		}
		fmt.Println()
	}

	m := make(map[int]int)

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] != 9 {
				m[data[y][x]]++
			}
		}
	}

	fmt.Print(m)

	for _, item := range m {
		if item > first {
			third = second
			second = first
			first = item
		} else if item >= second {
			third = second
			second = item
		} else if item >= third {
			third = item
		}
	}
	fmt.Printf("%v %v %v", first, second, third)
	fmt.Println()
	fmt.Println()
	fmt.Println(first * second * third)

}

func check(y, x, val int, data [][]int) [][]int {
	if data[y][x] >= 9 {
		return data
	}
	data[y][x] = val
	if x+1 != len(data[y]) {
		check(y, x+1, val, data)
	}
	if x != 0 {
		check(y, x-1, val, data)
	}
	if y+1 != len(data) {
		check(y+1, x, val, data)
	}
	if y != 0 {
		check(y-1, x, val, data)
	}
	return data
}
