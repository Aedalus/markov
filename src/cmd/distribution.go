package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"markov/distribution"
	"markov/tokenizer"
)

var outFile string

func initDistCmd() {
	distCmd.Flags().StringVarP(&outFile, "out", "o", "", "Print to a file instead")
}

var distCmd = &cobra.Command{
	Use:   "dist [files]",
	Short: "Tokenizes a file, calculates the distribution",
	Long:  "Tokenizes a file and calculates the markob distribution. The resulting structure is printed to stdout",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		combined := make(distribution.TextDistribution)
		for _, path := range args {
			abspath, err := filepath.Abs(path)
			if err != nil {
				println("Error: Could not resolve filepath")
				os.Exit(1)
			}

			file, err := os.Open(abspath)
			if err != nil {
				println("Error: Could not open file " + path)
				os.Exit(1)
			}

			scanner := bufio.NewScanner(file)
			words := tokenizer.FromScanner(scanner)
			dist := distribution.FromArray(words)

			combined = combined.Combine(dist)
		}

		b, err := json.Marshal(combined)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if outFile == "" {
			fmt.Println(string(b))
		} else {
			outAbs, err := filepath.Abs(outFile)
			if err != nil {
				println("Error: Could not write to " + outAbs)
				os.Exit(1)
			}

			err = ioutil.WriteFile(outAbs, b, 0644)
			if err != nil {
				println("Error: Could not write to " + outAbs)
				os.Exit(1)
			}
		}
	},
}
