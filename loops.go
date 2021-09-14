package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func parseInput(input string) string {
	if runtime.GOOS == "windows" {
		input = strings.Replace(input, "\r\n", "", -1)
	} else {
		input = strings.Replace(input, "\n", "", -1)
	}

	return input
}

func getInputNumber() (int, error) {
	fmt.Println("Please enter a number: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	input = parseInput(input)
	chosenInput, err := strconv.ParseInt(input, 0, 64)

	if err != nil {
		return 0, err
	}

	return int(chosenInput), nil
}

func main() {
	selectedChoice, err := getUserChoice()

	if err != nil {
		fmt.Println("Invalid choice, exiting!")
	}

	switch selectedChoice {
	case "1":
		calculateSumUpToNumber()

	case "2":
		calculateFactorial()

	case "3":
		sumNumbers()

	case "4":
		sumList()

	}
}

func calculateSumUpToNumber() {
	chosenNumber, err := getInputNumber()

	sum := 0

	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	for i := 1; i <= chosenNumber; i++ {
		sum += i
	}

	fmt.Printf("Sum is %v", sum)
}

func calculateFactorial() {
	chosenNumber, err := getInputNumber()

	factorial := 1

	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	for i := 1; i <= chosenNumber; i++ {
		factorial *= i
	}

	fmt.Printf("factorial is %v", factorial)
}

func sumNumbers() {}
func sumList()    {}

func getUserChoice() (string, error) {
	fmt.Println("Please enter your choice:")
	fmt.Println("1) Add all numbers up to number X")
	fmt.Println("2) Find the factorial of number X")
	fmt.Println("3) Sum up the selected numbers")
	fmt.Println("4) Sum up a list of numbers")

	userChoice, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userChoice = parseInput(userChoice)

	if userChoice == "1" || userChoice == "2" || userChoice == "3" || userChoice == "4" {
		return userChoice, nil
	}

	return "", errors.New("user choice is incorrect")

}
