package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, senhaReal string, jobs <-chan int, resultCh chan<- string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case i, ok := <-jobs:
			if !ok {
				return
			}
			tentativa := fmt.Sprintf("%08d", i)
			if verificaSenha(tentativa, senhaReal) {
				select {
				case resultCh <- tentativa:
				case <-ctx.Done():
				}
				return
			}
		}
	}
}

func RunConcorrente(senhaReal string, workers int) {
	if workers <= 0 {
		workers = runtime.NumCPU()
	}

	inicio := time.Now()
	const max = 100_000_000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int, 4096)
	resultCh := make(chan string, 1)

	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go worker(ctx, &wg, senhaReal, jobs, resultCh)
	}

	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			default:
			}
			jobs <- i
		}
		close(jobs)
	}()

	var encontrada string
	select {
	case encontrada = <-resultCh:
		cancel()
	case <-ctx.Done():
	}

	wg.Wait()

	duracao := time.Since(inicio)
	if encontrada != "" {
		fmt.Printf("Senha encontrada: %s (workers: %d)\n", encontrada, workers)
	} else {
		fmt.Println("Senha não encontrada (isso não deveria acontecer).")
	}
	fmt.Printf("Tempo: %v\n", duracao)
}
