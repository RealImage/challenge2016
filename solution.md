# APIs For the Solution

### To assign permissions for a Distributor
The following API needs to be posted for assigning a new distributor with include and exclude permissions

```
curl -i -X POST \
   -H "Content-Type:application/json" \
   -d \
'{
  "for": "D2",
  "includes": ["AF","IN"],
  "excludes":[
    {
    "province": "BDS",
    "country": "AF"
  },
    {
     "city": "VELOR",
    "province": "TN",
    "country": "IN"
  },
    {
     "city": "KLRAI",
    "province": "TN",
    "country": "IN"
  }
  ]
}' \
 'http://localhost:8000/assign_distributor'
 
```


Parameters | Type | Description
--- | --- | ---
for | `string` | distributor ID
includes | array of string | contains only the countries code
excludes | array of object | contains the excluded city, province and country code

### To assign permissions for a Sub Distributor
The following API needs to be posted for assigning a new subdistributor with include and exclude permissions. And also we have to give distributor ID

```
curl -i -X POST \
   -H "Content-Type:application/json" \
   -d \
'{
  "from": "D1",
  "for": "D3",
  "sub_includes": [
    {
      "city": "SULUR",  
    "province": "TN",
    "country": "IN"
  },
    {
     "city": "CHYAR",
    "province": "TN",
    "country": "IN"
  }
  ],
  "excludes":[
   {
     "city": "VNASI",
    "province": "TN",
    "country": "IN"
  }
  ]
}' \
 'http://localhost:8000/assign_sub_distributor'
 
```
Parameters | Type | Description
--- | --- | ---
from| string | distributor ID
for | `string` | distributor ID
sub_includes | array of string | contains the included city, province and country code
excludes | array of object | contains the excluded city, province and country code

### To check for permission authorization in particular region
To check the distributor's permission to distribute the film in that particular region

```
curl -i -X POST \
   -H "Content-Type:application/json" \
   -d \
'{
  "for": "D3",
  "city": "SRNGA",
  "province": "MP",
  "country": "IN"
}
' \
 'http://localhost:8000/check_permission'
```

Parameters | Type | Description
--- | --- | ---
for | `string` | distributor ID
city | string | contains  the city code
province | string | contains the province code
country | string | contains the country code

For the assigning the distributor and sub-distributor, the response will be the following, even if you send the duplicate data
```
{ 
   Status: "Distributor permissions successfully assigned!"
}
```

For checking the permission authorization, the response will be the following
```
{
    "is_authorized": "no"
}
```


Features Completed:
1. Parsed CSV file and stored in a map
2. Assigned and saved a distributor in a map
3. Validated the include and exclude city codes before assigning for distributors
4. Assigned sub-distributor permissions from distributor's permission
5. Checked the permission for the particular distributor on that given region

The programs may have repetitive code. It can be optimized if there is more time.

