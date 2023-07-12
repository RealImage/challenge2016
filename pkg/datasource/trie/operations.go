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
	// fmt.Println("segements : ", segments)
	if len(segments) == 1 {
		// Mark the leaf node to indicate permission
		if include {
			node.Children["*"] = nil
		} else {
			node.Children["-"] = nil
		}
	}

	for _, segment := range segments {
		segment = strings.TrimSpace(segment)
		// fmt.Println("segment : ", segment)
		child, ok := node.Children[segment]
		// fmt.Printf("child node - %#v\n", child)
		// fmt.Println("ok - ", ok)
		if !ok {
			child = &models.TrieNode{Children: make(map[string]*models.TrieNode)}

			// fmt.Printf("not ok child - %#v\n", child)
			child.Children["*"] = nil
			node.Children[segment] = child
			// fmt.Printf("node after not ok set - %#v\n", node)

		}

		node = child
		// fmt.Printf("next process node - %#v\n", node)

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
	// fmt.Printf("root - %#v\n", node)
	segments := []string{country, state, city}

	for _, segment := range segments {
		// fmt.Printf("node - %#v\n", node)
		segment = strings.TrimSpace(segment)
		// fmt.Println("segment : ", segment)
		child, ok := node.Children[segment]
		// fmt.Printf("child - %#v\n", child)
		// fmt.Printf("ok - %#v\n", ok)
		if !ok {
			// Check if the leaf node indicates permission or exclusion
			_, okInclude := node.Children["*"]
			_, okExclude := node.Children["-"]

			if okInclude && !okExclude {
				// fmt.Println("go on")
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
