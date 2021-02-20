package sudoku

import "testing"

func TestPrintBoard(t *testing.T) {
	board := [][]int{
		{2, 4, 1, 6, 8, 5, 3, 9, 7},
		{3, 8, 9, 7, 4, 1, 2, 6, 5},
		{6, 7, 5, 2, 3, 9, 4, 8, 1},
		{7, 1, 6, 3, 2, 4, 9, 5, 8},
		{5, 3, 4, 8, 9, 7, 1, 2, 6},
		{8, 9, 2, 1, 5, 6, 7, 3, 4},
		{4, 2, 7, 5, 6, 3, 8, 1, 9},
		{9, 6, 8, 4, 1, 2, 5, 7, 3},
		{1, 5, 3, 9, 7, 8, 6, 4, 2},
	}

	PrintBoard(board)
}