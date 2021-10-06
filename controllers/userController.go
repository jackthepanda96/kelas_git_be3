package controllers

import "fmt"

func TestPrintUserControl() {
	fmt.Println("Print from User Controller")
}

func PrintUserInput(input string) {
	fmt.Println(input)
}

func FiturDevelop1() {
	fmt.Println("Ini fitur baru 1")
	fmt.Println(InsertBuku("Atomic Habits", "Mark Manson"))
}

func InsertBuku(judul, pengarang string) string {
	return fmt.Sprint(judul, " Karangan : ", pengarang)
}
