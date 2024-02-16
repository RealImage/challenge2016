# Distributor Management System

This project is a distributor management system designed to handle the creation of distributors, sub-distributors, checking permissions, and viewing distributor information. The system utilizes a CSV file containing data about cities, states, and countries for region validation.

## File Structure

The project directory structure is organized as follows:

- **dto/**
  - dto.go
  - city.go
  - state.go
  - country.go

- **input/**
  - inputHandlers.go
  - ...

- **parser/**
  - parser.go
  - ...

- **permission/**
  - permissionChecker.go
  - ...

- **validator/**
  - validator.go
  - ...

- cities.csv
- go.mod
- main.go

css
Copy code

- **dto/**: Contains the data transfer object (DTO) files representing the structure of the data used in the project.
- **input/**: Contains functions to handle user input and interactions.
- **parser/**: Contains the parser functions responsible for parsing the CSV file containing city, state, and country data.
- **permission/**: Contains functions for checking permissions.
- **validator/**: Contains functions for validating distributor and sub-distributor data.
- **cities.csv**: CSV file containing city, state, and country data.
- **go.mod**: Go module file.
- **main.go**: Main entry point of the application.

## Usage

To run the program, execute the `main.go` file. Upon execution, the program will prompt the user with a menu to select various options:

1. Create a new distributor: Allows the user to create a new distributor.
2. Create a sub-distributor: Allows the user to create a sub-distributor under an existing distributor.
3. Check permission for a distributor: Allows the user to check permissions for a distributor.
4. View Distributors information: Displays information about existing distributors.
5. Exit the program: Exits the program.

## How to Run
Ensure you have Go installed on your system. Navigate to the root directory of the project and run the following command:

```bash
go run main.go
```
Follow the on-screen prompts to interact with the program.

### Sample Inputs/Outputs:
> **_NOTE:_**  List of choices is prompted to perform actions based on user input and continues to prompt the user until the program is exited.
- List of choices
```
Select one of the below choices:
1. Create a new distributor
2. Create a sub distributor
3. Check permission for a distributor
4. View Distributors information     
5. Exit the program
```

- Choosing "Create a new distributor" by press 1 and enter

```
Enter distributor name: kumaran
Enter the regions you want to include for this distributor: india,china,pakistan
Enter the regions you want to exclude for this distributor: chennai-tamil nadu-india
```

- Choosing "Create a sub distributor" by press 2 and enter

```
Enter distributor name: poorna
Enter the regions you want to include for this distributor: india
Enter the regions you want to exclude for this distributor: salem-tamil nadu-india
Enter the name of the parent distributor: kumaran
```

- Choosing "Check permission for a distributor" by press 3 and enter

```
Enter distributor name that needs to be checked: poorna
Enter regions that need to be checked: india
Check Permission Result: [POORNA has access to INDIA]
```

- Choosing "View Distributors information" by press 4 and enter

```
Distributor Information:
Name: KUMARAN, Include: [INDIA CHINA PAKISTAN], Exclude: [CHENNAI-TAMIL NADU-INDIA], Parent:
Name: POORNA, Include: [INDIA], Exclude: [SALEM-TAMIL NADU-INDIA], Parent: KUMARAN
```

- Choosing "View Distributors information" by press 5 and enter then it will exit the program.

### Dependencies
This project does not depend on any external libraries or packages beyond the standard Go libraries.

### Contributors
Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

### Author

[kumaran](https://github.com/kumaranElavazhagn)