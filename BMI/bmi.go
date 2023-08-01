package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	//output information
	fmt.Println("BMI Calculator")
	fmt.Println("---------------------------------")
	//prompt for user input
	fmt.Println("Enter your weght in (kg): ")

	weightInput, _ := reader.ReadString('\n')

	fmt.Print("Enter your height in (m): ")

	heightInput, _ := reader.ReadString('\n')

	fmt.Print(weightInput, heightInput)
	//saving user inputs to variables
	//calculate BMI (weight/(height*height))
	//output the calculator BMI
}
