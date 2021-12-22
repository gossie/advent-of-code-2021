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
	return p.points > limit
}

type gameState struct {
	turn        int
	rolls       int
	die         int
	currentEyes int
	playerOne   player
	playerTwo   player
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

func playGame(game gameState) (int, int, bool) {
	if game.currentEyes > 0 && game.rolls%3 == 0 {
		if game.turn == 0 {
			game.playerOne.move(game.currentEyes)
		} else {
			game.playerTwo.move(game.currentEyes)
		}

		game.turn = (game.turn + 1) % 2
		game.currentEyes = 0

		if game.playerOne.won(21) {
			return 1, 0, true
		} else if game.playerTwo.won(21) {
			return 0, 1, true
		} else {
			return playGame(game)
		}
	} else {
		game.rolls++
		game.currentEyes += 1
		winsOneOne, winsTwoOne, oneMoveLetToWin := playGame(game)

		winsOneTwo, winsTwoTwo, twoMoveLetToWin := winsOneOne, winsTwoOne, oneMoveLetToWin
		if !oneMoveLetToWin {
			game.currentEyes += 1
			winsOneTwo, winsTwoTwo, twoMoveLetToWin = playGame(game)
		}

		winsOneThree, winsTwoThree := winsOneTwo, winsTwoTwo
		if !twoMoveLetToWin {
			game.currentEyes += 1
			winsOneThree, winsTwoThree, _ = playGame(game)
		}

		return winsOneOne + winsOneTwo + winsOneThree, winsTwoOne + winsTwoTwo + winsTwoThree, false
	}
}

func MultipleUniverses() int {
	// startState := gameState{playerOne: player{position: 6}, playerTwo: player{position: 9}}

	// winsOne, winsTwo, _ := playGame(startState)

	dp := [22][22][11][11][2]int{}
	dp[0][0][6][9][0] = 1

	rolls := make([]int, 10)
	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				rolls[r1+r2+r3]++
			}
		}
	}

	w1, w2 := 0, 0

	for s1 := 0; s1 <= 20; s1++ {
		for s2 := 0; s2 <= 20; s2++ {
			for p1 := 1; p1 <= 10; p1++ {
				for p2 := 1; p2 <= 10; p2++ {
					for r := 1; r <= 9; r++ {
						if dp[s1][s2][p1][p2][0] > 0 {
							np1 := p1 + r
							for np1 > 10 {
								np1 -= 10
							}
							ns := s1 + np1
							nc := dp[s1][s2][p1][p2][0] * rolls[r]
							if ns > 20 {
								w1 += nc
							} else {
								dp[ns][s2][np1][p2][1] += nc
							}
						}
						if dp[s1][s2][p1][p2][1] > 0 {
							np2 := p2 + r
							for np2 > 10 {
								np2 -= 10
							}
							ns := s2 + np2
							nc := dp[s1][s2][p1][p2][1] * rolls[r]
							if ns > 20 {
								w2 += nc
							} else {
								dp[s1][ns][p1][np2][0] += nc
							}
						}
					}
				}
			}
		}
	}

	return int(math.Max(float64(w1), float64(w2)))
}
