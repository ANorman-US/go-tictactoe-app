package main

import (
	"fmt"
)

type board struct {
	state         [][]byte
	player_x_turn bool
}

func (b board) get_tile() byte {
	if b.player_x_turn {
		return 'X'
	}
	return 'O'
}

// returns board and bool (true = move made)
func (b board) make_move(index_row int, index_col int) (board, bool) {

	if index_col > 3 || index_col < 1 {
		fmt.Println("Column index out of range")
		return b, false
	}
	if index_row > 3 || index_row < 1 {
		fmt.Println("Row index out of range")
		return b, false
	}
	if b.state[index_row-1][index_col-1] != '_' {
		fmt.Println("Invalid Move")
		return b, false
	}

	b.state[index_row-1][index_col-1] = b.get_tile()
	return b, true

}

func (b board) print_board() {
	fmt.Println("  123")
	for row_num, row := range b.state {
		fmt.Println(row_num+1, string(row))
	}
}

func (b board) is_end_state() bool {

	for row := range 3 {
		if b.state[row][0] == b.state[row][1] && b.state[row][0] == b.state[row][2] && b.state[row][0] != '_' {
			return true
		}
	}

	for col := range 3 {
		if b.state[0][col] == b.state[1][col] && b.state[0][col] == b.state[2][col] && b.state[0][col] != '_' {
			return true
		}
	}

	if b.state[0][0] == b.state[1][1] && b.state[0][0] == b.state[2][2] && b.state[0][0] != '_' {
		return true
	}

	if b.state[0][2] == b.state[1][1] && b.state[0][2] == b.state[2][0] && b.state[0][2] != '_' {
		return true
	}

	return false
}
