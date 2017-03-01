angular.module('adminCtrl', ['adminService'])
.controller('adminCtrl', function(Admin, $mdDialog) {
	self = this;
	self.distributors = []
	getAllDistributors()
	function getAllDistributors(){
		Admin.getAllDistributors().then(function (result) {
				if(result.data.status){
					self.distributors = result.data.data.filter(function (distributor) {
						return distributor.name;
					})
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


	self.showTabDialog = function(ev) {
		$mdDialog.show({
			controller: DialogController,
			templateUrl: 'modules/admin/views/addDistributor.html',
			parent: angular.element(document.body),
			targetEvent: ev,
			clickOutsideToClose:true
		})
		.then(function(answer) {
			self.status = 'You said the information was "' + answer + '".';
		}, function() {
			self.status = 'You cancelled the dialog.';
		});
	};


	function DialogController($scope, $mdDialog, $element) {
		$scope.hide = function() {
			$mdDialog.hide();
		};

		$scope.countries = self.countries;


		$scope.saveDistributor = function(distributorName, selectedCountries, removedProvinces, removedCities) {
			Admin.addDistributor({
				name				: distributorName,
				selectedCountries	: [{
					assignedBy	: 'ADMIN',
					countries	: selectedCountries
				}],
				removedProvinces	: [{
					assignedBy	: 'ADMIN',
					provinces	: removedProvinces
				}],
				removedCities		: [{
					assignedBy	: 'ADMIN',
					cities		: removedCities
				}]
			}).then(function () {
				$mdDialog.hide();
				getAllDistributors()
			})
		};
		$scope.selectedCountries = []

		$scope.cities = []
		$scope.removedCities = []

		$scope.provinces = []
		$scope.removedProvinces = []


		$scope.countySearch  =function (criteria) {
			return criteria ? $scope.countries.filter(queryFilter(criteria)) : [];
		}
		function queryFilter(query) {
			var lowercaseQuery = angular.lowercase(query);
			return function filterFn(country) {
				return (country.searchName.indexOf(lowercaseQuery) != -1);
			};

		}

		$scope.provinceSearch  =function (criteria) {
			return criteria ? $scope.provinces.filter(queryFilter(criteria)) : [];
		}

		$scope.citySearch  =function (criteria) {
			return criteria ? $scope.cities.filter(queryFilter(criteria)) : [];
		}




		function getCitiesBasedOnCountrys(countries){
			Admin.getCitiesBasedOnCountrys({

					countries	: countries

				}).then(function (result) {
					if(result.data.status){
						$scope.cities = result.data.data.filter(function (city) {
							return city.Country_Name;
						}).map(function (city) {
							city.searchName = angular.lowercase(city.City_Name)
							return city;
						})
						var provinvceMap = new Object()
						$scope.cities.forEach(function (city) {
							provinvceMap[city.Province_Code] ={
								Province_Code : city.Province_Code,
								Province_Name : city.Province_Name,
								searchName	  : angular.lowercase(city.Province_Name)
							};
						})
						$scope.provinces = Object.keys(provinvceMap).map(function (key) {
							return provinvceMap[key];
						})
					}
					else {
						$scope.cities = [];
						$scope.provinces =[];
					}
				})
		}


		$scope.$watchCollection('selectedCountries', function (newVal) {
			getCitiesBasedOnCountrys(newVal);
		})

	}

});
