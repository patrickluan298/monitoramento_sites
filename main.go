package main

import (
	"fmt"
	"os"

	"github.com/monitoramento_sites/handler"
)

func main() {
	for {
		handler.ExibeMenu()
		comando := handler.LeComando()

		switch comando {
		case 1:
			handler.IniciaMonitoramento()
		case 2:
			handler.ExibeLogs()
		case 0:
			fmt.Println("Encerrando o programa...")
			os.Exit(0)
		default:
			fmt.Printf("Opção Inválida!\n\n")
			main()
		}
	}
}
