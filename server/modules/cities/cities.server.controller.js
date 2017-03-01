var DB				= require('../../data/dbmodule');
var Promise 		= require('bluebird');

const CITY_MODEL	= 'cities'
const COUNTRY_MODEL = 'countries'

/**
* gets the list of cities
* @param {Object} data payload
* @author gokul
*/
exports.getCities = function (data) {
	var countryCodes = data.countries.map(function (objCountry) {
		return objCountry.Country_Code;
	});


	objQuery = {Country_Code : {$in: countryCodes}};
	return DB.find(CITY_MODEL, objQuery);

};

/**
* gets the list of cities
* @param {Object} data payload
* @author gokul
*/
exports.getAllCountries = function () {

	return DB.find(COUNTRY_MODEL, {});

};
