var Promise					= require('bluebird');

const ERR_NO_COUNTRY_CODE	= 'No country code provided'


/**
* gets the list of cities
* @param {String} data payload
* @author gokul
*/
exports.validateGetCities = function (data) {
	//Country_Code is mandatory

	if(!data.Country_Code){
		return Promise.reject(new Error(ERR_NO_COUNTRY_CODE))
	}
	return Promise.resolve(data)

};

/**
* gets the list of cities
* @param {String} data payload
* @author gokul
*/
exports.validateGetAllCountries = function (data) {
	return Promise.resolve(data)
};
