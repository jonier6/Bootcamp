package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"flag"
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

func ProcesarEntrada(datos []string, contarLineas, contarBytes bool ) int {
	contador := 0

	for _, linea := range datos {
		if contarLineas {
			contador++
		}else if contarBytes{
			contador += len([]byte(linea)) + 1

		} else {
			palabras := strings.Fields(linea)
			contador += len(palabras)
		}
	}

	return contador
}

func main() {

	fmt.Println("Enter the text (type 'exit' to finish):")

	countLines := flag.Bool("l", false, "count lines")
	countBytes := flag.Bool("b", false, "count bytes")
	
	flag.Parse()

	contenido := GetInput(os.Stdin)

	resultado := ProcesarEntrada(contenido, *countLines, *countBytes)

	unidad := "words"

	if *countLines {
		unidad = "lines"
	}else if *countBytes{
		unidad = "bytes"
	}
	fmt.Printf("\nfinal result: %d %s.\n", resultado, unidad)
}
