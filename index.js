// Import required modules
const fs = require('fs');               // File system module for reading files
const { parse } = require('csv-parse'); // CSV parsing module
const express = require('express');     // Express.js for creating a web server
const inquirer = require('inquirer');   // Inquirer.js for interactive command-line prompts
const prompt = inquirer.createPromptModule(); // Create a prompt module using Inquirer

const app = express();
const port = 3001;
// app.use(express.json());


let distributorInformation = [];
let geoData

// The main function that initializes the program, reads CSV file, initializes geoData, asks questions to the use
async function main() {
    try {
        const csvFilePath = '\cities.csv';
        geoData = await parseCSVFile(csvFilePath)
        await askNextQuestion();
        console.log("Exiting the program");
    } catch (error) {
        console.log("Error in main function :", error)
    }

}

//Parse a CSV file and group records by country, states and cities in the state.
async function parseCSVFile(csvFilePath) {
    return new Promise((resolve, reject) => {

        // Read the CSV file using fs
        fs.readFile(csvFilePath, 'utf8', (err, data) => {
            if (err) {
                console.error('Error reading the file:', err);
                reject(err);
            }

            // Parse the CSV content using csv-parse
            parse(data, {
                columns: true, // Treat the first row as headers
                skip_empty_lines: true // Skip empty lines
            }, (err, records) => {
                if (err) {
                    console.error('Error parsing the CSV:', err);
                    return;
                }
                try {
                    // Group records by country, states and cities in the state.
                    const groupedData = records.reduce((acc, city) => {
                        const countryIndex = acc.findIndex(item => item.country === city['Country Name'].toUpperCase());

                        if (countryIndex === -1) {
                            // If country doesn't exist, add it with the first state and city
                            acc.push({
                                country: city['Country Name'].toUpperCase(),
                                states: [
                                    {
                                        state: city['Province Name'].toUpperCase(),
                                        cities: [city['City Name'].toUpperCase()]
                                    }
                                ]
                            });
                        } else {
                            const stateIndex = acc[countryIndex].states.findIndex(item => item.state === city['Province Name'].toUpperCase());
                            if (stateIndex === -1) {
                                // If state doesn't exist, add it with the first city
                                acc[countryIndex].states.push({
                                    state: city['Province Name'].toUpperCase(),
                                    cities: [city['City Name'].toUpperCase()]
                                });
                            } else {
                                // State exists, add the city to the existing state
                                acc[countryIndex].states[stateIndex].cities.push(city['City Name'].toUpperCase());
                            }
                        }
                        return acc;
                    }, []);
                    resolve(groupedData);
                } catch (error) {
                    console.error('Error grouping the data:', error);
                    reject(error);
                }
            });
        });
    })
}

// askNextQuestion() prompt the user with a list of choices and perform actions based on user input and continues to prompt the user until the program is exited.
async function askNextQuestion() {
    while (true) {
        try {

            const questions = [
                {
                    type: 'list',
                    name: 'NewDistributor',
                    message: 'Select one of the below choices :\n',
                    choices: ['Create a new distributor', 'Create a sub distributor',
                        "Check permission for a distributor", "View Distributors information", "Exit the program"]
                }
            ];

            const getDistributorDataQuestions = [
                {
                    type: 'input',
                    name: 'name',
                    message: 'Enter distributor name: \n'
                },
                {
                    type: 'input',
                    name: 'include',
                    message: 'Enter the regions you want to include for this distributor :\n'
                },
                {
                    type: 'input',
                    name: 'exclude',
                    message: 'Enter the regions you want to exclude for this distributor :\n'
                }]
            const getSubDistributorDataQuestions = [
                {
                    type: 'input',
                    name: 'name',
                    message: 'Enter distributor name: '
                },
                {
                    type: 'input',
                    name: 'include',
                    message: 'Enter the regions you want to include for this distributor :\n'
                },
                {
                    type: 'input',
                    name: 'exclude',
                    message: 'Enter the regions you want to exclude for this distributor :\n'
                },
                {
                    type: 'input',
                    name: 'parentName',
                    message: 'Enter the name of the parent distributor :\n'
                }]
            const checkPermissionData = [
                {
                    type: 'input',
                    name: 'distributorName',
                    message: 'Enter distributor name that need to checked:\n'
                }, {
                    type: 'input',
                    name: 'testData',
                    message: 'Enter regions that need to checked:\n'
                }]
            const answer = await prompt(questions);
            if (answer.NewDistributor == "Create a new distributor") {
                const distributorData = await prompt(getDistributorDataQuestions);
                let errors = validateDistributorData(distributorData);
                if (errors.length > 0) {
                    console.error("Validation errors:", errors);
                    continue;
                }
                let distributorObject = createNewDistributor(distributorData)

                // Add a valid distributor data in the distributorInformation array
                distributorInformation.push(distributorObject)


            } else if (answer.NewDistributor == "Create a sub distributor") {
                const subDistributorData = await prompt(getSubDistributorDataQuestions);
                let errors = validateSubDistributorData(subDistributorData);
                if (errors.length > 0) {
                    console.error("Validation errors:", errors);
                    continue;
                }
                let parentDistributorData = getDistributorData(subDistributorData.parentName.toUpperCase())
                if (parentDistributorData.exclude.length != 0) {
                    subDistributorData.exclude = subDistributorData.exclude != "" ?
                        subDistributorData.exclude + ',' + parentDistributorData.exclude.join() : parentDistributorData.exclude.join();
                }

                let subDistributorObject = createNewDistributor(subDistributorData, subDistributorData.parentName.toUpperCase())
                // Add a valid distributor data in the distributorInformation array
                distributorInformation.push(subDistributorObject)

            } else if (answer.NewDistributor == "Check permission for a distributor") {
                let errorMsg = []
                const checkData = await prompt(checkPermissionData);
                let errors = validateCheckPermissionData(checkData)
                if (errors.length > 0) {
                    console.error("Validation errors:", errors);
                    continue;
                }
                testData = checkData.testData.toUpperCase().split(',')
                let checkPermssionResult = checkPermission(checkData.distributorName.trim(), testData, "checkDistibutorPermission")
                console.log("Check Permssion Result :", checkPermssionResult)



            } else if (answer.NewDistributor == "View Distributors information") {
                displayDistributorInformation()
            } else if (answer.NewDistributor == "Exit the program") {
                // Resolve to exit the loop
                process.exit(0);

            }

        } catch (err) {
            console.error("Error in askNextQuestion():", err);
            // Don't resolve; continue to the next iteration of the loop
            continue;
        }
    }


}

// validateDistributorData function checks if the distributor name is not empty, unique, and if include and exclude regions are valid based on the geoData array.
function validateDistributorData(data) {
    try {
        let errorMsg = []
        if (data.name.trim() == '') {
            errorMsg.push("Distributor Name must not be empty, please enter a valid distributor name")
        } else if (validateDistributorName(data.name.trim().toUpperCase())) {
            errorMsg.push("Distributor Name already exists")
        }

        if (data.include.trim() == '') {
            errorMsg.push("Include Regions must not be empty, please enter a valid regions")
        } else if (!validateRegions(data.include.trim())) {
            errorMsg.push("Include Region is not present in csv, please enter a valid region")
        }

        if (data.exclude.trim() != '') {
            if (!validateRegions(data.exclude.trim())) {
                errorMsg.push("Exclude Region is not present in csv, please enter a valid region")
            }
        }
        return errorMsg;
    } catch (error) {
        console.error('Error in validateDistributorData() :', error);
    }
}


// Validate sub-distributor data to ensure it meets the required criteria
function validateSubDistributorData(data) {
    try {
        let errorMsg = []
        if (data.name.trim() == '') {
            errorMsg.push("Distributor Name must not be empty, please enter a valid distributor name")
        } else if (validateDistributorName(data.name.trim().toUpperCase())) {
            errorMsg.push("Distributor Name already exists")
        }

        if (data.include.trim() == '') {
            errorMsg.push("Include Regions must not be empty, please enter a valid regions")
        } else if (!validateRegions(data.include.trim())) {
            errorMsg.push("Include Region is not present in csv, please enter a valid region")
        }

        if (data.exclude.trim() != '') {
            if (!validateRegions(data.exclude.trim())) {
                errorMsg.push("Exclude Region is not present in csv, please enter a valid region")
            }
        }

        if (data.parentName.trim() == '') {
            errorMsg.push("Parent distributor Name must not be empty, please enter a valid parent distributor name")
        } else if (!validateDistributorName(data.parentName.trim().toUpperCase())) {
            errorMsg.push("Parent distributor Name does not exists, please enter existing parent distributor name")
        }
        if (errorMsg.length == 0) {

            let testData = data.exclude.trim() != '' ?
                (data.include.trim() + ',' + data.exclude.trim()).toUpperCase().split(',')
                : data.include.trim().toUpperCase().split(',')
            let checkPermissionWithParent = checkPermission(data.parentName.trim(), testData, "subDistributionCreation")
            if (checkPermissionWithParent.length > 0) {
                errorMsg = [...errorMsg, ...checkPermissionWithParent];
            }

        }


        return errorMsg;
    } catch (error) {
        console.error('Error in validateSubDistributorData() :', error);
    }

}

//  Validate distributor name to ensure it is unique.
function validateDistributorName(distributorName) {
    try {
        if (distributorInformation.length > 0) {
            for (let i = 0; i < distributorInformation.length; i++) {
                if (distributorInformation[i].distributorName === distributorName) {
                    return true;
                }
            }
            return false;
        } else {
            return false;
        }
    } catch (error) {
        console.error('Error in validateDistributorName() :', error);
    }
}


// validateRegions function checks if the specified regions exist in the geoData array(csv file data).
function validateRegions(data) {
    try {
        splitTestData = data.split(',')
        for (let i = 0; i < splitTestData.length; i++) {
            let testData = splitTestData[i].split('-').map(part => part.toUpperCase());

            if (testData.length > 0 && testData.length < 4) {
                if (testData.length == 1) {
                    for (let i = 0; i < geoData.length; i++) {
                        if (geoData[i].country == testData[0]) {
                            return true
                        }
                    }
                    return false
                } else if (testData.length == 2) {
                    for (let i = 0; i < geoData.length; i++) {
                        if (geoData[i].country == testData[1]) {
                            for (let j = 0; j < geoData[i].states.length; j++) {
                                if (geoData[i].states[j].state == testData[0]) {
                                    return true
                                }
                            }
                        }
                    }
                    return false
                } else if (testData.length == 3) {
                    for (let i = 0; i < geoData.length; i++) {
                        if (geoData[i].country == testData[2]) {
                            for (let j = 0; j < geoData[i].states.length; j++) {
                                if (geoData[i].states[j].state == testData[1]) {
                                    if (geoData[i].states[j].cities.includes(testData[0])) {
                                        return true
                                    }
                                }
                            }
                        }
                    }
                    return false
                }
            } else {
                return false
            }
        }
    } catch (error) {
        console.error('Error in validateRegions() :', error);
    }
}

// createNewDistributor function constructs a new distributor object with the distributorName, include regions, exclude regions and parent for sub distributor object.
function createNewDistributor(data, parent) {
    try {
        let distributorData = {
            distributorName: data.name.toUpperCase(),
            include: data.include.trim().toUpperCase().split(','),
            exclude: data.exclude.trim() != "" ? data.exclude.toUpperCase().split(',') : [],
            parent: parent
        };
        return distributorData
    } catch (error) {
        console.error('Error in createNewDistributor() :', error);
    }
}

// getDistributorData function searches for the distributor with the specified name in the distributorInformation array.
function getDistributorData(distributorName) {
    try {
        for (let i = 0; i < distributorInformation.length; i++) {
            if (distributorInformation[i].distributorName == distributorName.toUpperCase()) {
                return distributorInformation[i]
            }
        }
    } catch (error) {
        console.error('Error in getDistributorData() :', error);
    }
}

// validateCheckPermissionData function checks if the provided region data for checking permissions of the distributor is valid.
function validateCheckPermissionData(data) {
    try {
        let errorMsg = []
        if (data.distributorName.trim() == '') {
            errorMsg.push("Distributor Name must not be empty, please enter a valid distributor name")
        } else if (!validateDistributorName(data.distributorName.trim().toUpperCase())) {
            errorMsg.push("Distributor name does not exists")
        }
        let testData = data.testData.split(',')
        for (let i = 0; i < testData.length; i++) {
            if (!validateRegions(testData[i])) {
                errorMsg.push(testData[i].toUpperCase() + " does not exists in the csv file, please enter a valid region")
            }
        }
        return errorMsg
    } catch (error) {
        console.log("Error in validateCheckPermissionData() :", error)
    }

}

// checkPermission function evaluates whether a distributor has access to specified regions based on the testData.
function checkPermission(distributorName, testData, origin) {
    try {
        let validationResult = [], errorMsg = [];
        let distributorData = getDistributorData(distributorName)
        for (let i = 0; i < testData.length; i++) {
            let newTestData = testData[i].split('-');
            if (newTestData.length == 1) {
                if (distributorData.include.includes(testData[i])) {
                    validationResult.push(distributorData.distributorName + " have access to " + testData[i])
                } else {
                    validationResult.push(distributorData.distributorName + " do not have access to " + testData[i])
                    errorMsg.push(distributorData.distributorName + " do not have access to " + testData[i])
                }
            }
            else if (newTestData.length == 2) {
                if (distributorData.include.includes(newTestData[1])) {
                    if (distributorData.exclude.includes(testData[i])) {
                        validationResult.push(distributorData.distributorName + " do not have access to " + testData[i])
                        errorMsg.push(distributorData.distributorName + " do not have access to " + testData[i])
                    } else {
                        validationResult.push(distributorData.distributorName + " have access to " + testData[i])
                    }
                } else if (distributorData.include.includes(testData[i])) {
                    validationResult.push(distributorData.distributorName + " have access to " + testData[i])

                } else {
                    validationResult.push(distributorData.distributorName + " do not have access to " + testData[i])
                    errorMsg.push(distributorData.distributorName + " do not have access to " + testData[i])
                }
            } else if (newTestData.length == 3) {
                if (distributorData.include.includes(newTestData[2])) {
                    if (distributorData.include.includes(newTestData[1] + '-' + newTestData[2])) {
                        if (distributorData.include.includes(testData[i])) {
                            validationResult.push(distributorData.distributorName + " have access to " + testData[i])
                        } else {
                            if (distributorData.exclude.includes(testData[i])) {
                                validationResult.push(distributorData.distributorName + " do not have access to " + testData[i])
                                errorMsg.push(distributorData.distributorName + " do not have access to " + testData[i])
                            } else {
                                validationResult.push(distributorData.distributorName + " have access to " + testData[i])
                            }
                        }
                    } else {
                        if (distributorData.exclude.includes(newTestData[1] + '-' + newTestData[2])) {
                            validationResult.push(distributorData.distributorName + " do not have access to " + testData[i])
                            errorMsg.push(distributorData.distributorName + " do not have access to " + testData[i])
                        } else {
                            if (distributorData.exclude.includes(testData[i])) {
                                validationResult.push(distributorData.distributorName + " do not have access to " + testData[i])
                                errorMsg.push(distributorData.distributorName + " do not have access to " + testData[i])
                            } else {
                                validationResult.push(distributorData.distributorName + " have access to " + testData[i])
                            }
                        }
                    }
                }
                else {
                    if (distributorData.include.includes(newTestData[1] + '-' + newTestData[2])) {
                        if (distributorData.exclude.includes(testData[i])) {
                            validationResult.push(distributorData.distributorName + " do not have access to " + testData[i])
                            errorMsg.push(distributorData.distributorName + " do not have access to " + testData[i])
                        } else {
                            validationResult.push(distributorData.distributorName + " have access to " + testData[i])
                        }
                    } else if (distributorData.include.includes(testData[i])) {
                        validationResult.push(distributorData.distributorName + " have access to " + testData[i])
                    } else {
                        validationResult.push(distributorData.distributorName + " do not have access to " + testData[i])
                        errorMsg.push(distributorData.distributorName + " do not have access to " + testData[i])
                    }
                }
            }
        }
        return origin == "subDistributionCreation" ? errorMsg : validationResult;
    } catch (error) {
        console.log("Error in checkPermission() :", error)
    }
}

// displayDistributorInformation function prints the distributor information to the console for informational purposes.
function displayDistributorInformation() {
    try {
        console.log("Distributor Information", distributorInformation)
    } catch (error) {
        console.log("Error in displayDistributorInformation()", error)
    }
}

app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
    main()
});
