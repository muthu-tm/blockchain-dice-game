package main

import (
	"encoding/json"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
)

// PutState - update the state database
func PutState(stub shim.ChaincodeStubInterface, key string, value []byte) error {
	return stub.PutState(key, value)
}

// IsExistingGame - checks whether the Game already set or not
func IsExistingGame(stub shim.ChaincodeStubInterface, key string) (bool, error) {
	// If the key does not exist in the state database, (nil, nil) is returned.
	byteVal, err := stub.GetState(key)
	if err != nil {
		return false, err
	}
	if byteVal != nil {
		return true, nil
	}

	return false, nil
}

// GetGameByID - Get the Game and returns the Game struct after conversion
func GetGameByID(stub shim.ChaincodeStubInterface, gameID string) (Dice, error) {
	bytesData, err := getState(stub, gameID)
	var dice Dice
	if err != nil {
		return dice, err
	}

	err = json.Unmarshal(bytesData, &dice)
	if err != nil {
		return dice, err
	}

	return dice, nil
}

func getState(stub shim.ChaincodeStubInterface, key string) ([]byte, error) {
	resValue, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}

	return resValue, nil
}
