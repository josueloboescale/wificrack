package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	interfaceName := "wlan0"

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o nome da rede Wi-Fi: ")
	networkName, _ := reader.ReadString('\n')
	networkName = strings.TrimSpace(networkName)

	fmt.Print("Digite o nome do arquivo de senhas: ")
	passwordFile, _ := reader.ReadString('\n')
	passwordFile = strings.TrimSpace(passwordFile)

	file, err := os.Open(passwordFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		password := strings.TrimSpace(scanner.Text())

		cmd := exec.Command("sudo", "iwconfig", interfaceName, "essid", networkName, "key", password)

		err := cmd.Run()
		if err == nil {
			fmt.Printf("Conexão Wi-Fi estabelecida com sucesso! Senha utilizada: %s\n", password)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Nenhuma senha do arquivo funcionou. Conexão Wi-Fi não estabelecida.")
}
