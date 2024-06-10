/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"time"

	"github.com/Survialander/stress-test/internal"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "Stress test.",

	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		startTime := time.Now()

		internal.ExecuteStressTest(url, requests, concurrency, startTime)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("url", "http://google.com.br", "URL to make the requests.")
	rootCmd.PersistentFlags().Int("requests", 100, "Number of requests.")
	rootCmd.PersistentFlags().Int("concurrency", 10, "Number of workers to be used.")
}
