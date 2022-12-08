package main

import (
	"fmt"
	"os"
)

func main() {

	if 1 < len(os.Args) && len(os.Args) <= 10 {
		board := GetInput(os.Args)

		if board == [9][9]int{} {
			fmt.Printf("Error")
		} else {
			// printBoard(board)
			if backtrack(&board) {
				// fmt.Println("The Sudoku was solved successfully:")
				printBoard(board)
			} else {
				fmt.Printf("Error")
			}
		}

	} else {
		fmt.Printf("Error")
	}

}

func backtrack(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func printBoard(board [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

/*func parseInput(input string) [9][9]int {
	board := [9][9]int{}
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanRunes)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			scanner.Scan()
			i1, _ := strconv.Atoi(scanner.Text())
			board[row][col] = i1
		}
	}
	return board
}*/

func GetInput(input []string) [9][9]int {

	var matrix [9][9]int
	var counter int

	for i, _ := range input {
		if i > 0 {
			for j, col := range input[i] {
				if len(input[i]) == 9 {
					if col == '.' {
						matrix[i-1][j] = 0
					} else if '1' <= col && col <= '9' {
						matrix[i-1][j] = int(col - 48)
						counter++
					} else {
						return [9][9]int{}
					}
				} else {
					return [9][9]int{}
				}
			}
		}
	}
	if counter < 17 {
		return [9][9]int{}
	}
	return matrix
}
