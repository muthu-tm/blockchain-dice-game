package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// Game - chaincode interface for Dice Game
type Game struct {
}

// Init is called during Instantiate transaction after the chaincode container
// has been established for the first time, allowing the chaincode to
// initialize its internal data
func (cc *Game) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke is called to update or query the ledger in a proposal transaction.
// Updated state variables are not committed to the ledger until the
// transaction is committed.
func (cc *Game) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()

	var result []byte
	var err error

	if function == SETUPGAME {
		result, err = setupGame(stub, args)
	} else if function == PLAY {
		result, err = play(stub, args)
	} else if function == GETPLAYERPOINTS {
		result, err = getPlayerPoints(stub, args)
	} else if function == GETPLAYERTURN {
		result, err = getPlayerTurn(stub, args)
	} else if function == "" {
		err = errors.New("Chaincode invoke function name should not be empty")
	} else {
		err = errors.New("Invalid chaincode invoke function name")
	}

	if err != nil {
		fmt.Println("Error occured on chaincode invoke: - ", err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(result)
}

func setupGame(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("Incorrect number of arguments to SETUPGAME! Expected 3")
	}

	gameID := args[0]
	firstPlayerID := args[1]
	secondPlayerID := args[2]

	existingGame, err := IsExistingGame(stub, gameID)

	if existingGame {
		return nil, fmt.Errorf("Found an existing game for the given gameID")
	}

	game := NewGame(gameID, firstPlayerID, secondPlayerID)
	// convert data into byte array and update the ledger data
	gameAsBytes, _ := json.Marshal(game)
	err = PutState(stub, gameID, gameAsBytes)
	if err != nil {
		return nil, err
	}

	return json.Marshal(game)
}

func play(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 4 {
		return nil, fmt.Errorf("Incorrect number of arguments to PLAY! Expected 4")
	}

	gameID := args[0]
	playerID := args[1]
	pointsEarned := args[2]
	isHold := args[3]

	existingGame, err := IsExistingGame(stub, gameID)
	if !existingGame {
		return nil, fmt.Errorf("Not Found an existing game for the given gameID")
	}

	game, err := GetGameByID(stub, gameID)
	if err != nil {
		return nil, err
	}

	pointsAsInt, _ := strconv.Atoi(pointsEarned)
	if (playerID == game.FirstPlayer && game.PlayerTurn == 1) || (playerID == game.SecondPlayer && game.PlayerTurn == 2) {
		if pointsAsInt == 1 || isHold == "hold" {
			game.TotalRounds = game.TotalRounds + 1
			playerPoints := game.PlayerPoints[playerID] + game.CurrentTurnTotal
			game.PlayerPoints[playerID] = playerPoints
			game.CurrentTurnTotal = 0

			if playerID == game.FirstPlayer {
				game.PlayerTurn = 2
			} else {
				game.PlayerTurn = 1
			}
		} else {
			game.CurrentTurnTotal = game.CurrentTurnTotal + pointsAsInt
		}
	} else {
		return nil, fmt.Errorf("This is not your Turn")
	}

	// convert game data into byte array and update the ledger data
	gameAsBytes, _ := json.Marshal(game)
	err = PutState(stub, gameID, gameAsBytes)
	if err != nil {
		return nil, err
	}

	return json.Marshal(game)
}

func getPlayerPoints(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("Incorrect number of arguments for GETPLAYERPOINTS! Expected 2")
	}

	gameID := args[0]
	playerID := args[1]

	game, err := GetGameByID(stub, gameID)
	if err != nil {
		return nil, err
	}

	playerPoints := game.PlayerPoints[playerID]
	pointsAsBytes := []byte(strconv.Itoa(playerPoints))

	return pointsAsBytes, nil
}

func getPlayerTurn(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Incorrect number of arguments for GETPLAYERTURN! Expected 1")
	}

	gameID := args[0]

	game, err := GetGameByID(stub, gameID)
	if err != nil {
		return nil, err
	}

	pointsAsBytes := []byte(strconv.Itoa(game.PlayerTurn))

	return pointsAsBytes, nil
}

func main() {
	if err := shim.Start(new(Game)); err != nil {
		fmt.Println("Error creating Dice Game ChainCode: ", err)
	} else {
		fmt.Println("Dice Game ChainCode was created successfully")
	}
}
