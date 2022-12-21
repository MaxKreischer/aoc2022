package main

import (
	"encoding/json"
	"fmt"
	"os"
	// "aoc22_lib/arraylib"
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

// (1 for Rock, 2 for Paper, and 3 for Scissors)
//  plus the score for the outcome of the round
//   (0 if you lost, 3 if the round was a draw, and 6 if you won).

type TicTacToeRounds struct {
	TTTRounds [][2]string `json:"turns"`
}

func main() {

	someStruct := TicTacToeRounds{}
	file_name := "./input/input_aoc2202.json"
	fileBytes, _ := os.ReadFile(file_name)

	err := json.Unmarshal(fileBytes, &someStruct)

	if err != nil {
		fmt.Printf("%v", err)
	}

	// strategy_guide := [][2]string{{"A", "Y"}, {"B", "X"}, {"C", "Z"}}
	strategy_guide := someStruct.TTTRounds

	total_score := 0
	for _, turn := range strategy_guide {
		score := game_outcome_scores[turn]
		total_score += score
	}
	fmt.Printf("Total score: %v\n", total_score)

}
