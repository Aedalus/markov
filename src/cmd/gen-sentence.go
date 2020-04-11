package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"markov/distribution"
	"markov/generation"
)

func openMarkovFiles(files []string) []distribution.TextDistribution {
	distributions := []distribution.TextDistribution{}

	for _, file := range files {
		abspath, err := filepath.Abs(file)
		if err != nil {
			println("Could not resolve file " + file)
			os.Exit(1)
		}

		jsonFile, err := os.Open(abspath)
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			println("Error reading file: " + abspath)
			os.Exit(1)
		}

		var dist distribution.TextDistribution
		err = json.Unmarshal(byteValue, &dist)
		if err != nil {
			panic(err)
		}

		distributions = append(distributions, dist)
	}
	return distributions
}

var genSentenceCmd = &cobra.Command{
	Use:   "gen-sentence [...files]",
	Short: "Generates a sentence from a markov file.",
	Long:  "Generates a sentence from a markov file.",
	Run: func(cmd *cobra.Command, args []string) {

		textDists := openMarkovFiles(args)

		combined := make(distribution.TextDistribution)

		for _, dist := range textDists {
			combined = combined.Combine(dist)
		}

		snt := generation.GenerateSentence(combined)
		println(snt)
	},
}
