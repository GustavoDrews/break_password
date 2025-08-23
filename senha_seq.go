package main

import (
	"fmt"
	"time"
)

func verificaSenha(tentativa, senhaReal string) bool {
	return tentativa == senhaReal
}

func RunSequencial(senhaReal string) {
	inicio := time.Now()

	for i := 0; i < 100000000; i++ {
		tentativa := fmt.Sprintf("%08d", i)

		if verificaSenha(tentativa, senhaReal) {
			fmt.Println("Senha encontrada:", tentativa)
			break
		}
	}
	duracao := time.Since(inicio)
	fmt.Printf("Tempo: %v\n", duracao)
}
