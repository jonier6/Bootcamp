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
	
	fmt.Print("Ingresa el valor mínimo: ")
	minStr := getInput()
	min, _ := strconv.ParseFloat(minStr, 64)

	
	fmt.Print("Ingresa el valor máximo: ")
	maxStr := getInput()
	max, _ := strconv.ParseFloat(maxStr, 64)

	
	fmt.Print("Ingresa los valores separados por espacio: ")
	listStr := getInput()
	
	
	strFields := strings.Fields(listStr)
	var floatValues []float64
	for _, s := range strFields {
		f, _ := strconv.ParseFloat(s, 64)
		floatValues = append(floatValues, f)
	}

	
	filtered := minmax(min, max, floatValues...)

	
	fmt.Printf("Resultado: %v\n", filtered)
}
