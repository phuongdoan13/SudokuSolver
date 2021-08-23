package CharArray

const length int = 3
const area int = length * length
var rows [area][area + 1]int
var cols [area][area + 1]int
var boxes [area][area + 1]int
var isSolved bool = false

func SolveSudoku(board [][]byte)  {
	for i := 0; i < area; i++ {
		for j := 0; j < area; j++ {
			var num byte = board[i][j]
			if num != '.' {
				var d int = int(num - '0')

				placeNumber(d, i, j, board)
			}
		}
	}
	backtrack(0, 0, board)
}

func placeNumber(d int, row int, col int, board [][]byte){
	/*
	   Add number to a cell
	*/
	var box int = (row /length) *length + col /length
	rows[row][d]++
	cols[col][d]++
	boxes[box][d]++
	board[row][col] = byte(d + 48)
}

func removeNumber(d int, row int, col int, board [][]byte){
	/*
	   Remove number from a cell
	*/
	var box int = (row /length) *length + col /length
	rows[row][d]--
	cols[col][d]--
	boxes[box][d]--
	board[row][col] = '.'
}

func isPlaceable(d int, row int, col int) bool{
	/*
	   Check if it's valid to place the digit d in a cell
	*/
	var box int = (row /length) *length + col /length
	return (rows[row][d] + cols[col][d] + boxes[box][d]) == 0
}

func placeNextNumbers(row int, col int, board [][]byte){
	/*
	   Call backtrack function in recursion
	   to continue to place numbers
	   till the moment we have a solution
	*/

	if (col == area- 1) && (row == area- 1) {
		isSolved = true
	}else{
		if col == area- 1{
			backtrack(row + 1, 0, board)
		}else{
			backtrack(row, col + 1, board)
		}
	}
}

func backtrack(row int, col int, board [][]byte){
	/*
	   Backtracking
	*/
	if board[row][col] == '.' {
		for d := 1; d <= 9; d++ {
			if isPlaceable(d, row, col) {
				placeNumber(d, row, col, board)
				placeNextNumbers(row, col, board)
				if !isSolved{
					removeNumber(d, row, col, board)
				}
			}
		}
	}else{
		placeNextNumbers(row, col, board)
	}
}

