package report

import (
	"fmt"
	"sync"
	"time"
)

var reportMutex = &sync.Mutex{}

type Report struct {
	StartTime     time.Time
	TotalRequests int
	Items         map[string]map[string]int
}

func NewReport() *Report {
	return &Report{
		TotalRequests: 0,
		Items:         map[string]map[string]int{},
	}
}

func (r *Report) Start() {
	r.StartTime = time.Now()
}

func (r *Report) Print(verbose bool) {
	if verbose {
		fmt.Printf("\n\n")
	}

	fmt.Printf("----------------------------------------------------------------------\n")
	fmt.Printf("Relatório de execução\n")
	fmt.Printf("----------------------------------------------------------------------\n")
	fmt.Printf("Tempo total gasto na execução: %s\n", time.Since(r.StartTime))
	fmt.Printf("Quantidade total de requests realizados: %d\n", r.TotalRequests)
	fmt.Printf("----------------------------------------------------------------------\n")

	if _, ok := r.Items["0"]; ok {
		fmt.Printf("Quantidade de requests não concluídos: %d\n", r.Items["0"]["total"])
	}

	for key, value := range r.Items {
		if key != "0" {
			fmt.Printf("Quantidade de requests com status HTTP %s: %d\n", key, value["total"])
			fmt.Printf("----------------------------------------------------------------------\n")
		}
	}
}

func (r *Report) AddItem(statusCode string) {
	reportMutex.Lock()
	defer reportMutex.Unlock()

	r.TotalRequests++

	if _, ok := r.Items[statusCode]; !ok {
		r.Items[statusCode] = map[string]int{}
	}

	r.Items[statusCode]["total"]++
}
