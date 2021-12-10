package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
)

func main() {
	// part1()
	part2()
}

func getData() (data []string) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func part1() {
	m := make(map[string]int)
	m[")"] = 3
	m["]"] = 57
	m["}"] = 1197
	m[">"] = 25137

	data := getData()
	var corruptData []string
	score := 0
	for _, line := range data {
		var p string
		s := make(stack,0)
		for _, bracket := range line {
			if strings.ContainsAny(string(bracket), "{[(<") {
				s = s.Push(string(bracket))
				continue
			} else {
				s, p = s.Pop()
				isCorrupt, expected, actual := matchBracket(p,string(bracket))
				if isCorrupt {
					corruptData = append(corruptData, line)
					fmt.Printf("expected: %v actual: %v \n", expected, actual)
					score += m[actual]
					break
				}
			}
		}
	}
	for _, line := range corruptData {
		fmt.Println(line)
	}
	fmt.Print(score)
}

func part2() {

	data := getData()
	m := make(map[string]int64)
	m["("] = 1
	m["["] = 2
	m["{"] = 3
	m["<"] = 4
	var corruptData, goodData []string
	var score []int64
	for _, line := range data {
		var p string
		lineLength := len(line)
		s := make(stack,0)
		for h, bracket := range line {
			if h >= lineLength -1 {
				goodData = append(goodData, line)
				break
			}
			if strings.ContainsAny(string(bracket), "{[(<") {
				s = s.Push(string(bracket))
				continue
			} else {
				s, p = s.Pop()
				isCorrupt, expected, actual := matchBracket(p,string(bracket))
				if isCorrupt {
					corruptData = append(corruptData, line)
					fmt.Printf("expected: %v actual: %v \n", expected, actual)
					break
				}
			}
			fmt.Printf("%v %v %v \n", h, lineLength -2, line)

		}
	}

	fmt.Println(goodData)
	
	for _, line := range goodData {
		var lineScore int64
		lineScore = 0
		s := make(stack,0)
		for i := 0; i < len(line); i++ {
			if strings.ContainsAny(string(line[i]), "{[(<") {
				s = s.Push(string(line[i]))
			} else {
				s, _ = s.Pop()
			}
		}
		fmt.Print(s)
		for x := len(s)-1; x >= 0 ; x-- {
			fmt.Println(lineScore)
			lineScore = lineScore * 5 + m[string(s[x])]
		}
		fmt.Println(lineScore)
		score = append(score, lineScore)
	}
	sort.Slice(score, func(i, j int) bool { return score[i] < score[j] })
	/* for x := 0; x < 20; x++ {
		score[x] = 0
		score[len(score)-x-1] = 0
	} */

	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1889585934 1921358936 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1889585934 1921358936 
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

	fmt.Println(score)
	fmt.Println("FUCK")
	fmt.Println(score[len(score)/2])
}

type stack []string

func (s stack) Push(v string) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, string) {
    l := len(s)
	if l == 0 {
		return s, ""
	}
    return  s[:l-1], s[l-1]
}

func matchBracket(open, close string) (bool, string, string) {
	if open == "{" && close != "}" {	
		return true, "}", close
	}
	if open == "[" && close != "]" {
		return true, "]", close
	}
	if open == "(" && close != ")"  {
		return true, ")", close
	}
	if open == "<" && close != ">"  {
		return true, ">", close
	} 
	return false, "", ""
}