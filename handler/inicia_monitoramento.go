package handler

import (
	"fmt"
	"time"
)

const monitoramentos = 2
const delay = 5

func IniciaMonitoramento() {
	fmt.Println("Monitorando...")
	sites := LeSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			TestaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}
