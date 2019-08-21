const inquirer = require('inquirer')
const autoComplete = require('autocomplete-cli')

const {
    set, get
} = require('./../store/store')

const {
    removeEntityFromCollection,
    validate,
    getDistributorById
} = require('./util')

const confirmationPrompt = async (message, defaultValue) => {
    try {
        const { value } = await inquirer
        .prompt([
            {
                type: 'confirm',
                name: 'value',
                message,
                default: defaultValue
            }
        ])
        return value
    } catch (error) {
        console.log(`\nError while promting for confirmation to the user :: `, error)
    }
}

const inputPrompt = async (message, required) => {
    try {
        const { value } = await inquirer
        .prompt([
            {
                type: 'input',
                name: 'value',
                message,
                validate: function(value) {
                    if (value) {
                      return true
                    }
                    return 'Please enter a name'
                  }
            }
        ])
        return value
    } catch (error) {
        console.log(`\nError while prompting input from user :: `, error)
    }
}

const listPrompt = async (message, choices) => {
    try {
        const { value } = await inquirer
        .prompt([
            {
                type: 'list',
                name: 'value',
                message,
                choices,
                filter: function (val) {
                    return val.toLowerCase();
                }
            }
        ])
        return value
    } catch (error) {
        console.log(`\nerror while prompting list input to user :: `, error)
    }
}

const promptNewDistributor = async () => {
    try {
        const distributors = get('distributors')
        const distributor = {
            id: distributors.length + 1
        }
        distributor.name = await inputPrompt('Enter distributor name', true)

        if (distributors.length) {
            distributor.isSubDistributor = await confirmationPrompt('Is he a sub distributor ?', false)

            if (distributor.isSubDistributor) {
                const parentDistributor = await listPrompt('choose a distributor', distributors.map((distributor) => {
                    return `${distributor.name}-${distributor.id}`
                }))
                distributor.parentDistributor = parentDistributor.split('-')[1]
            }
        }
        let locations = get('locations', true)
        // get the allowed area
        distributor.allowedLocations = await getAllowedLocation(locations, distributor.parentDistributor)
        // get the restricted area
        locations = get('locations', true)
        distributor.restrictedLocations = await getRestrictedLocation(locations, distributor.parentDistributor)
        if (distributor.parentDistributor) {
            const parentDistributor = getDistributorById(distributor.parentDistributor)
            distributor.restrictedLocations = distributor.restrictedLocations.concat(parentDistributor.restrictedLocations)
        }
        distributors.push(distributor)
        set('distributors', distributors)
        const canExit = await confirmationPrompt('\nDo you want to add another distributor ?', true)
        if (canExit) {
            await promptNewDistributor()
        }
    } catch (error) {
        console.error("\nError while prompting user for the getting distributor related information", error)
        return Promise.reject(error)
    }
}

const getAllowedLocation = async (locations, parentDistributor) => {
    try {
        let selectedLocations = []
        const start = '> Enter allowed location: '
        const selectedLocation = await autoComplete({ start, suggestions: locations })
        const { isValid, message } = await validate(selectedLocation, parentDistributor)
        if (!isValid) {
            console.log('\nLocation entered was not valid ...', message)
            selectedLocations = selectedLocations.concat(await getAllowedLocation(locations, parentDistributor))
            return selectedLocations
        } else {
            selectedLocations.push(selectedLocation)
        }
        locations = removeEntityFromCollection(selectedLocation, locations)
        const addAnotherLocation = await confirmationPrompt('\nDo you want to add another location ?', true)
        if (addAnotherLocation) {
            selectedLocations = selectedLocations.concat(await getAllowedLocation(locations, parentDistributor))
        }
        return selectedLocations
    } catch (error) {
        // need to add comment
        return Promise.reject(error)
    }
}

const getRestrictedLocation = async (locations, parentDistributor) => {
    try {
        let selectedLocations = []
        const start = '> Enter restricted location: '
        const selectedLocation = await autoComplete({ start, suggestions: locations })
        const { isValid, message } = await validate(selectedLocation, parentDistributor)
        if (!isValid) {
            console.log('\nLocation entered was not valid ...', message || '')
            selectedLocations = selectedLocations.concat(await getRestrictedLocation(locations, parentDistributor))
            return selectedLocations
        } else {
            selectedLocations.push(selectedLocation)
        }
        locations = removeEntityFromCollection(selectedLocation, locations)
        const addAnotherLocation = await confirmationPrompt('\nDo you want to add another location ?', true)
        if (addAnotherLocation) {
            selectedLocations = selectedLocations.concat(await getRestrictedLocation(locations, parentDistributor))
        }
        return selectedLocations
    } catch (error) {
        return Promise.reject(error)
    }
}

const promptQuestion = async () => {
    try {
        const distributors = get('distributors')
        let distributor = await listPrompt('choose a distributor', distributors.map((distributor) => {
            return `${distributor.name}-${distributor.id}`
        }))
        distributorId = distributor.split('-')[1]
        distributor = getDistributorById(distributorId)
        const start = '> Enter allowed location: '
        const locations = get('locations', true)
        const selectedLocation = await autoComplete({ start, suggestions: locations })
        const { isValid, message } = validate(selectedLocation, distributorId)
        console.log(`\nThe distributor ${isValid ? '' : 'does not' } has access to this location`)
        const repeat = await confirmationPrompt('\nDo you want to check again ?', false)
        if (repeat) {
            promptQuestion()
        }
    } catch (error) {
        console.log(`\nerror while prompting question to user`, error)
    }
}
module.exports = {
    promptNewDistributor,
    promptQuestion
}