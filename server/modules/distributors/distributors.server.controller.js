var DB				= require('../../data/dbmodule');
var Promise 		= require('bluebird');

const DISTRIBUTORS	= 'distributors';


exports.getAllDistributors = function (data) {
	return DB.find(DISTRIBUTORS, {});
};


exports.addDistributor = function (data) {
	return DB.insert(DISTRIBUTORS, data)
};
