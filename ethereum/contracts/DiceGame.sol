pragma solidity ^0.5.0;

contract DiceGame {

    struct Dice {
        uint    totalRounds;
        uint    playerTurn;
        address payable player1;
        address payable player2;
        mapping(address => uint) playerPoints;
    }
    
    uint[] private Player1EachTurnPoints;
    uint[] private Player2EachTurnPoints;
    uint private currentTurnTotal;

    Dice private game;

    /**
     * @dev setup a new Dice game for the given two players
     * @param _player1 - First player
     * @param _player2 - Second player
     */
    function setupGame(address payable _player1, address payable _player2) public {
        require(_player1 != address(0) && _player2 != address(0), "Provide both player address to setup the GAME!");
        require (_player1 != _player2, "Both players shouln't be same! please provide different player address");
        game = Dice(0, 1, _player1, _player2);
        currentTurnTotal = 0;

    }

    /**
     * @dev As per the player's turn, needs to send the point scored now and the string value of "hold" to change the player turn
     * @param _points - First player
     * @param _stringVal - String value of "hold" to change the player turn
     */
    function play(uint _points, string memory _stringVal) public returns (string memory) {
        require (game.player1 != address(0), "It seems the GAME has not started yet");
        address _player = msg.sender;
        string memory val_hold = "hold";
        if ((_player == game.player1 && game.playerTurn == 1) || (_player == game.player2 && game.playerTurn == 2)) {
            if (compareStrings(val_hold, _stringVal) || _points == 1) {
                if (_player == game.player1) {
                    Player1EachTurnPoints.push(currentTurnTotal);
                } else {
                    Player2EachTurnPoints.push(currentTurnTotal);
                }
                game.playerPoints[_player] += currentTurnTotal;
                currentTurnTotal = 0;
                game.totalRounds = game.totalRounds + 1;
                game.playerTurn = changePlayerTurn(_player);
                return "Next Player Turn";
            }
            currentTurnTotal += _points;
        } else {
            revert("This is not your Turn!!");
        }
    }

    function getPlayerTurn() public view returns (uint) {
        require (game.player1 != address(0), "It seems the GAME has not started yet");
        return game.playerTurn;
    }

    function getPlayerPoints(address _player) public view returns (uint) {
        require (game.player1 != address(0), "It seems the GAME has not started yet");
        return game.playerPoints[_player];
    }


    function changePlayerTurn(address _player) private view returns (uint) {
        if (_player == game.player1) {
            return 2;
        } else {
            return 1;
        }
    }

    function compareStrings (string memory str_val1, string memory str_val2) private pure returns (bool) {
            return (keccak256(abi.encodePacked((str_val1))) == keccak256(abi.encodePacked((str_val2))) );
       }

}