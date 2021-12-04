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

type bingoBoard struct {
	values [][]string
	hasWon bool
}

func getData() (lines []string) {
	content, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(content), "\n")
}

func part1() {
	data := getData()
	seed := strings.Split(data[0], ",")
	var boards []bingoBoard

	for x := range data {
		var newBoard bingoBoard
		if (x % 6) == 2 {
			for y := 0; y < 5; y++ {
				if x+y >= len(data) {
					continue
				} else {
					line := strings.Split(strings.Trim(strings.ReplaceAll(data[x+y], "  ", " "), " "), " ")
					newBoard.values = append(newBoard.values, line)
				}
			}
			boards = append(boards, newBoard)

		}
	}

	fmt.Println("---------- HERE ARE THE BOARDS ----------")
	for _, item := range boards {
		fmt.Println(item)
		fmt.Println()
	}
	fmt.Println("---------- GAME TIME ----------")

	for _, seedNo := range seed {
		fmt.Println("NUMBER:" + seedNo)
		for _, board := range boards {
			for x, row := range board.values {
				for y, item := range row {
					if seedNo == strings.TrimSpace(item) {
						board.values[x][y] = "*"
						for _, line := range board.values {
							fmt.Println(line)
						}
						fmt.Println()
					}
				}
			}
		}
		fmt.Println("---------- CHECKING FOR BINGO ----------")
		for _, board := range boards {
			if checkBingo(board) {
				fmt.Println("BINGO")
				for _, line := range board.values {
					fmt.Println(line)
				}
				fmt.Println(calculateWin(board, seedNo))
				return
			}
		}
	}
}

func calculateWin(winBoard bingoBoard, seedNo string) (total int) {
	var sum int = 0
	seeeeed, _ := strconv.Atoi(seedNo)
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if winBoard.values[y][x] != "*" {
				val, _ := strconv.Atoi(strings.TrimSpace(winBoard.values[y][x]))
				sum += val
			}
		}
	}
	fmt.Println(sum)
	fmt.Println(seeeeed)
	return sum * seeeeed
}

func checkBingo(board bingoBoard) (bingo bool) {
	for x := 0; x < 5; x++ {
		if board.values[x][0] == "*" && board.values[x][1] == "*" && board.values[x][2] == "*" && board.values[x][3] == "*" && board.values[x][4] == "*" {
			return true
		}
		if board.values[0][x] == "*" && board.values[1][x] == "*" && board.values[2][x] == "*" && board.values[3][x] == "*" && board.values[4][x] == "*" {
			return true
		}
	}
	return false
}

func part2() {
	weHaveALoser := false
	data := getData()
	seed := strings.Split(data[0], ",")
	var boards []bingoBoard

	for x := range data {
		var newBoard bingoBoard
		if (x % 6) == 2 {
			for y := 0; y < 5; y++ {
				if x+y >= len(data) {
					continue
				} else {
					line := strings.Split(strings.Trim(strings.ReplaceAll(data[x+y], "  ", " "), " "), " ")
					newBoard.values = append(newBoard.values, line)
				}
			}
			boards = append(boards, newBoard)

		}
	}

	fmt.Println("---------- HERE ARE THE BOARDS ----------")
	for _, item := range boards {
		fmt.Println(item)
		fmt.Println()
	}
	fmt.Println("---------- GAME TIME ----------")

	for _, seedNo := range seed {
		fmt.Println("NUMBER:" + seedNo)
		for _, board := range boards {
			for x, row := range board.values {
				for y, item := range row {
					if seedNo == strings.TrimSpace(item) {
						board.values[x][y] = "*"
						for _, line := range board.values {
							fmt.Println(line)
						}
						fmt.Println()
					}
				}
			}
		}
		fmt.Println("---------- CHECKING FOR BINGO ----------")
		count := 0
		for x, board := range boards {
			if checkBingo(board) {
				if weHaveALoser && !board.hasWon {
					fmt.Println(calculateWin(board, seedNo))
					return
				}
				boards[x].hasWon = true
			}
			if boards[x].hasWon {
				count++
			}
		}
		if len(boards)-count == 1 {
			weHaveALoser = true
		}
	}
}
