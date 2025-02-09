package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to TicTacToe!")
	var b board = board{
		state: [][]byte{
			{'_', '_', '_'},
			{'_', '_', '_'},
			{'_', '_', '_'},
		},
		player_x_turn: true,
	}

	for {

		fmt.Println("Player to make a move:", string(b.get_tile()))
		b.print_board()

		var num_row, num_col int
		//loop until move made
		for {
			fmt.Println("Enter row and column number, seperated by a space")
			_, err := fmt.Scanln(&num_row, &num_col)

			if err != nil {
				fmt.Println(err)
				continue
			}

			var move_made bool = false
			b, move_made = b.make_move(num_row, num_col)

			if !move_made {
				continue
			}
			break
		}

		//check is goal state
		if b.is_end_state() {
			fmt.Println("Player", string(b.get_tile()), "wins!")
			b.print_board()
			break
		}
		//if not advance turn
		b.player_x_turn = !b.player_x_turn
	}

}
