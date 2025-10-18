package main

import "fmt"

type chessBoard = [8][8]byte
type placePieceFn = func(board chessBoard) chessBoard

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

func placeBlackPawns(board chessBoard) chessBoard {
	rowNumber := 1
	for column := range board[rowNumber] {
		board[rowNumber][column] = 'p'
	}
	return board
}
func placeBlackRooks(board chessBoard) chessBoard {
	board[0][0] = 'r'
	board[0][7] = 'r'
	return board
}
func placeBlackKnights(board chessBoard) chessBoard {
	board[0][1] = 'n'
	board[0][6] = 'n'
	return board
}
func placeBlackBishops(board chessBoard) chessBoard {
	board[0][2] = 'b'
	board[0][5] = 'b'
	return board
}
func placeBlackQueen(board chessBoard) chessBoard {
	board[0][3] = 'q'
	return board
}
func placeBlackKing(board chessBoard) chessBoard {
	board[0][4] = 'k'
	return board
}

func placeWhitePawns(board chessBoard) chessBoard {
	rowNumber := 6
	for column := range board[rowNumber] {
		board[rowNumber][column] = 'P'
	}
	return board
}
func placeWhiteRooks(board chessBoard) chessBoard {
	board[7][0] = 'R'
	board[7][7] = 'R'
	return board
}
func placeWhiteKnights(board chessBoard) chessBoard {
	board[7][1] = 'N'
	board[7][6] = 'N'
	return board
}
func placeWhiteBishops(board chessBoard) chessBoard {
	board[7][2] = 'B'
	board[7][5] = 'B'
	return board
}
func placeWhiteQueen(board chessBoard) chessBoard {
	board[7][3] = 'Q'
	return board
}
func placeWhiteKing(board chessBoard) chessBoard {
	board[7][4] = 'K'
	return board
}

func main() {
	var board chessBoard
	placePieceFunctions := [...]placePieceFn{
		placeWhitePawns,
		placeWhiteRooks,
		placeWhiteKnights,
		placeWhiteBishops,
		placeWhiteQueen,
		placeWhiteKing,
		placeBlackPawns,
		placeBlackRooks,
		placeBlackKnights,
		placeBlackBishops,
		placeBlackQueen,
		placeBlackKing,
	}
	for _, placePieceFunction := range placePieceFunctions {
		board = placePieceFunction(board)
	}
	drawBoard(board)
}
