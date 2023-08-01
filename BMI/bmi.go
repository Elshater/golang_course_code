package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)
	
	fmt.Print(weight)

	fmt.Print(height)

	//saving user inputs to variables
	//calculate BMI (weight/(height*height))
	//output the calculator BMI
}
