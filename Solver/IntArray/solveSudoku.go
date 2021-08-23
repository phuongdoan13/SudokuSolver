package IntArray


const length int = 3
const area int = length * length
var rows [area][area + 1]int
var cols [area][area + 1]int
var boxes [area][area + 1]int
var isSolved bool = false

func SolveSudoku(board [][]int)  {
	/*
		Trigger function to solve the sudoku
	 */
	for i := 0; i < area; i++ {
		for j := 0; j < area; j++ {
			var num  = board[i][j]
			if num != -1 {
				placeNumber(num, i, j, board)
			}
		}
	}
	backtrack(0, 0, board)
}



func placeNextNumbers(row int, col int, board [][]int){
	/*
	   Call backtrack function in recursion
	   to continue to place numbers
	   till the moment we have a solution
	*/

	if (col == area- 1) && (row == area- 1) { // If we are at the final cell in the board
		isSolved = true
	}else{                 // If not
		if col == area- 1{ // If we are at the end of a row
			backtrack(row + 1, 0, board)
		}else{ // If we are at the end of a col
			backtrack(row, col + 1, board)
		}


	}
}

func backtrack(row int, col int, board [][]int){
	/*
	   Backtracking
	*/
	if board[row][col] == -1 {
		for d := 1; d <= 9; d++ {
			if isPlaceable(d, row, col) {
				placeNumber(d, row, col, board)
				placeNextNumbers(row, col, board)
				if !isSolved {
					removeNumber(d, row, col, board)
				}
			}
		}
	}else{
		placeNextNumbers(row, col, board)
	}
}

func placeNumber(d int, row int, col int, board [][]int){
	/*
	   Add number to a cell
	*/
	var box int = (row /length) *length + col /length
	rows[row][d]++
	cols[col][d]++
	boxes[box][d]++
	board[row][col] = d
}

func removeNumber(d int, row int, col int, board [][]int){
	/*
	   Remove number from a cell
	*/
	var box int = (row /length) *length + col /length
	rows[row][d]--
	cols[col][d]--
	boxes[box][d]--
	board[row][col] = -1
}

func isPlaceable(d int, row int, col int) bool{
	/*
	   Check if it's valid to place the digit d in a cell
	*/
	var box int = (row /length) *length + col /length
	return (rows[row][d] + cols[col][d] + boxes[box][d]) == 0
}

