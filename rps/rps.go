package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK          = 0
	PAPER         = 1
	SCISSORS      = 2
	DRAW          = 0
	PLAYER_WINS   = 1
	COMPUTER_WINS = 2
)

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoise string `json:"computer_choise"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computervalue := rand.Intn(3)
	computerChoise := ""
	roundResult := ""
	winner := 0

	switch computervalue {
	case ROCK:
		computerChoise = "Computer chose ROCK"
		break
	case PAPER:
		computerChoise = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoise = "Computer chose SCISSORS"
		break
	}

	if playerValue == computervalue {
		roundResult = "It's a draw"
		winner = 0
	} else if playerValue == (computervalue+1)%3 {
		roundResult = "Player wins!"
		winner = 1
	} else {
		roundResult = "Computer wins!"
		winner = 2
	}

	var result Round
	result.ComputerChoise = computerChoise
	result.RoundResult = roundResult
	result.Winner = winner

	return result
}
