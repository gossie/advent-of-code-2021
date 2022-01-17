package day21

import (
	"math"
)

type player struct {
	position int
	points   int
}

func (p *player) move(steps int) {
	p.position = p.position + steps
	for p.position > 10 {
		p.position -= 10
	}
	p.points += p.position
}

func (p *player) won(limit int) bool {
	return p.points >= limit
}

type gameState struct {
	turn      int
	rolls     int
	die       int
	playerOne player
	playerTwo player
}

type wins struct {
	one int
	two int
}

func PlayTestGame() int {
	game := gameState{playerOne: player{position: 6}, playerTwo: player{position: 9}}

	for !game.playerOne.won(1000) && !game.playerTwo.won(1000) {
		eyes := 0
		for i := 0; i < 3; i++ {
			game.rolls++
			game.die++
			if game.die > 100 {
				game.die -= 100
			}
			eyes += game.die
		}

		if game.turn == 0 {
			game.playerOne.move(eyes)
		} else {
			game.playerTwo.move(eyes)
		}

		game.turn = (game.turn + 1) % 2
	}

	return int(math.Min(float64(game.playerOne.points), float64(game.playerTwo.points))) * game.rolls
}

func playGame(game gameState, gameCache map[gameState]wins, diceFrequencies map[int]int) wins {
	if result, present := gameCache[game]; present {
		return result
	}

	if game.playerOne.won(21) {
		return wins{one: 1}
	}
	if game.playerTwo.won(21) {
		return wins{two: 1}
	}

	totalWins := wins{}
	for die, frequency := range diceFrequencies {
		newState := gameState{turn: (game.turn + 1) % 2, playerOne: player{position: game.playerOne.position, points: game.playerOne.points}, playerTwo: player{position: game.playerTwo.position, points: game.playerTwo.points}}
		if game.turn == 0 {
			newState.playerOne.move(die)
		} else {
			newState.playerTwo.move(die)
		}

		subWins := playGame(newState, gameCache, diceFrequencies)
		totalWins.one += subWins.one * frequency
		totalWins.two += subWins.two * frequency
	}
	gameCache[game] = totalWins
	return totalWins
}

func MultipleUniverses() int64 {
	startState := gameState{playerOne: player{position: 6}, playerTwo: player{position: 9}}

	gameCache := make(map[gameState]wins)

	diceFrequencies := make(map[int]int)
	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				diceFrequencies[r1+r2+r3]++
			}
		}
	}

	totalWins := playGame(startState, gameCache, diceFrequencies)

	return int64(math.Max(float64(totalWins.one), float64(totalWins.two)))
}
