package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func itoa8(n int) string {
	var b [8]byte
	for i := 7; i >= 0; i-- {
		b[i] = byte('0' + n%10)
		n /= 10
	}
	return string(b[:])
}

func workerRange(start, end int, senhaReal string, found *atomic.Bool, result *atomic.Value, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i < end; i++ {
		if found.Load() {
			return
		}

		tentativa := itoa8(i)
		if verificaSenha(tentativa, senhaReal) {
			result.Store(tentativa)
			found.Store(true)
			return
		}
	}
}

func RunConcorrente(senhaReal string, workers int) {
	if workers <= 0 {
		workers = runtime.NumCPU()
	}

	const max = 100_000_000 // 00000000..99999999

	inicio := time.Now()

	// Divide o espaço de busca em fatias aproximadamente iguais
	chunk := (max + workers - 1) / workers

	var wg sync.WaitGroup
	var found atomic.Bool
	var result atomic.Value

	wg.Add(workers)
	for w := 0; w < workers; w++ {
		start := w * chunk
		end := start + chunk
		if end > max {
			end = max
		}
		go workerRange(start, end, senhaReal, &found, &result, &wg)
	}

	wg.Wait()

	duracao := time.Since(inicio)
	if found.Load() {
		fmt.Printf("Senha encontrada: %s (workers: %d)\n", result.Load().(string), workers)
	} else {
		fmt.Println("Senha não encontrada (isso não deveria acontecer).")
	}
	fmt.Printf("Tempo: %v\n", duracao)
}
