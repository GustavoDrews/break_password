package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func lerSenha() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Digite sua senha (8 números): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) != 8 {
			fmt.Println("A senha deve ter exatamente 8 caracteres.")
			continue
		}
		soDigitos := true
		for _, r := range input {
			if !unicode.IsDigit(r) {
				soDigitos = false
				break
			}
		}
		if !soDigitos {
			fmt.Println("A senha deve conter apenas dígitos (0-9).")
			continue
		}
		return input
	}
}

func main() {
	mode := flag.String("mode", "seq", "modo de execução: seq ou conc")
	workers := flag.Int("workers", 0, "quantidade de workers (apenas para -mode=conc). 0 usa runtime.NumCPU()")
	flag.Parse()

	senhaReal := lerSenha()

	switch *mode {
	case "seq":
		RunSequencial(senhaReal)
	case "conc":
		RunConcorrente(senhaReal, *workers)
	default:
		fmt.Println("Modo inválido. Use -mode=seq ou -mode=conc")
	}
}
