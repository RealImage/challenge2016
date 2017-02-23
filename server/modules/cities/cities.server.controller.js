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
	var objQuery = new Object();
	objQuery.Country_Code = data.Country_Code;

	if(data.Province_Code){
		objQuery.Province_Code	= data.Province_Code;
	}

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
