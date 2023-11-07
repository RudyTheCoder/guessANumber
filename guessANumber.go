// main.go file
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type IPlayer interface { // interface
	guess() int
}

type Game struct {
	player           IPlayer // player is of type Iplayer.
	guessesMadeSoFar int     // keeps a count of the number of guesses made so far
	randomNumber     int     // random number generated uses the golang library.
	recentGuess      int     // holds the most recent guess.
}

type Autoguess struct {
	p        *Game // A pointer to the Game
	min, max int   //Integers for the min and max possible values
}

type Human struct { //empty struct
}

// This plays the game by calling the player’s guess() to get the next guess and outputting
// the appropriate response
func (g *Game) play() {

	g.recentGuess = g.player.guess() // g.recentGuess is holding the most recent guess.

	// this next if statement checks to see if we have already used up our 3 guesses and whether
	// recent guess is not equal to the random number.
	if (g.guessesMadeSoFar == 2) && (g.recentGuess != g.randomNumber) {
		fmt.Println("You ran out of guesses. Game over")
		return

	} else if g.recentGuess < g.randomNumber { // is the recent guess less than the random number.
		fmt.Println("Too low") // prints to the screen that the guess is too low.
		g.guessesMadeSoFar++   // increases the number of guesses made.
		g.play()               // calls the play method again.

	} else if g.recentGuess > g.randomNumber { // is recent guess larger than the random number.
		fmt.Println("Too high")
		g.guessesMadeSoFar++
		g.play()
	} else if g.recentGuess == g.randomNumber { // is the recent guess equal to the random number.
		fmt.Println("You Win!")
		return

	}

	return
}

// asks the user for the next number to guess
func (h *Human) guess() (num int) {

	var b int

	fmt.Print("Enter your next guess: ") // ask the user then return this number

	fmt.Scan(&b) // Returns the input

	return b
}

// will return an appropriate guess based on choosing the middle value of the possible remaining values

func (a *Autoguess) guess() (num int) {
	var midpoint int

	fmt.Print("The computer has chosen ") // prints to the screen a message.

	midpoint = a.min + ((a.max - a.min) / 2) // Binary search for the midpoint between min and max

	// if midpoint is not equal to the recent guess
	if midpoint != a.p.recentGuess {
		a.p.recentGuess = midpoint // make recent guess equal to midpoint

		fmt.Println(midpoint) // print the midpoint to the screen.

		return midpoint // return midpoint

		// else if midpoint is equal to recent guess and midpoint is greater than random number

	} else if (midpoint == a.p.recentGuess) && (midpoint > a.p.randomNumber) {
		a.max = midpoint - 1                     // The random number is less than midpoint, so shift max.
		midpoint = a.min + ((a.max - a.min) / 2) // calculate the new midpoint.
		a.p.recentGuess = midpoint
		fmt.Println(midpoint)

		return midpoint

		// else if midpoint is equal to recent guess and midpoint is less than random number.

	} else if (midpoint == a.p.recentGuess) && (midpoint < a.p.randomNumber) {
		a.min = midpoint + 1 // The random number is greater than midpoint, so shift min.
		midpoint = a.min + ((a.max - a.min) / 2)
		a.p.recentGuess = midpoint
		fmt.Println(midpoint)

		return midpoint
	}

	return midpoint

}

func main() {

	var d string // this stores the users input.
	var c Game   //c is a variable of type Game

	fmt.Println("Guess a number to demo interfaces\n")
	fmt.Println("You have 3 guesses to guess a number from 1 to 10")
	fmt.Print("Do you want to make the guesses? (y/n -- if n guesses will be generated for you):")
	fmt.Scanln(&d) // Reads input from the user.

	if d == "y" {
		c.player = &Human{} // This is Inheritance in Golang.

	} else {

		c.player = &Autoguess{} // This is Inheritance in Golang.
		var a Autoguess
		a.p = &c
		a.min = 1
		a.max = 10
		c.player = &a

	}

	c.guessesMadeSoFar = 0 // initializing gussesMadeSoFar in the Game struct.
	min := 1
	max := 10
	rand.Seed(time.Now().UnixNano()) // uses the built in library to create a new random number
	//each time.
	c.randomNumber = (rand.Intn(max-min) + min) // initializing "randomNumber" in Game struct.
	c.play()                                    // plays the game.

}
