package datasource

import (
	cache "chng2016/pkg/datasource/cache"
	localdb "chng2016/pkg/datasource/localDB"
	trie "chng2016/pkg/datasource/trie"
	"chng2016/pkg/models"
)

// Datasource ...
type Datasource interface {
	SetCache(key string, value *models.Distributor)
	GetCache(key string) *models.Distributor

	SetCountryDetails(country *models.Country)
	IsCountryCodeValid(countryCode string) bool
	IsStateCodeValid(stateCode string) bool
	IsCityCodeValid(cityCode string) bool

	AddRegionToTrie(node *models.TrieNode, region string, include bool) error
	HasPermission(node *models.TrieNode, include []string, exclude []string, country, state, city string) bool
}

// DatasourceClient ...
type DatasourceClient struct {
	cache   cache.Cache
	localDB localdb.LocalDB
	trie    trie.Trie
}

// NewDatasourceClient ...
func NewDatasourceClient(cache cache.Cache, localDB localdb.LocalDB, trie trie.Trie) *DatasourceClient {
	return &DatasourceClient{cache: cache, localDB: localDB, trie: trie}
}
