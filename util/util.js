
const {
    get
} = require('./../store/store')

const removeEntityFromCollection = (stringValue, collection = []) => {
    const index = collection.indexOf(stringValue)
    if (index > -1) {
        collection.splice(index, 1)
    }
    return collection
}

const validate = (selectedLocation, parentDistributorId) => {
    let validationResponse = {} 
    if (!parentDistributorId) {
        validationResponse.isValid = true
        return validationResponse
    }
    validationResponse = validateLocations(selectedLocation, parentDistributorId, 'add')
    if (validationResponse.isPresent) {
        const { isPresent, message } = validateLocations(selectedLocation, parentDistributorId, 'restrict')
        validationResponse.isValid = !isPresent
    } else {
        validationResponse.isValid = validationResponse.isPresent       
    }
    return validationResponse
}

const validateLocations = (selectedLocation, parentDistributorId, type) => {
    const response = {
        isPresent: false
    }
    if (!selectedLocation) {
        return response
    }
    const { allowedLocations, restrictedLocations } = getDistributorById(parentDistributorId)
    const locations = type === 'add' ? allowedLocations : restrictedLocations
    const selectedLocationSplit = selectedLocation.split('-')
    const selectedLocationSplitLength = selectedLocationSplit.length
    switch (selectedLocationSplitLength) {
        case 1:
            response.isPresent = locations.indexOf(selectedLocation) > -1
            break
        case 2:
            response.isPresent = locations.indexOf(selectedLocation) > -1
            if (!response.isPresent) {
                response.isPresent = locations.indexOf(selectedLocationSplit[1]) > -1
            }
            break
        case 3:
            response.isPresent = locations.indexOf(selectedLocation) > -1
            if (!response.isPresent) {
                response.isPresent = locations.indexOf(selectedLocationSplit[1]+'-'+selectedLocationSplit[2]) > -1
                if (!response.isPresent) {
                    response.isPresent = locations.indexOf(selectedLocationSplit[2]) > -1
                }
            }
            break
    }
    if (!response.isPresent) {
        response.message = `\nThe parent distributor does not have access to the selected country :: ${selectedLocation}. So can Provide access to the selected country.`
    }
    return response
}

const getDistributorById = (distributorId) => {
    const distributors = get('distributors')
    const index = distributors.findIndex(function(localdistributor) {
        return localdistributor.id == distributorId
    })
    let distributor
    if (index > -1) {
        distributor = distributors[index]
    }
    return distributor
}

module.exports = {
    removeEntityFromCollection,
    validate,
    getDistributorById
}