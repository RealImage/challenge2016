package models

// Structure to represent a node in the Trie
type TrieNode struct {
	Children map[string]*TrieNode
}
