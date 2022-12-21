package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var game_outcome_scores = map[[2]string]int{
	[2]string{"A", "X"}: 1 + 3,
	[2]string{"A", "Y"}: 2 + 6,
	[2]string{"A", "Z"}: 3 + 0,
	[2]string{"B", "X"}: 1 + 0,
	[2]string{"B", "Y"}: 2 + 3,
	[2]string{"B", "Z"}: 3 + 6,
	[2]string{"C", "X"}: 1 + 6,
	[2]string{"C", "Y"}: 2 + 0,
	[2]string{"C", "Z"}: 3 + 3,
}

type GamePartialScoreKey struct {
	points     int
	their_move string
}

var desired_outcomes_and_their_move_to_partial_score = map[GamePartialScoreKey]int{
	{0, "A"}: 3,
	{3, "A"}: 1,
	{6, "A"}: 2,
	{0, "B"}: 1,
	{3, "B"}: 2,
	{6, "B"}: 3,
	{0, "C"}: 2,
	{3, "C"}: 3,
	{6, "C"}: 1,
}

func get_points_by_strategy(moves [2]string) int {
	var move_theirs string = moves[0]
	var desired_outcome string = moves[1]
	var round_score int = 0
	if desired_outcome == "X" {
		round_score += 0

	} else if desired_outcome == "Y" {
		round_score += 3
	} else if desired_outcome == "Z" {
		round_score += 6
	} else {
		panic("Desired outcome not in supported outcomes!")
	}
	gp_key := GamePartialScoreKey{round_score, move_theirs}
	round_score += desired_outcomes_and_their_move_to_partial_score[gp_key]
	return round_score
}

// (1 for Rock, 2 for Paper, and 3 for Scissors)
//  plus the score for the outcome of the round
//   (0 if you lost, 3 if the round was a draw, and 6 if you won).

type TicTacToeRounds struct {
	TTTRounds [][2]string `json:"turns"`
}

func main() {

	someStruct := TicTacToeRounds{}
	// file_name := "./input/test_in.json"
	file_name := "./input/input_aoc2202.json"
	fileBytes, _ := os.ReadFile(file_name)

	err := json.Unmarshal(fileBytes, &someStruct)

	if err != nil {
		fmt.Printf("%v", err)
	}

	strategy_guide := someStruct.TTTRounds

	total_score := 0
	for _, turn := range strategy_guide {
		// score := game_outcome_scores[turn]
		score := get_points_by_strategy(turn)
		total_score += score
	}
	fmt.Printf("Total score: %v\n", total_score)

}
