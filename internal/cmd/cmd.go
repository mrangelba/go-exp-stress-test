package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/mrangelba/go-exp-stress-test/internal/report"
	"github.com/spf13/cobra"
)

var reportRequest *report.Report

var rootCmd = &cobra.Command{
	Use:   "go-exp-stress-test",
	Short: "Teste de stress em serviços HTTP.",
	Long:  "Teste de stress em serviços HTTP.",
	Run: func(cmd *cobra.Command, args []string) {
		reportRequest = report.NewReport()
		reportRequest.Start()
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt32("requests")
		concurrency, _ := cmd.Flags().GetInt32("concurrency")
		verbose, _ := cmd.Flags().GetBool("verbose")

		var wg sync.WaitGroup

		for i := int32(0); i < concurrency; i++ {
			wg.Add(1)

			go func(routine int32) {
				defer wg.Done()

				for j := int32(0); j < requests/concurrency; j++ {
					fetchURL(cmd.Context(), int32(routine), int32(j), url, verbose)
				}
			}(i)
		}

		wg.Wait()

		reportRequest.Print(verbose)
	},
}

func fetchURL(ctx context.Context, routine int32, request int32, url string, verbose bool) {
	if verbose {
		fmt.Printf("[%d] Request %d - URL %s\n", routine+1, request+1, url)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		reportRequest.AddItem("0")

		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		reportRequest.AddItem("0")

		if verbose {
			fmt.Printf("[%d] Request %d - Error: %s\n", routine+1, request+1, err.Error())
		}

		return
	}
	defer resp.Body.Close()

	reportRequest.AddItem(fmt.Sprintf("%d", resp.StatusCode))

	if verbose {
		fmt.Printf("[%d] Request %d - Status: %d\n", routine+1, request+1, resp.StatusCode)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL do serviço a ser testado (obrigatório)")
	rootCmd.MarkFlagRequired("url")
	rootCmd.Flags().Int32P("requests", "r", 10, "Número total de solicitações a serem enviadas (padrão: 10)")
	rootCmd.Flags().Int32P("concurrency", "c", 2, "Número de chamadas simultâneas (padrão: 2)")
	rootCmd.Flags().BoolP("verbose", "v", false, "Exibir detalhes das requisições (padrão: false)")
}
