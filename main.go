package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Player 1 = x
// Player 2 = o
// Input format: x y

func printBoard(board [3][3]string) {
	fmt.Printf(" %s | %s | %s \n", "\033[1m"+board[0][0]+"\033[0m", "\033[1m"+board[0][1]+"\033[0m", "\033[1m"+board[0][2]+"\033[0m")
	fmt.Printf("___________ \n")
	fmt.Printf(" %s | %s | %s \n", "\033[1m"+board[1][0]+"\033[0m", "\033[1m"+board[1][1]+"\033[0m", "\033[1m"+board[1][2]+"\033[0m")
	fmt.Printf("___________ \n")
	fmt.Printf(" %s | %s | %s \n", "\033[1m"+board[2][0]+"\033[0m", "\033[1m"+board[2][1]+"\033[0m", "\033[1m"+board[2][2]+"\033[0m")
}
func inputPoint(arr [3][3]string, player uint8) [3][3]string {
	var x, y int
	var msg error
	fmt.Printf("Player %d - Enter a point\n", player)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Sscanf(input, "%d %d", &x, &y)
	x, y, msg = sanitizeInput(input)
	if msg != nil {
		fmt.Println(msg)
		return inputPoint(arr, player)
	}
	if arr[x-1][y-1] != "" {
		fmt.Println("Point already taken")
		return inputPoint(arr, player)
	}

	if player == 1 {
		arr[x-1][y-1] = "x"
		return arr
	} else {
		arr[x-1][y-1] = "o"
		return arr
	}

}
func checkWinner(board [3][3]string) uint8 {
	// 0 = no winner
	// 1 = winner
	// 2 = draw

	// Check for draw
	draw := true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "" {
				draw = false
			}
		}
	}
	if draw {
		return 2
	}

	// Check rows
	for i := 0; i < 3; i++ {
		if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return 1
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if board[0][i] != "" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return 1
		}
	}

	// Check diagonals
	if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return 1
	}

	if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return 1
	}

	return 0
}
func sanitizeInput(input string) (int, int, error) {
	// Split the input into two parts
	inputParts := strings.Fields(input)
	if len(inputParts) != 2 {
		return 0, 0, fmt.Errorf("\033[31m Co-ordinates must contain two values\033[0m")
	}

	// Parse the two values as integers
	val1, err := strconv.Atoi(inputParts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("\033[31mInvalid input: %s\033[0m", inputParts[0])
	}
	val2, err := strconv.Atoi(inputParts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("\033[31mInvalid input: %s\033[0m", inputParts[1])
	}

	// Check ranges
	if val1 < 1 || val1 > 3 || val2 < 1 || val2 > 3 {
		return 0, 0, fmt.Errorf("\033[31mCo-ordinates must be between 1 and 3\033[0m")
	}
	return val1, val2, nil
}

func main() {
	var board [3][3]string
	for {

		board = inputPoint(board, 1)
		printBoard(board)
		if checkWinner(board) == 1 {
			fmt.Println("\033[33mPlayer 1 wins!\033[0m")
			break
		} else if checkWinner(board) == 2 {
			fmt.Println("\033[33mDraw!\033[0m")
			break
		}

		board = inputPoint(board, 2)
		printBoard(board)
		if checkWinner(board) == 1 {
			fmt.Println("\033[33mPlayer 2 wins!\033[0m")
			break
		} else if checkWinner(board) == 2 {
			fmt.Println("\033[33mDraw!\033[0m")
			break
		}

	}
}
