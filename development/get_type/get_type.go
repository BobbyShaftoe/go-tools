package main

import (
	"fmt"
	_ "fmt"
)

type data struct {
	a int
	b string
}

var myName string = "nick"
var myAge int = 40

var friendsAges = [4]int{3, 5, 7, 11}

var languagesKnown = map[string]string{
	"name": "nick", "knows": "Golang",
}

var mapStringString = "map[string]string"

var skillLevel = data{
	a: 10, b: "out of ten",
}

func main() {
	fmt.Printf("\nStarting...\n")
	getType(myName, "myName")
	getType(myAge, "myAge")
	getType(friendsAges, "friendsAges")
	getType(languagesKnown, "languagesKnown")
	getType(skillLevel, "skillLevel")

}

func getType(v interface{}, n string) {
		result, ok := (v).(int)
		if ok {
				fmt.Printf("\tValue type is integer, and dynamic value is: %v\n", result)
		}

		fmt.Printf("\tType of variable \"%s\" is: \"%T\"\n", n, v)

	fmt.Printf("\t\tvalue: %v\n\n", v)
	fmt.Printf("\t\tlength: %v\n\n", v)



		}


