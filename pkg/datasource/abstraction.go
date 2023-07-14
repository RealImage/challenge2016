package datasource

import "chng2016/pkg/models"

// SetCache ...
func (d *DatasourceClient) SetCache(key string, value *models.Distributor) {
	d.cache.SetCache(key, value)
}

// GetCache ...
func (d *DatasourceClient) GetCache(key string) *models.Distributor {
	return d.cache.GetCache(key)
}

// IsCountryCodeValid ...
func (d *DatasourceClient) IsCountryCodeValid(countryCode string) bool {
	return d.localDB.IsCountryCodeValid(countryCode)
}

// IsStateCodeValid ...
func (d *DatasourceClient) IsStateCodeValid(stateCode string) bool {
	return d.localDB.IsStateCodeValid(stateCode)
}

// IsCityCodeValid ...
func (d *DatasourceClient) IsCityCodeValid(cityCode string) bool {
	return d.localDB.IsCityCodeValid(cityCode)
}

// SetCountryDetails ...
func (d *DatasourceClient) SetCountryDetails(country *models.Country) {
	d.localDB.SetCountryDetails(country)
}

// AddRegionToTrie ...
func (d *DatasourceClient) AddRegionToTrie(node *models.TrieNode, region string, include bool) error {
	return d.trie.AddRegionToTrie(node, region, include)
}

// HasPermission ...
func (d *DatasourceClient) HasPermission(node *models.TrieNode, include []string, exclude []string, country, state, city string) bool {
	return d.trie.HasPermission(node, include, exclude, country, state, city)
}
