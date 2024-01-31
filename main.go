package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func insertValueIntoArray(arr *[]int, value int) {
	lo := int(0)
	hi := len(*arr)
	mid := int(0)

	for lo < hi {
		mid = (lo + hi) / 2

		if (*arr)[mid] < value {
			lo = mid + 1
		} else if (*arr)[mid] > value {
			hi = mid
		} else {
			return
		}
	}
	if lo == len(*arr) {
		*arr = append(*arr, value)
	} else {
		*arr = append((*arr)[:lo+1], (*arr)[lo:]...)
		(*arr)[lo] = value
	}

}

func main() {
	fmt.Println(os.Args)

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <dictionary-file-path>")
	}

	trie := NewTrie()

	startTime := time.Now()
	fmt.Println("Loading dictionary...")
	loadDictionary(&trie, os.Args[1])
	elapsedTime := time.Since(startTime)

	fmt.Printf("Dictionary loaded in %d ms!\n", elapsedTime.Milliseconds())

	for {
		var input string
		fmt.Print("Find words that start with: ")
		fmt.Scan(&input)

		wordsWithPrefix := trie.GetWords(input)

		for _, word := range wordsWithPrefix {
			fmt.Println(word)
		}
	}
}

func loadDictionary(t *Trie, path string) {
	dictionaryFile, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open file %s\n", path)
		return
	}
	defer dictionaryFile.Close()

	reader := bufio.NewReader(dictionaryFile)
	for {
		word, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		word = word[:len(word)-1]
		t.Add(word)
	}
}
