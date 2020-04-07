package main

import (
	"markov/cmd"
)

func main() {
	cmd.Execute()
}

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"time"

// 	"markov/distribution"
// )

// func scanDirectory() []string {
// 	var files []string

// 	root := "../texts"
// 	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			println("Error scanning directory")
// 			println(err.Error())
// 			os.Exit(1)
// 		}
// 		if info.IsDir() == false {
// 			files = append(files, path)

// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return files
// }

// func main() {
// 	println("Scanning text directory")

// 	// Scan the files, get word count
// 	files := scanDirectory()
// 	println("Found the following files")
// 	for _, filepath := range files {
// 		file, err := os.Open(filepath)
// 		if err != nil {
// 			panic(nil)
// 		}
// 		scanner := bufio.NewScanner(file)
// 		words := distribution.TokenizeScanner(scanner)
// 		file.Close()

// 		fmt.Printf(" | %d\n", len(words))
// 		dist := distribution.FromArray(words)

// 		rand.Seed(time.Now().UnixNano())
// 		sentenceLength := rand.Intn(10) + 5
// 		println(sentenceLength)
// 		sentence := []string{""}
// 		for len(sentence) < sentenceLength {
// 			lastWord := sentence[len(sentence)-1]
// 			nextWord := dist.GetRandWord(lastWord)
// 			sentence = append(sentence, nextWord)
// 		}

// 		str := strings.Join(sentence, " ")
// 		str = strings.TrimSpace(str)
// 		str += "."
// 		println(str)
// 	}
// }
