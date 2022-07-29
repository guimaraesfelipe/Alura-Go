package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	intro()

	for {
		showMenu()

		switch readOption() {
		case 1:
			startMonitor()
			fmt.Println("")
		case 2:
			printLogs()
			fmt.Println("")
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Opcão invalida")
			fmt.Println("")
			time.Sleep(2 * time.Second)
		}
	}

}

func intro() {
	name := "Felipe"
	version := 1.1
	fmt.Println("Olá,", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func readOption() int {
	var option int
	fmt.Scan(&option)

	return option
}

func startMonitor() {
	fmt.Println("Monitorando...")
	urls := readUrlsFile()

	for _, url := range urls {
		testResponse(url)

		time.Sleep(2 * time.Second)
	}

}

func testResponse(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	} else {
		if response.StatusCode == 200 {
			fmt.Println("Site:", url, "foi carregado com sucesso!")
			createLog(url, true)
		} else {
			fmt.Println("Site: ", url, "apresentou falha no acesso. StatusCode:", response.StatusCode)
			createLog(url, false)
		}
	}
}

func readUrlsFile() []string {
	var urls []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	} else {
		reader := bufio.NewReader(file)
		for {
			row, err := reader.ReadString('\n')
			row = strings.TrimSpace(row)

			urls = append(urls, row)

			if err == io.EOF {
				break
			}
		}

	}
	file.Close()
	return urls
}

func createLog(url string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + url + " - online: " + strconv.FormatBool(status) + "\n")
	}
	file.Close()
}

func printLogs() {
	fmt.Println("")

	file, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Logs:\n" + string(file))
	}
}
