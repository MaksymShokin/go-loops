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


func main() {
	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')

	if runtime.GOOS == "windows" {
		age = strings.Replace(age, "\r\n", "", -1)
	} else {
		age = strings.Replace(age, "\n", "", -1)
	}
	parsedAge, _ := strconv.ParseInt(age, 0, 64)



	if parsedAge > 18 {
		fmt.Println("Welcome to the club")
	} else if parsedAge < 18 {
		fmt.Println("Sorry, you are too young")
	}
}