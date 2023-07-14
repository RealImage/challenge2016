package datasource

import (
	"sync"

	"chng2016/pkg/models"
)

// Trie ...
type Trie interface {
	AddRegionToTrie(node *models.TrieNode, region string, include bool) error
	HasPermission(node *models.TrieNode, include []string, exclude []string, country, state, city string) bool
}

// TrieOperator ...
type TrieOperator struct {
	sync.RWMutex
}

// NewTrie ...
func NewTrie() *TrieOperator {
	return &TrieOperator{}
}
