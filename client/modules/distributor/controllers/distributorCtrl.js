angular.module('distributorCtrl', ['adminService'])
.controller('distributorCtrl', function(Admin, $mdDialog, $stateParams) {
	self = this;
	self.id	= $stateParams.id
	self.distributors = []
	getDistributorById(self.id	)
	function getDistributorById(id){
		Admin.getDistributorById(id).then(function (result) {
				if(result.data.status){
					self.data 	= result.data.data[0]
					self.countries	= [];
					self.data.selectedCountries.forEach(function (data) {
						self.countries = self.countries.concat(data.countries)
					})

					self.countries = self.countries.map(function (country) {
						country.searchName = angular.lowercase(country.Country_Name)
						return country;
					})
					if(!self.data.shared){
						self.data.shared = []
					}
					self.removedProvinceNames = [];
					self.removedCityNames = [];

					if(self.data.removedProvinces){
						self.data.removedProvinces.forEach(function (removedProvince) {
							provinces = removedProvince.provinces.map(function (data) {
								return data.Province_Name;
							});
							self.removedProvinceNames =	self.removedProvinceNames.concat(provinces)
						})

					}

					if(self.data.removedCities){
						self.data.removedCities.forEach(function (removedCity) {
							city = removedCity.cities.map(function (data) {
								return data.City_Name;
							})
							self.removedCityNames = self.removedCityNames.concat(city);
						});


					}



					getAllDistributors(self.data.name);
				}
			})
	}

	function getAllDistributors(currentDistributor) {
		Admin.getAllDistributors().then(function (result) {
			if(result.data.status){
				self.distributors = result.data.data.filter(function (distributor) {
					return distributor.name && distributor.name != currentDistributor;
				})
			}
		})
	}


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
						return city.Province_Name && self.removedProvinceNames.indexOf(city.Province_Name) == -1
									&& city.City_Name && self.removedCityNames.indexOf(city.City_Name) == -1
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


	self.addSharedLocation= function (data){
		var alreadyAssigned = data.shared.map(function (shared) {
			return shared.assignedTo;
		})
		self.distributors = self.distributors.filter(function (distributor) {
			return alreadyAssigned.indexOf(distributor.name) == -1;
		})
		data.shared.push({
			assignedTo : '',
			provinces: [],
			cities: []
		});
	}

	self.saveSharedLocations = function (data) {
		Admin.saveSharedLocations(data)
	}




});
