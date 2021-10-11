package main

//the word 'main is used to make an executable type package. If you
//do not name it main.go, it will not build an executable.
//Convention dictates we call the main package 'main'.

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	//Make sure GOPATH is set to the go folder in dev on desktop!
	//remember to name folder according to whats in .mod file or face errors!
	"github.com/PeterEnglish/go-hello-world/doctor"
)

//Give package access to code contained in fmt package.

func main() {

	reader := bufio.NewReader(os.Stdin)
	sayHelloWorld("Hello World!")

	var whatToSay string
	whatToSay = doctor.Intro()
	fmt.Println(whatToSay)

	for {
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			response := doctor.Response(userInput)
			fmt.Println(response)
		}
	}

}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}

//GO COMMANDS
//'go run main.go' to run ->Compiles and Executes.
//'go build main.go' - just compiles a program
//'go fmt main.go' - formats code in each file in current directory
//'go get' - downloads the raw source code of someone elses package.
//'go test' = runs any tests associated with current project.
//A package is a collection of common source code files.
//Functions, types, variables, and constants defined in one source
//file are visible to all other source files within the same package
