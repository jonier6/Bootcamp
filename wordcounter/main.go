package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetInput(r io.Reader) []string {
	var lineas []string
	scanner := bufio.NewScanner(r)


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

	fmt.Println("Enter the text (type 'exit' to finish):")

	modoLinea := false
	if len(os.Args) > 1 && os.Args[1] == "-l" {
		modoLinea = true
	}

	contenido := GetInput(os.Stdin)

	resultado := ProcesarEntrada(contenido, modoLinea)
	unidad := "words"
	if modoLinea {
		unidad = "lines"
	}
	fmt.Printf("\nfinal result: %d %s.\n", resultado, unidad)
}
