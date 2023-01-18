package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please input your guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input, please try again.", err)
			continue
		}
		input = strings.Trim(input, "\r\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number, please try again.")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number, please try again.")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
