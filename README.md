# Guess a Number Game with Go

This Go program implements a simple "Guess a Number" game that allows the user to guess a random number within a specified range. The user can choose to make guesses themselves or let the computer generate guesses automatically. The program utilizes interfaces, structs, and methods to achieve its functionality.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Running the Program](#running-the-program)
- [How to Play](#how-to-play)
- [Program Structure](#program-structure)
- [Contributing](#contributing)


## Getting Started

### Prerequisites

Before running the program, ensure that you have Go installed on your computer. You can download and install Go from the [official Go website](https://golang.org/dl/).

### Running the Program

1. Clone or download this repository to your local machine.

2. Open a terminal or command prompt and navigate to the directory where the `guess_num.go` file is located.

3. Run the program using the following command:

   ```shell
   go run guessANumber.go


1. Follow the on-screen instructions to play the game.

## How to Play

The game presents you with the following options:

* You have 3 guesses to guess a number from 1 to 10.
* Do you want to make the guesses? (y/n -- if 'n,' guesses will be generated for you):

You can choose to make guesses yourself by entering 'y' and providing your guesses, or you can let the computer generate guesses for you by entering 'n'.

If you choose to make guesses yourself ('y'), you will be prompted to enter your guesses one by one. The program will provide feedback on each guess, indicating whether it's too low, too high, or correct.

If you choose to let the computer generate guesses ('n'), the computer will automatically make guesses and provide feedback based on its guesses.

The game ends when you correctly guess the number, run out of guesses, or decide to exit.

## Program Structure

* main.go: The main program file containing the game logic, including struct definitions, interface, methods, and 
  the game loop.
* IPlayer interface: Defines the guess method that all players must implement.
* Game struct: Represents the game state and player information.
* Autoguess struct: Represents the computer player for automated guesses.
* Human struct: Represents the human player.


## Contributing

Contributions to this project are welcome! If you have any improvements, bug fixes, or feature additions, feel free to open an issue or submit a pull request.