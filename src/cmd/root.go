package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "markov",
	Short:   "markov helps generate random texts",
	Long:    "Helps to parse texts and generate markov distributions from the tokens.\nAdditional texts can then be generated from the markov distributions.",
	Version: "0.0.1",
}

// Execute starts the CLI
func Execute() {
	// Add Commands
	initDistCmd()
	rootCmd.AddCommand(distCmd)
	rootCmd.AddCommand(genSentenceCmd)

	// Dist Cmd

	// Start
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
