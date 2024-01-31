package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrWordNotFound = errors.New("word not found in trie")

type TrieNode struct {
	Letter   byte
	children []*TrieNode
}

type Trie struct {
	root TrieNode
}

func NewTrie() Trie {
	return Trie{
		TrieNode{
			children: make([]*TrieNode, 0),
		},
	}
}

func (n *TrieNode) GetWords(word []byte, words *[]string) {
	if len(n.children) == 0 {
		*words = append(*words, string(word))
		return
	} else {
		for _, child := range n.children {
			word = append(word, child.Letter)
			child.GetWords(word, words)
			word = word[:len(word)-1]
		}
	}
}

func (n *TrieNode) Add(letter byte) *TrieNode {
	hi := len(n.children)
	lo := int(0)

	for lo < hi {
		mid := (lo + hi) / 2

		if n.children[mid].Letter == letter {
			return n.children[mid]
		} else if n.children[mid].Letter < letter {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	newNode := &TrieNode{
		Letter:   letter,
		children: make([]*TrieNode, 0),
	}

	if lo == len(n.children) {
		n.children = append(n.children, newNode)
	} else {
		n.children = append(n.children[:lo+1], n.children[lo:]...)
		n.children[lo] = newNode
	}

	return newNode
}

func (n *TrieNode) Search(letter byte) (*TrieNode, error) {
	hi := len(n.children)
	lo := int(0)

	for lo < hi {
		mid := (lo + hi) / 2

		if n.children[mid].Letter == letter {
			return n.children[mid], nil
		} else if n.children[mid].Letter < letter {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return nil, errors.New("unable to find letter in TrieNode children")
}

func (t *Trie) Add(word string) {
	currentNode := &t.root
	for _, letter := range []byte(word) {
		currentNode = currentNode.Add(letter)
	}
}

func (t *Trie) GetWords(word string) []string {
	words := make([]string, 0)

	startNode, err := t.Search(word)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	startNode.GetWords([]byte(word), &words)
	return words
}

func (t *Trie) Search(word string) (*TrieNode, error) {
	current := &t.root
	var err error

	for _, letter := range []byte(word) {
		current, err = current.Search(letter)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, ErrWordNotFound
		}
	}

	return current, nil
}
