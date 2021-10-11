package main

//the word 'main is used to make an executable type package. If you
//do not name it main.go, it will not build an executable.
//Convention dictates we call the main package 'main'.

import "fmt"

//Give package access to code contained in fmt package.

func main() {
	sayHelloWorld("Hello World!")
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}

//'go run main.go' to run ->Compiles and Executes.
//'go build main.go' - just compiles a program
//'go fmt main.go' - formats code in each file in current directory
//'go get' - downloads the raw source code of someone elses package.
//'go test' = runs any tests associated with current project.
//A package is a collection of common source code files.
//Functions, types, variables, and constants defined in one source
//file are visible to all other source files within the same package
