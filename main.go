package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	os.Exit(_main())
}

type sudoku struct {
	board [][]int
}

func (s *sudoku) print() {
	for i := range s.board {
		if i%3 == 0 && i != 0 {
			fmt.Println("- - - - - - - - - - - -")
		}

		for j := range s.board[i] {
			if j%3 == 0 && j != 0 {
				fmt.Printf(" | ")
			}

			if j == 8 {
				fmt.Printf("%d\n", s.board[i][j])
			} else {
				fmt.Printf("%d ", s.board[i][j])
			}
		}
	}
}

func (s *sudoku) findNextEmpty() (int, int, bool) {
	for i := range s.board {
		for j := range s.board[i] {
			if s.board[i][j] == 0 {
				return i, j, true
			}
		}
	}

	return -1, -1, false
}

func (s *sudoku) solve() bool {
	row, col, found := s.findNextEmpty()
	if !found {
		return true
	}

	for i := 1; i <= 9; i++ {
		if s.valid(row, col, i) {
			s.board[row][col] = i
			if s.solve() {
				return true
			}

			s.board[row][col] = 0
		}
	}

	return false
}

func (s *sudoku) valid(row int, col int, val int) bool {
	// check row
	for i := range s.board[row] {
		if s.board[row][i] == val {
			return false
		}
	}

	// check column
	for i := range s.board {
		if s.board[i][col] == val {
			return false
		}
	}

	// check box
	boxX := (col / 3) * 3
	boxY := (row / 3) * 3

	for i := boxX; i < boxX+3; i++ {
		for j := boxY; j < boxY+3; j++ {
			if s.board[j][i] == val {
				return false
			}
		}
	}

	return true
}

func _main() int {
	var board [][]int
	scanner := bufio.NewScanner(os.Stdin)
	rows := 0
	cols := -1
	for scanner.Scan() {
		line := scanner.Text()
		var nums []int
		for i, str := range strings.Split(line, ",") {
			num, err := strconv.ParseInt(str, 10, 8)
			if err != nil {
				fmt.Printf("invalid input %s(row: %d, column: %d): err:%v", str, rows, i, err)
				return 1
			}

			if !(num >= 0 && num <= 9) {
				fmt.Printf("invalid input %d(row: %d, column: %d): each number must be 0<=n<=9", num, rows, i)
				return 1
			}

			nums = append(nums, int(num))
		}

		if cols != -1 {
			if len(nums) != cols {
				fmt.Println("All rows must have same columns")
				return 1
			}
		}

		cols = len(nums)
		board = append(board, nums)
		rows++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return 1
	}

	game := sudoku{
		board: board,
	}

	fmt.Println("## Input")
	game.print()
	fmt.Println("")

	game.solve()

	fmt.Println("## Answer")
	game.print()

	return 0
}
