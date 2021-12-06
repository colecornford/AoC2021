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

func getData() []int {
	content, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		panic(err)
	}
	sFish := strings.Split(string(content), ",")
	var iFish []int
	for _, fish := range sFish {
		i, _ := strconv.Atoi(fish)
		iFish = append(iFish, i)
	}
	return iFish
}

func part1() {
	data := getData()
	iterations := 80
	p, sp1, sp10 := "", " ", " "

	for x := 1; x <= iterations; x++ {
		var newSchool []int
		for _, fish := range data {
			newSchool = append(newSchool, fish-1)
		}
		for y, fish := range data {
			if fish == 0 {
				newSchool = append(newSchool, 8)
				newSchool[y] = 6
			}
		}
		if x > 1 {
			p = "s"
			sp1 = ""
		}
		if x > 9 {
			sp10 = ""
		}
		fmt.Printf("After %v%v day%v:  %v %v\n", sp10, x, p, sp1, newSchool)
		data = newSchool
	}
	fmt.Println(len(data))
}

func getDataP2() map[int]int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	sFish := strings.Split(string(content), ",")
	school := make(map[int]int)
	for _, fish := range sFish {
		i, _ := strconv.Atoi(fish)
		school[i]++
	}
	return school
}

func part2() {
	school := getDataP2()
	iterations := 256
	p, sp1, sp10 := "", " ", " "
	for x := 1; x <= iterations; x++ {
		newSchool := make(map[int]int)
		for y := 8; y >= 0; y-- {
			newSchool[y] = school[y+1]
		}
		if school[0] > 0 {
			newSchool[8] = school[0] + newSchool[8]
			newSchool[6] = school[0] + newSchool[6]
		}
		if x > 1 {
			p = "s"
			sp1 = ""
		}
		if x > 9 {
			sp10 = ""
		}
		fmt.Printf("After %v%v day%v:  %v %v\n", sp10, x, p, sp1, newSchool)
		school = newSchool
	}
	count := 0
	for _, fish := range school {
		count += fish
	}
	fmt.Println(count)
}
