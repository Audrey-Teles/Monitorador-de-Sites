package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	for {
		showIntroduction()

		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
		}
	}

}

func showIntroduction() {
	name := "Audrey"
	version := 1.1

	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)

}

func readCommand() int {
	var commandRead int

	fmt.Scan(&commandRead)
	fmt.Println("O comando escolhido foi", commandRead)

	return commandRead
}

func showMenu() {

	fmt.Println("#---------- MENU ----------#")
	fmt.Println("|[1] Iniciar Monitoramento |")
	fmt.Println("|[2] Exibir Logs           |")
	fmt.Println("|[0] Sair do Programa      |")
	fmt.Println("#--------------------------#")

}

func startMonitoring() {
	fmt.Println("Monitorando...")

	site := "https://random-status-code.herokuapp.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully :)")
	} else {
		fmt.Println("Site:", site, "having problems :(")
		fmt.Println("Status code:", resp.StatusCode)
	}
}
