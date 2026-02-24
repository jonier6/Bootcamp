package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}


func minmax(min, max float64, values ...float64) []float64 {
	var result []float64

	for _, v := range values {
		if v >= min && v <= max {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	
	fmt.Print("Enter the minimum value: ")
	minStr := getInput()
	min, err := strconv.ParseFloat(minStr, 64)
	if err != nil {
    	fmt.Println("Invalid minimum value")
    	return
	}

	
	fmt.Print("Enter the maximum value: ")
	maxStr := getInput()
	max, err := strconv.ParseFloat(maxStr, 64)
	if err != nil {
    	fmt.Println("Invalid maximum value")
    	return
	}

	if max <= min {
        fmt.Printf("Error, the maximun value must be grater than the minimum value entered")
        return 
    }
	
	fmt.Print("Enter the values separated by spaces: ")
	listStr := getInput()
	
	strFields := strings.Fields(listStr)
	var floatValues []float64
	for _, s := range strFields {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
    	fmt.Println("invalid ungresed values")
    	continue
	}
		floatValues = append(floatValues, f)
	}

	
	filtered := minmax(min, max, floatValues...)

	
	fmt.Printf("Resultado: %v\n", filtered)
}
