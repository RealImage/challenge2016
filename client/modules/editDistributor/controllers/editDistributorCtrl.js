angular.module('editDistributorCtrl', ['adminService'])
.controller('editDistributorCtrl', function(Admin, $mdDialog, $stateParams) {
	self = this;
	self.id	= $stateParams.id
	self.distributors = []
	getDistributorById(self.id	)
	function getDistributorById(id){
		Admin.getDistributorById(id).then(function (result) {
				if(result.data.status){
					self.data = result.data.data[0]
				}
			})
	}

	function getAllCountries(){
		Admin.getAllCountries().then(function (result) {
				if(result.data.status){
					self.countries = result.data.data.filter(function (country) {
						return country.Country_Name;
					}).map(function (country) {
						country.searchName = angular.lowercase(country.Country_Name)
						return country;
					})
				}
			})
	}
	getAllCountries();

	self.updateDistributor = function(data) {
		Admin.updateDistributor(data).then(function () {
			$mdDialog.hide();
			getAllDistributors()
		})
	};
	self.selectedCountries = []

	self.cities = []
	self.removedCities = []

	self.provinces = []
	self.removedProvinces = []


	self.countySearch  =function (criteria) {
		return criteria ? self.countries.filter(queryFilter(criteria)) : [];
	}
	function queryFilter(query) {
		var lowercaseQuery = angular.lowercase(query);
		return function filterFn(country) {
			return (country.searchName.indexOf(lowercaseQuery) != -1);
		};

	}

	self.provinceSearch  =function (criteria) {
		return criteria ? self.provinces.filter(queryFilter(criteria)) : [];
	}

	self.citySearch  =function (criteria) {
		return criteria ? self.cities.filter(queryFilter(criteria)) : [];
	}




	function getCitiesBasedOnCountrys(countries){
		Admin.getCitiesBasedOnCountrys({

				countries	: countries

			}).then(function (result) {
				if(result.data.status){
					self.cities = result.data.data.filter(function (city) {
						return city.Country_Name;
					}).map(function (city) {
						city.searchName = angular.lowercase(city.City_Name)
						return city;
					})
					var provinvceMap = new Object()
					self.cities.forEach(function (city) {
						provinvceMap[city.Province_Code] ={
							Province_Code : city.Province_Code,
							Province_Name : city.Province_Name,
							searchName	  : angular.lowercase(city.Province_Name)
						};
					})
					self.provinces = Object.keys(provinvceMap).map(function (key) {
						return provinvceMap[key];
					})
				}
				else {
					self.cities = [];
					self.provinces =[];
				}
			})
	}

	self.getCities = function () {
		var countries = [];
		self.data.selectedCountries.forEach(function (data) {
			countries = countries.concat(data.countries)
		})
		getCitiesBasedOnCountrys(countries);
	}

	self.update = function (data) {
		Admin.updateDistributor(data)
	}





});
