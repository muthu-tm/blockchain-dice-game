# Fabric-Dice-Game #
> Smartcontract for a simple Dice game

## Language ##
* Go - go1.11

---

## Run ##
To test or run this smart contract, we need to deploy this in an simple fabric network and call the invoke functions;

### Steps to play Dice! ###
* Once After successfull deployment, you can query or invoke the following function from Fabric SDK
* _setupGame_ - First need to setup the game with two player's addresses
...Function Name: setupGame, arguments: gameID, firstPlayerId, secondPlayerID
* _play_ - Second start playing the game by providing points scored and string value "hold" to change the player turn
...Fucntion Name: play, arguments: gameID, playerID, pointsEarned, hold
* _getPlayerPoints_ - Get a player's total points. Points used to summed once the current turn has been finished fully; untill then the currentTurnTotal is calculated separately as per the conditions followed 
...Fucntion Name: getPlayerPoints, arguments: gameID, playerId
* _getPlayerTurn_ - You can get the current player turn
...Fucntion Name: getPlayerTurn, arguments: gameID

#### Arguments Type ####
gameID: string - Each game will have a unique id
playerID: string - Each player has their own unique identity and ID
pointsEarned: uint - Points earned on this dice roll; if 1 no points and the plyer turn gets changed as per the rule
hold: string - "hold" or can be empty

## TODO ##
* Go Test for the chaincode:
...TDD using [go test](https://golang.org/pkg/testing/ "godoc - golang testing")
...BDD using [Go Convey](http://goconvey.co/, "GoConvey - Go behavioral Tests")
* Encryption of data
...Encrypt the value before storing into the ledger using [fabric bccsp](https://godoc.org/github.com/hyperledger/fabric/bccsp)