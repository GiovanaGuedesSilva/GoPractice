package main

/*
	Demonstração de leitura de arquivos em Go.

	- Uso da biblioteca padrão "os" para abrir arquivos.
	- Uso de "io/ioutil" ou "os.ReadFile" para ler o conteúdo completo.
	- Uso de "bufio.Scanner" para ler linha por linha.
	- Tratamento de erros com if err != nil.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Caminho para o arquivo que será lido
	filePath := "example.txt"

	// Tentativa de abrir o arquivo
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close() // Garante que o arquivo será fechado ao final

	fmt.Println("Lendo o conteúdo do arquivo linha por linha:")

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Linha %d: %s\n", lineNumber, line)
		lineNumber++
	}

	// Verificando erro durante o scan
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}
}
