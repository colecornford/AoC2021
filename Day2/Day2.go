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

	var depth, position int = 0, 0
	for _, item := range data {
		var command []string = strings.Split(item, " ")
		var distance, _ = strconv.Atoi(command[1])
		switch command[0] {
		case "forward":
			position += distance
		case "up":
			depth -= distance
		case "down":
			depth += distance
		}
	}

	fmt.Println(depth * position)
}

func part2() {
	data := getData()

	var depth, position, aim int = 0, 0, 0
	for _, item := range data {
		var command []string = strings.Split(item, " ")
		var distance, _ = strconv.Atoi(command[1])
		switch command[0] {
		case "forward":
			position += distance
			depth = distance*aim + depth
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}

	fmt.Println(depth * position)

}
