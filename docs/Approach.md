### Technical approach

1. Parse the cities CSV and build the base permission object as follows:
```
{
	...
	...
	"IN": {
		"TN": {
			"THYAU": false,
			"SWMIL": false		
		},
		"KA": {
			"HU": false,
			"MY": false		
		},
		"TL": {
			"HY": false,
			"TP": false
		},
		...
		...
	},
	"US": {
		"NJ": {
			"BINGH": false,
			"CETRA": false
		},
		"WS": {
			"WA": false,
			"OLYPI": false		
		},
		...
		...
	},
	...
	...
}
```
2. When a distributor or sub-distributor is created, the permissions are cloned from the base permission object.
3. Based on the INCLUDE & EXCLUDE inputs, update the permission object of the distributor. Here is the sample 
4. If the distributor has a parent (if it is a sub-distributor), the INCLUDE & EXCLUDE inputs will be validated against the permission of immediate parent distributor.
5. A distributor is said to have access to a COUNTRY, if all the cities of all the provinces have permission value to true
6. A distributor is said to have access to a PROVINCE, if all the cities of that province have permission value to true
7. A distributor is said to have access to a CITY, if that city have permission value to true.


