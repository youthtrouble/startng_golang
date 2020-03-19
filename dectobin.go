package main

import (
	"fmt"
	"strconv"
)

func main() {
	var value int64
	fmt.Println("Enter a value")
	fmt.Scan(&value)
	output := strconv.FormatInt(value, 2)
	fmt.Print("The Binary Equivalent of your input is ", output)
}
