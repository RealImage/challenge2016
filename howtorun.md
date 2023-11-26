```
> go run main.go
press 0 to exit
enter distrbuter name :
enter a valid distributer
press 0 to exit
enter distrbuter name : distributer
Permissions for : DISTRIBUTER
INCLUDE:
enter a valid include permission

INCLUDE:
enter a valid include permission

INCLUDE:India

INCLUDE:0
ExCLUDE:Tamil Nadu-India

ExCLUDE:0
press 0 to exit
enter distrbuter name : distributer1 < distributer
Permissions for : DISTRIBUTER1<DISTRIBUTER
INCLUDE:Karnataka-India

INCLUDE:Keelakarai-Tamil Nadu-India
Parent distributer dont have access to grant permission- Keelakarai-Tamil Nadu-India

INCLUDE:0
ExCLUDE:Keelakarai-Tamil Nadu-India

ExCLUDE:0
press 0 to exit
enter distrbuter name : 0
INPUT  : [{"Name":"","Permission":"","AuthType":""},{"Name":"","Permission":"","AuthType":""},{"Name":"","Permission":"","AuthType":""},{"Name":"DISTRIBUTER","Permission":"India","AuthType":"include"},{"Name":"DISTRIBUTER","Permission":"Tamil Nadu-India","AuthType":"exclude"},{"Name":"DISTRIBUTER1\u003cDISTRIBUTER","Permission":"Karnataka-India","AuthType":"include"},{"Name":"DISTRIBUTER1\u003cDISTRIBUTER","Permission":"Keelakarai-Tamil Nadu-India","AuthType":"include"},{"Name":"DISTRIBUTER1\u003cDISTRIBUTER","Permission":"Keelakarai-Tamil Nadu-India","AuthType":"exclude"}]
OUTPUT : {"DISTRIBUTER":{"IncludeMap":{"India":"country"},"ExcludeMap":{"Tamil Nadu-India":"state"}},"DISTRIBUTER1":{"IncludeMap":{"Karnataka-India":"state"},"ExcludeMap":{"Keelakarai-Tamil Nadu-India":"city"}}}
```

