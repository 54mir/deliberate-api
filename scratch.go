package main

import (
	"fmt"
	"strings"
)

//structs

func main() {
	listOfStrings := []string{"a", "d", "f"}
	stringOfChars := strings.Join(listOfStrings, ", ")
	fmt.Println(stringOfChars)
	fmt.Printf("%T", stringOfChars)

}
