package main

import "fmt"

func main() {
	var age string
	fmt.Println("How old are you?")
	fmt.Scan(&age)
	fmt.Println("Oh, you are already " + age + " :O")

}
