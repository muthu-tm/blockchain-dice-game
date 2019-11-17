# Ethereum-Dice-Game #
> Smartcontract for a simple Dice game

## Language ##
* Solidity - ^0.5.0

---

## Run ##
To test or run this smart contract use the [REMIX](https://remix-ide.readthedocs.io/en/latest/ "Remix latest documentation - Ethereum IDE") online by following below steps

---

### Steps to  setup the smartcontract ###
* Open [Remix Website](https://remix.ethereum.org/, "Remix - Ethereun IDE")
* Create new Solididty file namely __DiceGame.sol__
* Copy and paste this solidity code into that file
* Compile and check for any errors (Hopfully, you won't get any)
* Deploy the DiceGame contract with Environment setting as JavaXcript VM

### Steps to play Dice! ###
* Once After successfull deployment, you will see a new Contract under Deployed Contract side list
* Expand it and you will see the following public functions to call; **_play_, _setupGame_, _getPlayerPoints_, _getPlayerTurn_**
* _setupGame_ - First need to setup the game with two player's addresses
* _play_ - Second start playing the game by providing points scored and string value "hold" to change the player turn
* _getPlayerPoints_ - Get a player's total points. Points used to summed once the current turn has been finished fully; untill then the currentTurnTotal is calculated separately as per the conditions followed 
* _getPlayerTurn_ - You can get the current player turn