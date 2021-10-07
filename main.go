package main

import (
	"fmt"
	"project/git_class/controllers"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello gais")

	controllers.PrintUserInput("Jerry")
	controllers.PrintUserInput("This one too")
	controllers.PrintUserInput("Test For Merge")
}
