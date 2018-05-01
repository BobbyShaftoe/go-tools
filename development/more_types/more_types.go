package main

import (
	"fmt"
)

type data struct {
	a int
	b string
}

type Stats struct {
	height   int64
	weight   int
	starSign string
}

var myName string = "nick"
var myAge int = 40
var isAwesome bool = true
var friendsAges = [4]int{3, 5, 7, 11}

type Computers struct {
	Mac   bool
	Linux bool
	Windows bool
}

var computersOwned = Computers{true, true, false}

var languagesKnown = map[string]string{
	"name": "nick", "knows": "Golang",
}

var vitalStats = map[string]Stats{
	"nick": {195, 85, "Aries"},
}

var nickLikes = map[string]interface{}{
	"languages": struct {
		first  string
		second string
	}{"Golang", "Python"},
}

var nickLoves = map[string][]string{
	"drinks": {"coffee", "coke light"},
	"food":   {"ramen", "ice cream"},
}

var skillLevel = data{
	a: 10, b: "out of ten",
}

func main() {

	fmt.Printf("\nStarting...\n")

	const VarLength int = len(friendsAges)

	var I interface{}

	typeFunc := func(i interface{}) {
		switch varType := I.(type) {

		case [VarLength]int:
			result := I.([VarLength]int)
			fmt.Printf("\n\tType of variable is: %T\n", varType)
			fmt.Printf("\tDynamic values are:\n")
			for i := range result {
				fmt.Printf("\t\t[%v] %v\n", i, result[i])
			}

		case string:
			result := I.(string)
			fmt.Printf("\n\tType of variable is: %T\n", varType)
			fmt.Printf("\tDynamic value is: %v\n", result)

		case int:
			result := I.(int)
			fmt.Printf("\n\tType of variable is: %T\n", varType)
			fmt.Printf("\tDynamic value is: %v\n", result)

		case bool:
			result := I.(bool)
			fmt.Printf("\n\tType of variable is: %T\n", varType)
			fmt.Printf("\tDynamic value is: %v\n", result)

		case map[string]string:
			result := I.(map[string]string)
			fmt.Printf("\n\tType of variable is: %T\n", varType)
			fmt.Printf("\tDynamic values are:\n")
			for k, v := range result {
				fmt.Printf("\t\t%s: %s\n", k, v)
			}

		case map[string]int:
			result := I.(map[string]int)
			fmt.Printf("\n\tType of variable is: %T\n", varType)
			fmt.Printf("\tDynamic value is: %v\n", result)
			for k, i := range result {
				fmt.Printf("\t\t%s: %v\n", k, i)
			}

		case map[string][]string:
			result := I.(map[string][]string)
			fmt.Printf("\n\tType of variable is: %T\n", varType)
			fmt.Printf("\tDynamic values are:\n")
			for k, v := range result {
				fmt.Printf("\t\t%s:\n", k)
				for s := range v {
					fmt.Printf("\t\t  - %s\n", v[s])
				}
			}

		case map[string]struct{}:
			result := I.(map[string]struct{})
			fmt.Printf("\n\tType of variable is: %T\n", varType)
			fmt.Printf("\tDynamic value is: %v\n", result)

		case map[string]interface{}:
			result := I.(map[string]interface{})
			fmt.Printf("\n\tType of variable is: %T\n", varType)
			fmt.Printf("\tDynamic values are:\n")
			for k, v := range result {
				fmt.Printf("\t\t%s: %v\n", k, v)
			}

		case map[string]Stats:
			result := I.(map[string]Stats)
			fmt.Printf("\n\tType of variable is: %T\n", varType)
			fmt.Printf("\tDynamic values are:\n")
			for k, v := range result {
				fmt.Printf("\t\t%s:\n", k)
				fmt.Printf("\t\t  - Height: %d\n", v.height)
				fmt.Printf("\t\t  - Weight: %d\n", v.weight)
				fmt.Printf("\t\t  - Star Sign: %s\n", v.starSign)
			}

		case Computers:
			result := I.(Computers)
			fmt.Printf("\n\tType of variable is: %T\n", varType)
			fmt.Printf("\tValues are: \n")
			fmt.Printf("\t\tMac: %v \n\t\tLinux: %v\n\t\tWindows: %v\n\n", result.Mac, result.Linux, result.Windows)



		}
	}

	I = myName
	typeFunc(I)

	I = myAge
	typeFunc(I)

	I = friendsAges
	typeFunc(I)

	I = languagesKnown
	typeFunc(I)

	I = skillLevel
	typeFunc(I)

	I = vitalStats
	typeFunc(I)

	I = nickLikes
	typeFunc(I)

	I = nickLoves
	typeFunc(I)

	I = computersOwned
	typeFunc(I)

	I = isAwesome
	typeFunc(I)

}
