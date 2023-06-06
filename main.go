package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"math/rand"
)

const constString = "and don't type your nubmer in, just  press enter when ready."

func main() {

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// generate a random number between 2 to 10 so that we simplify the solution
	var firstNum = rand.Intn(8) + 2
	var secondNum = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer int

	// create our reader variable, which  reads input from standard in (the keyboard)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Now we are going to play Number mind game, Let's Start")
	fmt.Println("+++++++++++++++**********++++++++++++++++")
	fmt.Println("")

	fmt.Println("Think a number between 1 to 10", constString)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNum, constString)
	reader.ReadString('\n')
	
	fmt.Println("Now multiply the result by", secondNum, constString)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", constString)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, constString)
	reader.ReadString('\n')

	// now the time to answer
	answer = firstNum*secondNum - subtraction
	fmt.Println("The answer is", answer)

}

