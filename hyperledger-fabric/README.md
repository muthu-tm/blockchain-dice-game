# Fabric-Dice-Game #
> Smartcontract for a simple Dice game

## Languages Used ##

| S.No   | Language    | Version  |
|:------:|:-----------:|:--------:|
| 1      | Go          | go1.11   |

---

## Run ##
To test or run this smart contract, we need to deploy this in an simple fabric network and call the invoke functions;

### Steps to play Dice! ###
* Once After successfull deployment, you can query or invoke the following function from Fabric SDK
* __setupGame__ - First need to setup the game with two player's addresses
  * Function Name: setupGame, arguments: gameID, firstPlayerId, secondPlayerID
* __play__ - Second start playing the game by providing points scored and string value "hold" to change the player turn
  * Function Name: play, arguments: gameID, playerID, pointsEarned, hold
* __getPlayerPoints__ - Get a player's total points. Points used to summed once the current turn has been finished fully; untill then the currentTurnTotal is calculated separately as per the conditions followed 
  * Function Name: getPlayerPoints, arguments: gameID, playerId
* __getPlayerTurn__ - You can get the current player turn
  * Function Name: getPlayerTurn, arguments: gameID

#### Arguments Type ####
* _gameID_: string - Each game will have a unique id
* _playerID_: string - Each player has their own unique identity and ID
* _pointsEarned_: uint - Points earned on this dice roll; if 1 no points and the plyer turn gets changed as per the rule
* _hold_: string - "hold" or can be empty

## TODO ##
* Go Test for the chaincode:
  * TDD using [go test](https://golang.org/pkg/testing/ "godoc - golang testing")
  * BDD using [Go Convey](http://goconvey.co/, "GoConvey - Go behavioral Tests")
* Encryption of data
  * Encrypt the value before storing into the ledger using [fabric bccsp](https://godoc.org/github.com/hyperledger/fabric/bccsp)