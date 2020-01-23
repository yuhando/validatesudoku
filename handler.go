package main

// ValidateBoard mapping all existing rows and columns on the board
func ValidateBoard(board [][]int) string {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if !CheckingBoardData(board, row, col) { // the result will return false, if there is a value not correct or not pass from checking process
				return "Incorrect"
			}
		}
	}
	// if all checking process pass, the result will return "Correct"
	return "Correct"
}

// CheckingBoardData make sure that value on all rows and columns are unique
func CheckingBoardData(board [][]int, row int, col int) bool {

	// each value of a row, column and 3x3 board ​​will be stored in this model to be checked
	itemSet := ItemsSet{}

	// check row value one by one
	for x := 0; x < 9; x++ { // loop the row from the first index(0) to the last index(8) -> 1-9

		// get current row and if value exist return false
		getCurrentRow := board[row][x]
		if itemSet.Exist(getCurrentRow) {
			return false
		}

		// if not stored into model
		if getCurrentRow > 0 && getCurrentRow <= 9 {
			itemSet.Add(getCurrentRow)
		}

	}

	// clean the model after checking values ​​from the row
	itemSet.Clear()

	// check column value one by one, the checking process is same like checking row value
	for y := 0; y < 9; y++ {
		getCurrentCol := board[y][col]
		if itemSet.Exist(getCurrentCol) {
			return false
		}

		if getCurrentCol > 0 && getCurrentCol <= 9 {
			itemSet.Add(getCurrentCol)
		}
	}

	itemSet.Clear()

	// check 3x3 board value one by one, the process still same like row and column the difference is on just the looping
	for x := 0; x < 3; x++ { // loop row
		for y := 0; y < 3; y++ { // loop column

			getRow := row/3*3 + x
			getCol := col/3*3 + y
			getCurrentBoard3x3 := board[getRow][getCol]

			if itemSet.Exist(getCurrentBoard3x3) {
				return false
			}

			if getCurrentBoard3x3 > 0 && getCurrentBoard3x3 <= 9 {
				itemSet.Add(getCurrentBoard3x3)
			}
		}
	}

	// if all checking processes pass then the results will return true,
	// which means all values ​​from the given(sudoku board) are correct.
	return true
}
