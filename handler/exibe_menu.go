package handler

import "fmt"

func ExibeMenu() {
	fmt.Println("####### Monitoramento de Sites #######")
	fmt.Println("[1] - Iniciar Monitoramento")
	fmt.Println("[2] - Exibir Logs")
	fmt.Println("[0] - Sair do Programa")
	fmt.Printf("\nOpção escolhida: ")
}

func LeComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("")

	return comandoLido
}
