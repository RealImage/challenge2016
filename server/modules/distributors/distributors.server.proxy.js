var Controller 		= require('./distributors.server.controller');
var ProxyValidator	= require('./distributors.server.pxvalidator');
var Promise 		= require('bluebird');
var ResponseHandler	= require('../../helpers/responseHandler');


/**
* gets the list of distributors
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
* add distributors
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


/**
* updates Distributor
* @param {Object} req Request object
* @param {Object} res Load Landing page
* @author gokul
*/
exports.updateDistributor = function (req,res) {
	// call controller function
	ProxyValidator.validateUpdateDistributor(req.body)
		.then(Controller.updateDistributor)
		.then(ResponseHandler.sendSuccessResponse.bind(null, res))
		.catch(ResponseHandler.sendErrorResponse.bind(null, res))

};

/**
* updates Distributor
* @param {Object} req Request object
* @param {Object} res Load Landing page
* @author gokul
*/
exports.getDistributorById = function (req,res) {
	// call controller function
	ProxyValidator.validateGetDistributorById(req.params)
		.then(Controller.getDistributorById)
		.then(ResponseHandler.sendSuccessResponse.bind(null, res))
		.catch(ResponseHandler.sendErrorResponse.bind(null, res))

};

/**
* updates Distributor
* @param {Object} req Request object
* @param {Object} res Load Landing page
* @author gokul
*/
exports.saveSharedLocations = function (req,res) {
	// call controller function
	ProxyValidator.validateSaveSharedLocations(req.body)
		.then(Controller.saveSharedLocations)
		.then(ResponseHandler.sendSuccessResponse.bind(null, res))
		.catch(ResponseHandler.sendErrorResponse.bind(null, res))

};
