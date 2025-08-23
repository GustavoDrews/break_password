package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Digite sua senha (8 números): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) != 8 {
			fmt.Println("A senha deve ter exatamente 8 caracteres.")
			continue
		}
	}
}
