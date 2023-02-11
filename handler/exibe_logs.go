package handler

import (
	"fmt"
	"io/ioutil"
)

func ExibeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Erro na abertura do arquivo:", err)
		fmt.Printf("Retornando para o menu...\n\n")
	} else {
		fmt.Println("Exibindo Logs...")
		fmt.Println(string(arquivo))
	}
}
