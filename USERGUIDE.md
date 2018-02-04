# Creating Distributors and Sub-Distributors with customized access

## Creating a Direct Distributor

Run the program by executing the command **go run main.go**

Once the code is executed you will be prompted in the console to enter the `DISTRIBUTOR` name as below

```
By default you need to create a Direct Distributor initially
Please enter your PERMISSION following the order Country, State, and the City
use `_` to separate Country Name, State Name and City Name (not case sensitive)
Sample: EXCLUDE: CHENNAI_TAMIL nadu_INDIa 

Enter the Distributor name:

```

you can type the `DISTRIBUTOR` name in the console, followed by that we can give all our **PERMISSIONS** that we need to assign for the Distributor by typing a PERMISSION with a prefix `INCLUDE:` or `EXCLUDE:` and clicking `ENTER` and this order should be followed for all the rules and once you are done click the `ENTER` button twice to submit for validation and storing.

```
Enter the Distributor name: DISTRIBUTOR1
Please list the permissions,
INCLUDE: INDIA
INCLUDE: CHINA
EXCLUDE: KERALA_INDIA
EXCLUDE: TAMIL NADU_INDIA
INCLUDE: Al-Hamdaniya_Ninawa_Iraq
INCLUDE: Istgah-e Rah Ahan-e Garmsar_Semnan_Iran

```



**Possible Response:**

In case if the rules we have entered are valid, console will be responding with 

```
Distributor created successfully !!

1. Continue adding distributors?
2. View the existing distributors?
Please select your choice:
```

In case if we have missed entering the `INCLUDE:` or `EXCLUDE:` for any of the `PERMISSION` - console will throw a message mentioning the invalid `PERMISSION` meanwhile all other permissions you have entered will be neglected

```
[chennai_tamil nadu_india] Not permitted - mention the INCLUDE/EXCLUDE operation
1. Continue adding distributors?
2. View the existing distributors? 
```

In case if you are trying to Add/Remove a **STATE/AREA** to which you have already created a `PERMISSION` saying Distributor is subjected to have/don't have access, your current input will be rejected you need to correct that and try again.


## Creating a Sub - Distributor

You can straight away create a Sub-Distributor by following the below-mentioned step.

```
1. Continue adding distributors?
2. View the existing distributors?
Please select your choice: 1

Enter the distributor type (Direct/Sub): Sub

Enter the Parent - Distributor name: DISTRIBUTOR1
Enter the Sub - Distributor name: DISTRIBUTOR2
Please list the permissions,
INCLUDE: Al-Hamdaniya_Ninawa_Iraq
INCLUDE: Istgah-e Rah Ahan-e Garmsar_Semnan_Iran
EXCLUDE: Maharashtra_India
INCLUDE: Chhattisgarh_india
INCLUDE: Fujian_China
INCLUDE: Shanhaiguan_Hebei_China
EXCLUDE: Yunnan_China

Sub-Distributor created successfully !!
```


**Key points we need focus while creating a Sub-Distributor:**

* Enter the correct Parent Distributor, in error case you will be redirected to the selection (Case sensitive)
* All Distributor and Sub-Distributor names are unique, in error case, you will be redirected to the selection
* Any issue with the `PERMISSION` will be will be notified to you referring to the actual `PERMISSION` and this will not be saved in this case


## Listing existing Distributors and checking for permission

To list all the existing Distributors created and check for a permission

```
1. Continue adding distributors?
2. View the existing distributors?
Please select your choice: 2
Direct Distributors:
 
DISTRIBUTOR1
DISTRIBUTOR2
DISTRIBUTOR3


Sub Distributors:

DISTRIBUTOR4 > DISTRIBUTOR1
DISTRIBUTOR5 > DISTRIBUTOR1
DISTRIBUTOR6 > DISTRIBUTOR2
DISTRIBUTOR7 > DISTRIBUTOR3
DISTRIBUTOR7 > DISTRIBUTOR4

Please enter any Distributor name to check permission: DISTRIBUTOR2
Please enter the permission to check whether it is valid or not:
INCLUDE: Chhattisgarh_india

Valid permission !!
```


The user will be notified whether the PERMISSION is valid or not and any other error case that has occurred.


## Test case

I have also included the test case to validate the main core functionality which will receive the object and query the data based on the distributor role and return the message excluded the user interaction part.

**Steps:**

```
go get github.com/stretchr/testify
go test ./... -v
```
