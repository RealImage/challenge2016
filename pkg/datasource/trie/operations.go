package datasource

import (
	"strings"

	"chng2016/pkg/models"
)

// AddRegionToTrie ...
func (t *TrieOperator) AddRegionToTrie(node *models.TrieNode, region string, include bool) error {
	if len(strings.TrimSpace(region)) == 0 {
		return ErrBlankRegionReceived
	}
	segments := strings.Split(region, ",")

	for _, segment := range segments {
		segment = strings.TrimSpace(segment)
		child, ok := node.Children[segment]
		if !ok {
			child = &models.TrieNode{Children: make(map[string]*models.TrieNode)}

			child.Children["*"] = nil
			node.Children[segment] = child
		}

		node = child

	}

	// Mark the leaf node to indicate permission
	if include {
		node.Children["*"] = nil
	} else {
		node.Children["-"] = nil
	}
	return nil
}

// HasPermission ...
func (t *TrieOperator) HasPermission(node *models.TrieNode, include []string, exclude []string, country, state, city string) bool {
	segments := []string{country, state, city}

	for _, segment := range segments {
		segment = strings.TrimSpace(segment)
		child, ok := node.Children[segment]
		if !ok {
			// Check if the leaf node indicates permission or exclusion
			_, okInclude := node.Children["*"]
			_, okExclude := node.Children["-"]

			if okInclude && !okExclude {
				return true
			}
			return false
		}

		node = child
	}

	// Check if the leaf node indicates permission or exclusion
	_, okInclude := node.Children["*"]
	_, okExclude := node.Children["-"]

	if okInclude && !okExclude {
		return true
	}

	return false
}
