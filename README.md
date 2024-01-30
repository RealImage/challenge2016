# Real Image Challenge 2016

This project is a Distributor Management System that allows users to create distributors, sub-distributors, check permissions, and view distributor’s information.

## Getting Started 
### Prerequisites 
- Node.js 
- npm (Node Package Manager) 
- CSV file with geo data (cities.csv) 

### Installation

1. Clone the repository: 
```bash 
git clone https://github.com/Deva45anbu/challenge2016.git 
```
2.	Install dependencies:
```bash
npm install
``` 
3.	Run the application:
```
npm start
```

### Usage
1.	Launch the application by running npm start.
2.	Follow the prompts to
     - Create a new distributor
     - Create a sub-distributor
     - Check permissions of a distributor
     - View distributor information
     - Exit the program.
3.	Enter the required information when prompted.

### File Structure
1. index.js: Main application file.
2. cities.csv: CSV file containing geographical data.
3. README.md: Project documentation. 


### Sample Inputs/Outputs:
> **_NOTE:_**  List of choices is prompted to perform actions based on user input and continues to prompt the user until the program is exited.

- List of choices
```
? Select one of the below choices :
 (Use arrow keys)
> Create a new distributor
  Create a sub distributor
  Check permission for a distributor
  View Distributors information
  Exit the program


```
- Choosing "Create a new distributor"
```
? Enter distributor name: 
 Distributor1
? Enter the regions you want to include for this distributor :
 India,Japan,China
? Enter the regions you want to exclude for this distributor :
 Tamil nadu-india,Maharashtra-India

```

- Choosing "Check permission for a distributor"
```
? Select one of the below choices :
 Check permission for a distributor
? Enter distributor name that need to checked:
 Distributor1
? Enter regions that need to checked:
 India,gujarat-india,chennai-tamil nadu-india,shingu-wakayama-japan
Check Permssion Result : [
  'DISTRIBUTOR1 have access to INDIA',
  'DISTRIBUTOR1 have access to GUJARAT-INDIA',
  'DISTRIBUTOR1 do not have access to CHENNAI-TAMIL NADU-INDIA',
  'DISTRIBUTOR1 have access to SHINGU-WAKAYAMA-JAPAN'
]
```
- Choosing "Create a sub distributor"
```
? Select one of the below choices :
 Create a sub distributor
? Enter distributor name: 
 Distributor2
? Enter the regions you want to include for this distributor :
 india
? Enter the regions you want to exclude for this distributor :
 gujarat-india
? Enter the name of the parent distributor :
 Distributor1
```
- Choosing "Check permission for a distributor"
```
? Select one of the below choices :
 Check permission for a distributor
? Enter distributor name that need to checked:
 Distributor2
? Enter regions that need to checked:
 chennai-tamil nadu-india,gujarat-india,uttar pradesh-india,bihar-india,new zealand
Check Permssion Result : [
  'DISTRIBUTOR2 do not have access to CHENNAI-TAMIL NADU-INDIA',
  'DISTRIBUTOR2 do not have access to GUJARAT-INDIA',
  'DISTRIBUTOR2 have access to UTTAR PRADESH-INDIA',
  'DISTRIBUTOR2 have access to BIHAR-INDIA',
  'DISTRIBUTOR2 do not have access to NEW ZEALAND'
]
```
## Contributing 

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

### Author

[Deva Anbu](https://github.com/Deva45anbu).
