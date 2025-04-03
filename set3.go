package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Relationship status:")
	fmt.Println(relationshipStatus("@chums", "@jobenilagan", socialGraph))

	fmt.Println("Tic tac toe:")
	fmt.Println(ticTacToe(board))

	fmt.Println("ETA:")
	fmt.Println(eta("dlsu", "upd", routeMap))
}

func relationshipStatus(fromMember string, toMember string, socialGraph map[string]map[string]string) string {
	// subject and object users
	subjectFollows := false
	objectFollows := false

	// split subject following list by user
	following := strings.Split(socialGraph[fromMember]["following"], ",")
	// iterate through following list
	for _, user := range following {
		// if object user found, set to true
		if user == toMember {
			subjectFollows = true
			break
		}
	}

	// repeat process for object user
	following = strings.Split(socialGraph[toMember]["following"], ",")
	for _, user := range following {
		if user == fromMember {
			objectFollows = true
			break
		}
	}

	// return appropriate status based on bool conditions
	if subjectFollows && objectFollows {
		return "friends"
	} else if subjectFollows {
		return "follower"
	} else if objectFollows {
		return "followed by"
	}

	// otherwise, return no relationship
	return "no relationship"
}

func ticTacToe(board [][]string) string {
	// get size of the board
	n := len(board)
	// arrays for storing diagonals
	diag_down, diag_up := make([]string, n), make([]string, n)

	// helper function to check for winning lines (horizontal, vertical, diagonals)
	checkLine := func(line []string) string {
		// get the first symbol in a line
		symbol := line[0]

		// if blank, can't win so return empty string
		if symbol == "" {
			return ""
		}

		// iterate through line, if any symbol does not match can't win so return empty string
		for _, point := range line {
			if point != symbol {
				return ""
			}
		}

		// otherwise, return the winning symbol
		return symbol
	}

	// iterate through each row and column
	for i := 0; i < n; i++ {
		// array for columns
		col := make([]string, n)

		// populate column array with respective symbols
		for j := 0; j < n; j++ {
			col[j] = board[j][i]
		}

		// check each row, return winner if not blank
		if winner := checkLine(board[i]); winner != "" {
			return winner
		}

		// do same for columns
		if winner := checkLine(col); winner != "" {
			return winner
		}
	}

	// populate diagonal arrays with values
	for i := 0; i < n; i++ {
		// downward and upward pattern across the board
		diag_down[i] = board[i][i]
		diag_up[i] = board[i][n-i-1]
	}

	// check downward diagonal
	if winner := checkLine(diag_down); winner != "" {
		return winner
	}

	// check upward diagonal
	if winner := checkLine(diag_up); winner != "" {
		return winner
	}

	// if no winners found, return no winner
	return "NO WINNER"
}

func eta(firstStop string, secondStop string, routeMap map[string]map[string]int) int {
	// initialize running total and current stop
	total := 0
	current := firstStop

	// if first stop and last stop identical, return 0
	if firstStop == secondStop {
		return 0
	}

	for {
		// create variable for next stop
		var next string
		// for each route & associated travel time:
		for route, val := range routeMap {
			// split list of stops
			stops := strings.Split(route, ",")
			if stops[0] == current {
				// set next to next stop and increment total
				next = stops[1]
				total += val["travel_time_mins"]
				break
			}
		}

		// if destination reached, return total
		if next == secondStop {
			return total
		}

		// set current stop to next stop
		current = next
	}
}
