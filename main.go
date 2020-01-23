package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Starting ... Validate Sudoku listen to http://localhost:8081/")
	http.HandleFunc("/", Sudoku)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

// Sudoku ..
func Sudoku(w http.ResponseWriter, r *http.Request) {
	sudokuBoard := [][]int{
		{5, 3, 2, 9, 8, 6, 7, 4, 1},
		{4, 8, 7, 2, 1, 5, 3, 6, 9},
		{6, 9, 1, 4, 3, 7, 5, 8, 2},
		{3, 2, 5, 1, 7, 4, 8, 9, 6},
		{7, 6, 4, 3, 9, 8, 1, 2, 5},
		{8, 1, 9, 5, 6, 2, 4, 3, 7},
		{1, 5, 6, 8, 2, 3, 9, 7, 4},
		{9, 7, 8, 6, 4, 1, 2, 5, 3},
		{2, 4, 3, 7, 5, 9, 6, 1, 8},
	}

	validateBoard := ValidateBoard(sudokuBoard)
	fmt.Fprintf(w, "The Sudoku Board : \n \n")
	for _, i := range sudokuBoard {
		convertToString := strings.Trim(strings.Replace(fmt.Sprint(i), " ", " ", -1), "[]")
		fmt.Fprintf(w, convertToString+" \n")
	}

	fmt.Fprintf(w, " \n Result : "+validateBoard)

}
