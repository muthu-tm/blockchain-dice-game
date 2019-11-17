package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func Test_setupGame(test *testing.T) {
	fmt.Println("\n Test for SETUP_GAME started!")

	diceCC := new(Game)
	stub := shim.NewMockStub("SetupGame", diceCC)
	if stub == nil {
		test.Fatalf("MockStub creation failed")
	}

	gameID := "game-11172019-01"
	firstPlayer := "player-test-01"
	secondPlayer := "player-test-02"

	// Setting up Game with correct args and values
	checkInvoke(test, stub, [][]byte{[]byte(SETUPGAME), []byte(gameID), []byte(firstPlayer), []byte(secondPlayer)})
	// Passing incorrect number of function arguments to check failure
	checkBadInvoke(test, stub, [][]byte{[]byte(SETUPGAME), []byte(gameID), []byte(firstPlayer)})
}

func Test_play(test *testing.T) {
	fmt.Println("\n Test for PLAY started!")

	diceCC := new(Game)
	stub := shim.NewMockStub("Play", diceCC)
	if stub == nil {
		test.Fatalf("MockStub creation failed")
	}

	gameID := "game-11172019-01"
	firstPlayer := "player-test-01"
	secondPlayer := "player-test-02"
	stringValue := ""

	// Setting up Game with correct args and values
	checkInvoke(test, stub, [][]byte{[]byte(SETUPGAME), []byte(gameID), []byte(firstPlayer), []byte(secondPlayer)})
	// Play the Game with correct args and values
	checkInvoke(test, stub, [][]byte{[]byte(PLAY), []byte(gameID), []byte(firstPlayer), []byte(strconv.Itoa(5)), []byte(stringValue)})
	// Passing incorrect number of function arguments to check failure
	checkBadInvoke(test, stub, [][]byte{[]byte(PLAY), []byte(gameID), []byte(firstPlayer), []byte(strconv.Itoa(5))})
	// let's change the player turn by sending 1
	checkInvoke(test, stub, [][]byte{[]byte(PLAY), []byte(gameID), []byte(firstPlayer), []byte(strconv.Itoa(1)), []byte(stringValue)})
	// check whether the player turn changed correctly
	checkQuery(test, stub, [][]byte{[]byte(GETPLAYERTURN), []byte(gameID)}, 2)
	// check whether the player point changed correctly; in this case it should be 5
	checkQuery(test, stub, [][]byte{[]byte(GETPLAYERPOINTS), []byte(gameID), []byte(firstPlayer)}, 5)

	// Play for second player and check the above cases
	// Play the Game with correct args and values
	checkInvoke(test, stub, [][]byte{[]byte(PLAY), []byte(gameID), []byte(secondPlayer), []byte(strconv.Itoa(4)), []byte(stringValue)})
	// let's change the player turn by sending "hold"
	checkInvoke(test, stub, [][]byte{[]byte(PLAY), []byte(gameID), []byte(secondPlayer), []byte(strconv.Itoa(2)), []byte("hold")})
	// check whether the player turn changed correctly
	checkQuery(test, stub, [][]byte{[]byte(GETPLAYERTURN), []byte(gameID)}, 1)
	// check whether the player point changed correctly; in this case it should be 4
	checkQuery(test, stub, [][]byte{[]byte(GETPLAYERPOINTS), []byte(gameID), []byte(secondPlayer)}, 4)

}

func Test_getPlayerPoints(test *testing.T) {
	fmt.Println("\n Test for GET_PLAYER_POINTS started!")
	diceCC := new(Game)
	stub := shim.NewMockStub("GetPlayerPoints", diceCC)
	if stub == nil {
		test.Fatalf("MockStub creation failed")
	}

	gameID := "game-11172019-01"
	firstPlayer := "player-test-01"
	secondPlayer := "player-test-02"

	// Setting up Game with correct args and values
	checkInvoke(test, stub, [][]byte{[]byte(SETUPGAME), []byte(gameID), []byte(firstPlayer), []byte(secondPlayer)})
	// Invoking getPlayerPoints function with correct arguments
	checkInvoke(test, stub, [][]byte{[]byte(GETPLAYERPOINTS), []byte(gameID), []byte(firstPlayer)})
	// Passing incorrect number of function arguments to check failure
	checkBadInvoke(test, stub, [][]byte{[]byte(GETPLAYERPOINTS), []byte(gameID)})
}

func Test_getPlayerTurn(test *testing.T) {
	fmt.Println("\n Test for GET_PLAYER_Turn started!")
	diceCC := new(Game)
	stub := shim.NewMockStub("GetPlayerTurn", diceCC)
	if stub == nil {
		test.Fatalf("MockStub creation failed")
	}

	gameID := "game-11172019-01"
	firstPlayer := "player-test-01"
	secondPlayer := "player-test-02"

	// Setting up Game with correct args and values
	checkInvoke(test, stub, [][]byte{[]byte(SETUPGAME), []byte(gameID), []byte(firstPlayer), []byte(secondPlayer)})
	// Invoking getPlayerPoints function with correct arguments
	checkInvoke(test, stub, [][]byte{[]byte(GETPLAYERTURN), []byte(gameID)})
	// Passing incorrect number of function arguments to check failure
	checkBadInvoke(test, stub, [][]byte{[]byte(GETPLAYERTURN)})
}

func checkBadInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("transact-01", args)
	if res.Status == shim.OK {
		fmt.Println("Error!! Invoke for method -", string(args[0]), " unexpectedly succeeded")
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("transact-01", args)
	if res.Status != shim.OK {
		fmt.Println("Error!! Invoke for method -", string(args[0]), " failed ", string(res.Message))
		t.FailNow()
	}
	fmt.Println(string(res.Payload))
}

func checkQuery(t *testing.T, stub *shim.MockStub, args [][]byte, expectedVal int) {
	gameID := args[1]
	res := stub.MockInvoke("transact-01", args)
	if res.Status != shim.OK {
		fmt.Println("Error!! Query for key -", gameID, " failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Error!! Query for key -", gameID, " failed to get value")
		t.FailNow()
	}

	actualVal, _ := strconv.Atoi(string(res.Payload))
	if actualVal != expectedVal {
		fmt.Println("Error!! Queried value for key -", gameID, " was ", actualVal, "  not ", expectedVal, " as expected")
		t.FailNow()
	}
}
