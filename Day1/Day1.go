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

	var priorNo, count int = 0, 0
	for _, item := range data {
		newNumber, _ := strconv.Atoi(item)
		if priorNo < newNumber {
			count++
		}
		priorNo = newNumber
	}
	fmt.Print(count - 1) // Do this to compensate for the first item having no comparison
}

func part2() {
	data := getData()

	var priorSum, newSum, count int = 0, 0, 0
	for index := range data {
		if index+2 >= len(data) {
			break
		}
		newNumber3, _ := strconv.Atoi(data[index+2])
		newNumber2, _ := strconv.Atoi(data[index+1])
		newNumber1, _ := strconv.Atoi(data[index])

		newSum = newNumber1 + newNumber2 + newNumber3
		if priorSum < newSum {
			count++
		}
		priorSum = newSum
	}
	fmt.Print(count - 1) // Do this to compensate for the first window having no comparison
}
