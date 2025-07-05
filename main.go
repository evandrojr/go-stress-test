package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
	Duration   time.Duration
	Error      error
}

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 1, "Número total de requests")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")

	flag.Parse()

	if *url == "" {
		fmt.Println("A URL é obrigatória. Use --url=<sua_url>")
		return
	}

	fmt.Printf("Iniciando teste de carga para %s com %d requests e %d chamadas simultâneas...\n", *url, *requests, *concurrency)

	results := make(chan Result, *requests)
	var wg sync.WaitGroup
	client := &http.Client{}

	startTime := time.Now()

	// Cria um pool de workers
	requestQueue := make(chan struct{}, *concurrency)

	for i := 0; i < *requests; i++ {
		wg.Add(1)
		requestQueue <- struct{}{} // Adquire um slot no pool de workers
		go func() {
			defer wg.Done()
			defer func() { <-requestQueue }() // Libera o slot

			startReq := time.Now()
			resp, err := client.Get(*url)
			duration := time.Since(startReq)

			if err != nil {
				results <- Result{Error: err, Duration: duration}
				return
			}
			defer resp.Body.Close()
			results <- Result{StatusCode: resp.StatusCode, Duration: duration}
		}()
	}

	wg.Wait()
	close(results)

	totalTime := time.Since(startTime)

	// Processa os resultados
	statusCodeCounts := make(map[int]int)
	totalRequests := 0
	successRequests := 0

	for res := range results {
		totalRequests++
		if res.Error != nil {
			// Lida com erros, talvez contá-los separadamente ou registrá-los
			continue
		}
		statusCodeCounts[res.StatusCode]++
		if res.StatusCode == http.StatusOK {
			successRequests++
		}
	}

	fmt.Println("\n--- Relatório do Teste de Carga ---")
	fmt.Printf("Tempo total gasto: %s\n", totalTime)
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200 (OK): %d\n", statusCodeCounts[http.StatusOK])
	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for code, count := range statusCodeCounts {
		if code != http.StatusOK {
			fmt.Printf("  Status %d: %d\n", code, count)
		}
	}
}