package main

import "fmt"

type chessBoard = [8][8]byte

func drawBoard(board chessBoard) {
	for _, row := range board {
		for _, cell := range row {
			if cell == 0 {
				cell = ' '
			}
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func placePieces(board chessBoard) chessBoard {
	rowNumber := 1
	for column := range board[rowNumber] {
		board[rowNumber][column] = 'p'
	}
	board[0][0] = 'r'
	board[0][7] = 'r'
	board[0][1] = 'n'
	board[0][6] = 'n'
	board[0][2] = 'b'
	board[0][5] = 'b'
	board[0][3] = 'q'
	board[0][4] = 'k'

	rowNumber = 6
	for column := range board[rowNumber] {
		board[rowNumber][column] = 'P'
	}
	board[7][0] = 'R'
	board[7][7] = 'R'
	board[7][1] = 'N'
	board[7][6] = 'N'
	board[7][2] = 'B'
	board[7][5] = 'B'
	board[7][3] = 'Q'
	board[7][4] = 'K'

	return board
}

func main() {
	var board chessBoard
	board = placePieces(board)
	drawBoard(board)
}
