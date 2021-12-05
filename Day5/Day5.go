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

type point struct {
	x int
	y int
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
	var count int = 0
	all := make(map[point]int)

	for _, inputLine := range data {
		var startFin = strings.Split(string(inputLine), " -> ")
		var point1, point2 point
		point1.x, _ = strconv.Atoi(strings.Split(startFin[0], ",")[0])
		point1.y, _ = strconv.Atoi(strings.Split(startFin[0], ",")[1])
		point2.x, _ = strconv.Atoi(strings.Split(startFin[1], ",")[0])
		point2.y, _ = strconv.Atoi(strings.Split(startFin[1], ",")[1])
		if (point1.x != point2.x && point1.y != point2.y) || (point1.x == point2.x && point1.y == point2.y) {
			continue
		}
		if point1.x == point2.x {
			for y := min(point1.y, point2.y); y <= max(point1.y, point2.y); y++ {
				var newPoint point
				newPoint.x = point1.x
				newPoint.y = y
				all[newPoint]++
			}
		}
		if point1.y == point2.y {
			for x := min(point1.x, point2.x); x <= max(point1.x, point2.x); x++ {
				var newPoint point
				newPoint.x = x
				newPoint.y = point1.y
				all[newPoint]++
			}
		}
	}
	fmt.Println(all)
	for _, item := range all {
		if item > 1 {
			count++
		}
	}
	fmt.Println(count)
}

func part2() {
	data := getData()
	var count int = 0
	all := make(map[point]int)

	for _, inputLine := range data {
		var startFin = strings.Split(string(inputLine), " -> ")
		var point1, point2 point
		point1.x, _ = strconv.Atoi(strings.Split(startFin[0], ",")[0])
		point1.y, _ = strconv.Atoi(strings.Split(startFin[0], ",")[1])
		point2.x, _ = strconv.Atoi(strings.Split(startFin[1], ",")[0])
		point2.y, _ = strconv.Atoi(strings.Split(startFin[1], ",")[1])

		if point1.x == point2.x {
			for y := min(point1.y, point2.y); y <= max(point1.y, point2.y); y++ {
				var newPoint point
				newPoint.x = point1.x
				newPoint.y = y
				all[newPoint]++
			}
		}
		if point1.y == point2.y {
			for x := min(point1.x, point2.x); x <= max(point1.x, point2.x); x++ {
				var newPoint point
				newPoint.x = x
				newPoint.y = point1.y
				all[newPoint]++
			}
		}
		if point1.x != point2.x && point1.y != point2.y {
			var dist int = max(point1.x, point2.x) - min(point1.x, point2.x)
			for x := 0; x <= dist; x++ {
				var newPoint point
				if point1.x-point2.x <= 0 {
					newPoint.x = point1.x + x
					if point1.y-point2.y <= 0 {
						newPoint.y = point1.y + x
					} else {
						newPoint.y = point1.y - x
					}
				} else {
					newPoint.x = point1.x - x
					if point1.y-point2.y <= 0 {
						newPoint.y = point1.y + x
					} else {
						newPoint.y = point1.y - x
					}
				}
				fmt.Println(newPoint)
				all[newPoint]++
			}
		}
	}
	fmt.Println(all)
	for _, item := range all {
		if item > 1 {
			count++
		}
	}
	fmt.Println(count)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
