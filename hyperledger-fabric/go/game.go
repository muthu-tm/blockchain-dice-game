package main

// Dice - Game struct with the game details. It will be saved as json string in the ledger
type Dice struct {
	GameID           string         `json:"gameID"`
	FirstPlayer      string         `json:"firstPlayer"`
	SecondPlayer     string         `json:"secondPlayer"`
	TotalRounds      int            `json:"totalRounds"`
	PlayerPoints     map[string]int `json:"gamePoints"`
	CurrentTurnTotal int            `json:"currentTurnTotal"`
	PlayerTurn       int            `json:"playerTurn"`
}

func newDice(gameID, player1, player2 string) (dice Dice) {
	dice.GameID = gameID
	dice.FirstPlayer = player1
	dice.SecondPlayer = player2
	dice.TotalRounds = 0
	dice.CurrentTurnTotal = 0
	dice.PlayerTurn = 1

	return
}

// NewGame - Setup a new Dice Game
func NewGame(gameID, firstPlayer, secondPlayer string) (dice Dice) {
	dice = newDice(gameID, firstPlayer, secondPlayer)

	// initialize the first player points
	dice.PlayerPoints[firstPlayer] = 0
	// initialize the second player points
	dice.PlayerPoints[secondPlayer] = 0

	return
}
