package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// opted to use longer variable names for clarity
func round(opponent, player string, line int) int {
	win, lose, draw := 6, 0, 3
	rock, paper, scissors := 1, 2, 3

	switch player {
	case "X":
		if opponent == "A" {
			player = "Z"
		} else if opponent == "B" {
			player = "X"
		} else if opponent == "C" {
			player = "Y"
		}
	case "Y":
		player = opponent
	case "Z":
		if opponent == "A" {
			player = "Y"
		} else if opponent == "B" {
			player = "Z"
		} else if opponent == "C" {
			player = "X"
		}
	}

	choiceMap := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
		"X": "Rock",     // lose
		"Y": "Paper",    // draw
		"Z": "Scissors", // win
	}

	opponentChoice := choiceMap[opponent]
	playerChoice := choiceMap[player]

	opp := strings.ToLower(opponentChoice)
	plr := strings.ToLower(playerChoice)

	if plr == opp {
		switch plr {
		case "rock":
			return draw + rock // 3 + 1
		case "paper":
			return draw + paper // 3 + 2
		case "scissors":
			return draw + scissors // 3 + 3
		}
	} else if plr == "rock" {
		if opp == "paper" {
			return lose + rock // 0 + 1
		} else if opp == "scissors" {
			return win + rock // 6 + 1
		}
	} else if plr == "paper" {
		if opp == "scissors" {
			return lose + paper // 0 + 2
		} else if opp == "rock" {
			return win + paper // 6 + 1
		}
	} else if plr == "scissors" {
		if opp == "rock" {
			return lose + scissors // 0 + 3
		} else if opp == "paper" {
			return win + scissors // 6 + 3
		}
	}

	return 0
}

func main() {
	var sm int
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	line := 1
	for scanner.Scan() {
		if scanner.Text() != "" {
			s := strings.Split(scanner.Text(), " ")

			opp, plr := s[0], s[1]

			r := round(opp, plr, line)
			sm += r
		}
		line += 1
	}

	fmt.Println(sm)
}
