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
	for i := 0; i < 100_000_000; i++ {
		tentativa := itoa8(i)
		if verificaSenha(tentativa, senhaReal) {
			fmt.Println("Senha encontrada:", tentativa)
			break
		}
	}
	duracao := time.Since(inicio)
	fmt.Printf("Tempo: %v\n", duracao)
}
