package utils

const (
	// Assign Permissions for Distributor
	WITHOUT_INCLUDES_FIELD           = `{"for": "D2","excludes":[{"province": "BDS","country": "AF"},{"city": "VELOR","province": "TN","country": "IN"},{"city": "KLRAI","province": "TN","country": "IN"}]}`
	WITHOUT_EXCLUDES_FIELD           = `{"for": "D2","includes": ["AF","IN"]}`
	WITHOUT_DISTRIBUTOR_FIELD        = `{"includes": ["AF","IN"],"excludes":[{"province": "BDS","country": "AF"},{"city": "VELOR","province": "TN","country": "IN"},{"city": "KLRAI","province": "TN","country": "IN"}]}`
	VALID_ASSIGN_DISTRIBUTOR_REQUEST = `{"for": "D1","includes": ["AF","IN"],"excludes":[{"province": "BDS","country": "AF"},{"city": "VELOR","province": "TN","country": "IN"},{"city": "KLRAI","province": "TN","country": "IN"}]}`
	WRONG_COUNTRY_CODE               = `{"for": "D1","includes": ["MOMO"],"excludes":[{"province": "BDS","country": "AF"},{"city": "VELOR","province": "TN","country": "IN"},{"city": "KLRAI","province": "TN","country": "IN"}]}`

	//Check Permissions Request for Distributor
	EXCLUDED_DISTRIBUTOR   = `{"city": "KLRAI","province": "TN","country": "IN"}`
	EXCLUDED_CITY          = `{"for": "D1","province": "TN","country": "IN"}`
	EXCLUDED_PROVINCE      = `{"for": "D1","city": "KLRAI","country": "IN"}`
	EXCLUDED_COUNTRY       = `{"for": "D1","city": "KLRAI","province": "TN"}`
	UNKNOWN_DISTRIBUTOR    = `{"for": "D3","city": "KLRAI","province": "TN","country": "IN"}`
	VALID_CHECK_PERMISSION = `{"for": "D1","city": "KLRAI","province": "TN","country": "IN"}`
	WITHOUT_CITY_NAME      = `{"for": "D1", "province": "TN","country": "IN"}`                 // all the cities will be permitted
	WRONG_CITY_CODE        = `{"for": "D1","city": "SDSDSD","province": "TN","country": "IN"}` // Wrong City Code

	//Assign Sub-Distributor Permissions
	EXCLUDED_FROM         = `{"for": "D3","sub_includes": [{"city": "SULUR","province": "TN","country": "IN"},{"city": "CHYAR","province": "TN","country": "IN"}],"excludes":[{"city": "VNASI","province": "TN","country": "IN"}]}`
	EXCLUDED_FOR          = `{"from": "D1","sub_includes": [{"city": "SULUR","province": "TN","country": "IN"},{"city": "CHYAR","province": "TN","country": "IN"}],"excludes":[{"city": "VNASI","province": "TN","country": "IN"}]}`
	EXCLUDED_SUB_INCLUDES = `{"from": "D1","for": "D3","excludes":[{"city": "VNASI","province": "TN","country": "IN"}]}`
	EXCLUDED_EXCLUDES     = `{"from": "D1","for": "D3","sub_includes": [{"city": "SULUR","province": "TN","country": "IN"},{"city": "CHYAR","province": "TN","country": "IN"}]}`

	VALID_SUB_DISTRIBUTION = `{"from": "D1","for": "D3","sub_includes": [{"city": "SULUR","province": "TN","country": "IN"},{"city": "CHYAR","province": "TN","country": "IN"}],"excludes":[{"city": "VNASI","province": "TN","country": "IN"}]}`

	// "BOZEN" cannot be added since "US" is not allocated to Distributor "D1"
	INVALID_SUB_DISTRIBUTION = `{"from": "D1","for": "D3","sub_includes": [{"city": "BOZEN","province": "MT","country": "US"}],"excludes":[{"city": "VNASI","province": "TN","country": "IN"}]}`

	// "SULUR" is valid because "IN" is included in D1 distributor
	VALID_PERMISSION = `{"for": "D3","city": "SULUR","province": "TN","country": "IN"}`

	// All the provinces in the country is valid for sub-distributor
	VALID_DIFF_PROVINCE1 = `{"for": "D3","city": "ARIYU","province": "TN","country": "IN"}`
	VALID_DIFF_PROVINCE2 = `{"for": "D3","city": "ARANI","province": "TN","country": "IN"}`

	// "VNASI" is excluded as it was excluded in distributor D1's data.
	INVALID_PERMISSION1 = `{"for": "D3","city": "VNASI","province": "TN","country": "IN"}`

	// "BOZEN" is not added in D3 distributor, so the authorization will be no
	INVALID_PERMISSION2 = `{"for": "D3","city": "BOZEN","province": "MT","country": "US"}`
)
