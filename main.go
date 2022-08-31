package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {

	for {
		showIntroduction()
		readSitesFile()
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("=========================== LEAVING THE PROGRAM... ===========================")
			os.Exit(0)
		default:
			fmt.Println("========================= I DON'T KNOW THIS COMMAND... =========================")

		}
	}

}

func showIntroduction() {
	name := "Audrey"
	version := 1.1

	fmt.Println("Olá,", name)
	fmt.Println("Este programa está na versão", version)

}

func readCommand() int {
	var commandRead int

	fmt.Scan(&commandRead)

	return commandRead
}

func showMenu() {

	fmt.Println("|========== MENU ==========|")
	fmt.Println("|[1] Start Monitoring      |")
	fmt.Println("|[2] Show Logs             |")
	fmt.Println("|[0] Exit                  |")
	fmt.Println("|==========================|")

}

func startMonitoring() {
	fmt.Println("================================ MONITORING... =================================")

	sites := readSitesFile()

	for i := 0; i < monitoring; i++ {
		fmt.Println("---------------------------------- TEST", i+1, "--------------------------------------")
		for i, site := range sites {
			fmt.Println("============================== SITE", i+1, "==========================================")
			testSite(site)
		}
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
	fmt.Println("================================================================================")

}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully :)")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "having problems. Status code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readSitesFile() []string {

	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " -> " + site + "- online:" + strconv.FormatBool(status) + "\n")

}

func showLogs() {
	fmt.Println("============================== SHOWING LOGS... ================================")

	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(file))
}
