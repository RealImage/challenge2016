var Controller 		= require('./distributors.server.controller');
var ProxyValidator	= require('./distributors.server.pxvalidator');
var Promise 		= require('bluebird');
var ResponseHandler	= require('../../helpers/responseHandler');


/**
* gets the list of cities
* @param {Object} req Request object
* @param {Object} res Load Landing page
* @author gokul
*/
exports.getAllDistributors = function (req,res) {

	// call controller function
	ProxyValidator.validateGetAllDistributors()
		.then(Controller.getAllDistributors)
		.then(ResponseHandler.sendSuccessResponse.bind(null, res))
		.catch(ResponseHandler.sendErrorResponse.bind(null, res))

};

/**
* gets the list of cities
* @param {Object} req Request object
* @param {Object} res Load Landing page
* @author gokul
*/
exports.addDistributor = function (req,res) {

	// call controller function
	ProxyValidator.validateAddDistributor(req.body)
		.then(Controller.addDistributor)
		.then(ResponseHandler.sendSuccessResponse.bind(null, res))
		.catch(ResponseHandler.sendErrorResponse.bind(null, res))

};
