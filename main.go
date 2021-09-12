package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)


func getUserAge() (int, error) {
	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')

	if runtime.GOOS == "windows" {
		age = strings.Replace(age, "\r\n", "", -1)
	} else {
		age = strings.Replace(age, "\n", "", -1)
	}
	parsedAge, err := strconv.ParseInt(age, 0, 64)

	return int(parsedAge), err
}

func main() {
	userAge, err := getUserAge()

	if err != nil {
		fmt.Println("Not correct input")
		return
	}

	if userAge > 18 && userAge < 65 {
		fmt.Println("Welcome to the club")
	} else if userAge < 18 {
		fmt.Println("Sorry, you are too young")
	} else {
		fmt.Println("Good over 65")
	}

	// switch parsedAge {
	// case "1":
	// 		fmt.Println("Option 1")
	// case "2":
	// 		fmt.Println("Option 2")
	// case "3":
	// 		fmt.Println("Option 3")
	// default:
	// 		fmt.Println("Invalid choice!")
	// }
}