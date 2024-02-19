# Distributor Management System

This project is a distributor management system designed to handle the creation of distributors, sub-distributors, checking permissions, and viewing distributor information. The system utilizes a CSV file containing data about cities, states, and countries for region validation.

## How to Run the code base
Ensure you have Go installed on your system. Navigate to the root directory of the project and run the following command:

```bash
go run main.go
```
Follow the on-screen prompts to interact with the program.

## How to Run the .exe file
To run the executable file (distibuter.exe), follow these steps:

1. Ensure that both the executable file and the cities.csv file are in the same folder.
2. Double-click on the executable file (distibuter.exe).
3. The program will start running, and you can interact with it through the command-line interface.
4. Follow the on-screen prompts to perform various actions like creating distributors, sub-distributors, checking permissions, and viewing distributor information.

## Usage

To run the program, execute the `main.go` file. Upon execution, the program will prompt the user with a menu to select various options:

1. Create a new distributor: Allows the user to create a new distributor.
2. Create a sub-distributor: Allows the user to create a sub-distributor under an existing distributor.
3. Check permission for a distributor: Allows the user to check permissions for a distributor.
4. View Distributors information: Displays information about existing distributors.
5. Exit the program: Exits the program.

### Sample Inputs/Outputs:
> **_NOTE:_**  List of choices is prompted to perform actions based on user input and continues to prompt the user until the program is exited.
- List of choices
```
Use the arrow keys to navigate: ↓ ↑ → ←
Select one of the below choices:
1. Create a new distributor
2. Create a sub distributor
3. Check permission for a distributor
4. View Distributors information     
5. Exit the program
```

- Choosing "Create a new distributor"

```
Enter distributor name: kumaran
Enter the regions you want to include for this distributor: india,china,pakistan
Enter the regions you want to exclude for this distributor: chennai-tamil nadu-india
```

- Choosing "Create a sub distributor"

```
Enter distributor name: poorna
Enter the regions you want to include for this distributor: india
Enter the regions you want to exclude for this distributor: salem-tamil nadu-india
Enter the name of the parent distributor: kumaran
```

- Choosing "Check permission for a distributor"

```
Enter distributor name that needs to be checked: poorna
Enter regions that need to be checked: india
Check Permission Result: [POORNA has access to INDIA]
```

- Choosing "View Distributors information"

```
Distributor Information:
Name: KUMARAN, Include: [INDIA CHINA PAKISTAN], Exclude: [CHENNAI-TAMIL NADU-INDIA], Parent:
Name: POORNA, Include: [INDIA], Exclude: [SALEM-TAMIL NADU-INDIA], Parent: KUMARAN
```

- Choosing "View Distributors information", it will exit the program.

### Author

[Kumaran](https://github.com/kumaranElavazhagn)