package main

import (
	"fmt"
	"project/git_class/controllers"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello gais")

	controllers.PrintUserInput("Jerry")
	controllers.PrintUserInput("Hasan")
	controllers.PrintUserInput("Hans")
	controllers.PrintUserInput("Darien")
}
