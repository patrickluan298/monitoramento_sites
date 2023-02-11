package handler

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func LeSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	defer arquivo.Close()

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	return sites
}
