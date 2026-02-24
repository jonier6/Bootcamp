package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func GetInput() []string {
	var lineas []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the text (type 'exit' to finish):")

	for scanner.Scan() {
		texto := scanner.Text()
		
		
		if strings.EqualFold(strings.TrimSpace(texto), "exit") {
			break
		}
		lineas = append(lineas, texto)
	}
	return lineas
}


func ProcesarEntrada(datos []string, contarLineas bool) int {
	contador := 0

	for _, linea := range datos {
		if contarLineas {
			
			contador++
		} else {
			
			palabras := strings.Fields(linea)
			contador += len(palabras)
		}
	}

	return contador
}

func main() {
	
	modoLinea := false
	if len(os.Args) > 1 && os.Args[1] == "-l" {
		modoLinea = true
	}


	contenido := GetInput()


	resultado := ProcesarEntrada(contenido, modoLinea)

	
	unidad := "words"
	if modoLinea {
		unidad = "lines"
	}
	fmt.Printf("\nfinal result: %d %s.\n", resultado, unidad)
}