var Controller 		= require('./cities.server.controller');
var ProxyValidator	= require('./cities.server.pxvalidator');
var Promise 		= require('bluebird');
var ResponseHandler	= require('../../helpers/responseHandler');


/**
* gets the list of cities
* @param {Object} req Request object
* @param {Object} res Load Landing page
* @author gokul
*/
exports.getCities = function (req,res) {

	// call controller function
	ProxyValidator.validateGetCities(req.body)
		.then(Controller.getCities)
		.then(ResponseHandler.sendSuccessResponse.bind(null, res))
		.catch(ResponseHandler.sendErrorResponse.bind(null, res))

};

/**
* gets the list of cities
* @param {Object} req Request object
* @param {Object} res Load Landing page
* @author gokul
*/
exports.getAllCountries = function (req,res) {

	// call controller function
	ProxyValidator.validateGetAllCountries()
		.then(Controller.getAllCountries)
		.then(ResponseHandler.sendSuccessResponse.bind(null, res))
		.catch(ResponseHandler.sendErrorResponse.bind(null, res))

};
