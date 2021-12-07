package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func getData() []int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(content), ",")
	var i []int
	for _, item := range s {
		z, _ := strconv.Atoi(item)
		i = append(i, z)
	}
	return i
}

func part1() {
	crabArmy := getData()
	max := 0
	for _, crab := range crabArmy {
		if crab > max {
			max = crab
		}
	}
	fmt.Println(max)
	minFuel := 123456789
	for x := 0; x <= max; x++ {
		totalFuel := 0
		for _, crab := range crabArmy {
			totalFuel += int(math.Abs((float64(crab) - float64(x))))
		}
		if totalFuel < minFuel {
			minFuel = totalFuel
		}
	}
	fmt.Println(minFuel)
}

func part2() {
	crabArmy := getData()
	max := 0
	for _, crab := range crabArmy {
		if crab > max {
			max = crab
		}
	}
	fmt.Println(max)
	minFuel := 123456789
	for x := 0; x <= max; x++ {
		totalFuel := 0
		for _, crab := range crabArmy {
			dist := int(math.Abs((float64(crab) - float64(x))))
			totalFuel += sumLower(dist)
		}
		if totalFuel < minFuel {
			minFuel = totalFuel
		}
	}
	fmt.Println(minFuel)
}

func sumLower(x int) int {
	fuel := 0
	for i := x; i != 0; i-- {
		fuel += i
	}
	return fuel
}
