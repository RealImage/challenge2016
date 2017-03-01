var Promise					= require('bluebird');

const ERR_NO_COUNTRY_CODE	= 'No country code provided'


/**
* gets the list of cities
* @param {String} data payload
* @author gokul
*/
exports.validateGetAllDistributors = function (data) {

	return Promise.resolve(data)

};

/**
* adds distributor
* @param {String} data payload
* @author gokul
*/
exports.validateAddDistributor = function (data) {
	return Promise.resolve(data)
};

/**
* updates distributor
* @param {String} data payload
* @author gokul
*/
exports.validateUpdateDistributor = function (data) {
	return Promise.resolve(data)
};


/**
* get distributor by id
* @param {String} data payload
* @author gokulklenty
* @since Feb 26, 2017 11:19 PM
*/
exports.validateGetDistributorById = function (data) {
	return Promise.resolve(data);

};

/**
* get distributor by id
* @param {String} data payload
* @author gokulklenty
* @since Feb 26, 2017 11:19 PM
*/
exports.validateSaveSharedLocations = function (data) {
	return Promise.resolve(data);

};
