angular.module('adminService', [])
.factory('Admin', function($http) {

	var adminFactory = {};

	adminFactory.getAllDistributors = function() {
		return $http.get('/api/distributors/getAllDistributors/');
	};

	adminFactory.getAllCountries = function() {
		return $http.get('/api/cities/getAllCountries/');
	};

	adminFactory.getCitiesBasedOnCountrys = function(data) {
		return $http.post('/api/cities/getCities/', data);
	};

	adminFactory.addDistributor = function(data) {
		return $http.post('/api/distributors/addDistributor/', data);
	};

	adminFactory.updateDistributor = function(data) {
		return $http.post('/api/distributors/updateDistributor/', data);
	};

	adminFactory.getDistributorById = function(id) {
		return $http.get('/api/distributors/getDistributorById/'+id);
	};

	adminFactory.saveSharedLocations = function(data) {
		return $http.post('/api/distributors/saveSharedLocations', data);
	};

	return adminFactory;

});
