package day21

import (
	"math"
)

func PlayTestGame() int {
	playerOnePosition := 6
	playerTwoPosition := 9
	playerOnePoints := 0
	playerTwoPoints := 0

	turn := 0

	rolls := 0
	die := 0

	for playerOnePoints < 1000 && playerTwoPoints < 1000 {
		eyes := 0
		for i := 0; i < 3; i++ {
			rolls++
			die++
			if die > 100 {
				die -= 100
			}
			eyes += die
		}

		if turn == 0 {
			playerOnePosition = playerOnePosition + eyes
			for playerOnePosition > 10 {
				playerOnePosition -= 10
			}
			playerOnePoints += playerOnePosition
		} else {
			playerTwoPosition = playerTwoPosition + eyes
			for playerTwoPosition > 10 {
				playerTwoPosition -= 10
			}
			playerTwoPoints += playerTwoPosition
		}

		turn = (turn + 1) % 2
	}

	return int(math.Min(float64(playerOnePoints), float64(playerTwoPoints))) * rolls
}
