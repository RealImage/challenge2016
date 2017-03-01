var DB				= require('../../data/dbmodule');
var Promise 		= require('bluebird');

const DISTRIBUTORS	= 'distributors';


exports.getAllDistributors = function (data) {
	return DB.find(DISTRIBUTORS, {});
};


exports.addDistributor = function (data) {
	return DB.insert(DISTRIBUTORS, data)
};


exports.updateDistributor = function (data) {
	var id			= data._id;
	return DB.update(DISTRIBUTORS,{_id: id}, {$set: data})
};


/**
* get distributor by id
* @param {String} data payload
* @author gokulklenty
* @since Feb 26, 2017 11:19 PM
*/
exports.getDistributorById = function (data) {

	return DB.find(DISTRIBUTORS, {_id: data.id});

};

/**
* get distributor by id
* @param {String} data payload
* @author gokulklenty
* @since Feb 26, 2017 11:19 PM
*/
exports.saveSharedLocations = function (data) {

	var id			= data._id;
	return DB.update(DISTRIBUTORS,{_id: id}, {$set: data}).then(function () {
		return Promise.map(data.shared
								, saveSharedLocations.bind(null, data.name));
	})

};

var saveSharedLocations = function (assignedBy, data) {

	return DB.find(DISTRIBUTORS, {name: data.assignedTo}).then(function (result) {
		var distributor		= result[0];
		var alreadyAssigned = false;
		if(!distributor.includedRegions){
			distributor.includedRegions	= [{
				assignedBy	: assignedBy,
				cities 		: data.cities,
				provinces	: data.provinces
			}];
		}
		else {
			distributor.includedRegions = distributor.includedRegions.map(function (region) {
				if(region.assignedBy == assignedBy){
					region.cities 		= data.cities;
					region.provinces	= data.provinces;
					alreadyAssigned		= true;
				}
				return region;
			});

			if(!alreadyAssigned){
				distributor.includedRegions.push({
					assignedBy	: assignedBy,
					cities 		: data.cities,
					provinces	: data.provinces,
				})
			}
		}
		return DB.update(DISTRIBUTORS,{_id: distributor._id}, {$set: distributor});

	})
}
