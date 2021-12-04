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
	content, err := ioutil.ReadFile("jay.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func part1() {
	data := getData()

	// var powerConsumption, gammaRate, epsilonRate int = 0,"",""
	var gammaRate, epsilonRate string = "", ""

	var position int = 0
	for {
		var countZero, countOne int = 0, 0

		for _, item := range data {
			if item[position:position+1] == "0" {
				countZero++
			} else {
				countOne++
			}
		}
		if countZero > countOne {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
		if position == len(data[0])-1 {
			g, _ := strconv.ParseInt(gammaRate, 2, 64)
			e, _ := strconv.ParseInt(epsilonRate, 2, 64)
			fmt.Println(g * e)
			break
		} else {
			position++
		}
	}

}

func part2() {
	oldData := getData()
	var carbonScrub, oxyGen string = "", ""
	var position int = 0
	for {
		var newData []string
		var countZero, countOne int = 0, 0
		for _, item := range oldData {
			if item[position:position+1] == "0" {
				countZero++
			} else {
				countOne++
			}
		}
		if countZero > countOne {
			for _, item := range oldData {
				if item[position:position+1] == "0" {
					newData = append(newData, item)
				}
			}
		} else {
			for _, item := range oldData {
				if item[position:position+1] == "1" {
					newData = append(newData, item)
				}
			}
		}
		position++
		oldData = newData
		if len(newData) == 1 {
			carbonScrub = newData[0]
			break
		}
	}
	oldData = getData()
	position = 0
	for {
		var newData []string
		var countZero, countOne int = 0, 0
		for _, item := range oldData {
			if item[position:position+1] == "0" {
				countZero++
			} else {
				countOne++
			}
		}
		if countZero > countOne {
			for _, item := range oldData {
				if item[position:position+1] == "1" {
					newData = append(newData, item)
				}
			}
		} else {
			for _, item := range oldData {
				if item[position:position+1] == "0" {
					newData = append(newData, item)
				}
			}
		}
		position++
		oldData = newData
		if len(newData) == 1 {
			oxyGen = newData[0]
			break
		}
	}
	c, _ := strconv.ParseInt(carbonScrub, 2, 64)
	o, _ := strconv.ParseInt(oxyGen, 2, 64)
	fmt.Println("110011010000")
	fmt.Println(oxyGen)
	fmt.Println("001100100110")
	fmt.Println(carbonScrub)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(c * o)
}
