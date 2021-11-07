package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

var WINNING = []string{"Good job!", "It's your lucky day", "You nail it!"}
var DRAW = []string{"Great minds think alike", "Try again", "That was close"}
var LOSSING = []string{"Better luck next time", "To bad =(", "Don't give up"}

type Round struct {
	Message        string `json:"message"`
	ComputerChoise string `json:"computer_choise"`
	PlayerChoise   string `json:"player_choise"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computervalue := rand.Intn(3)
	computerChoise := ""
	playerChoise := ""
	roundResult := ""
	message := ""

	switch playerValue {
	case ROCK:
		playerChoise = "Player chose ROCK"
		break
	case PAPER:
		playerChoise = "Player chose PAPER"
		break
	case SCISSORS:
		playerChoise = "Player chose SCISSORS"
		break
	}

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
		message = DRAW[rand.Intn(3)]
	} else if playerValue == (computervalue+1)%3 {
		roundResult = "Player wins!"
		message = WINNING[rand.Intn(3)]
	} else {
		roundResult = "Computer wins!"
		message = LOSSING[rand.Intn(3)]
	}

	var result Round
	result.ComputerChoise = computerChoise
	result.PlayerChoise = playerChoise
	result.RoundResult = roundResult
	result.Message = message

	return result
}
