	
	var distApp = angular.module('distApp', ['ngRoute']);
  var domain_path = "http://localhost:3000"

  var form_transformation = function(obj) {
	        var str = [];
	        for(var p in obj)
	        str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
	        return str.join("&");
	    	}
	// configure our routes
	distApp.config(function($routeProvider) {
		$routeProvider

			// route for the home page
			.when('/', {
				templateUrl : 'pages/home.html',
				controller  : 'mainController'
			})
		.when('/countries', {
				templateUrl : 'pages/countries-index.html',
				controller  : 'countriesIndexController'
			})
		.when('/countries/new', {
				templateUrl : 'pages/countries-new.html',
				controller  : 'countriesNewController'
			})
		.when('/countries/:country_id', {
				templateUrl : 'pages/country-show.html',
				controller  : 'countriesShowController'
			})
		.when('/states/new/:country_id', {
				templateUrl : 'pages/states-new.html',
				controller  : 'statesNewController'
			})
		.when('/states/:state_id', {
				templateUrl : 'pages/state-show.html',
				controller  : 'statesShowController'
			})
		.when('/cities/new/:state_id', {
				templateUrl : 'pages/cities-new.html',
				controller  : 'citiesNewController'
			})
		.when('/distributors', {
				templateUrl : 'pages/distributors-index.html',
				controller  : 'distributorsIndexController'
			})
		.when('/distributors/new', {
				templateUrl : 'pages/distributors-new.html',
				controller  : 'distributorsNewController'
			})
		.when('/distributors/new/:distributor_id', {
				templateUrl : 'pages/distributors-new.html',
				controller  : 'distributorsNewController'
			})
			
	});

	distApp.controller('mainController', function($scope) {});

	distApp.controller('countriesIndexController', function($scope, $http) {
		$http.get(domain_path+"/countries").success(function(data) {
  		$scope.countries = data.countries
  	})
	});
	distApp.controller('countriesNewController', function($scope, $http) {
		$scope.country = {}
		$scope.submitForm = function(){
			$http({
				url: domain_path+"/countries", 
				method: "post",
				data: $scope.country ,
				transformRequest: form_transformation,
        headers: {'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8',"Accept":"application/x-www-form-urlencoded;charset=utf-8"}
      }).success(function(data) {
        $scope.message = data.message
      });
		}
	});

	distApp.controller('countriesShowController', function($scope, $http, $routeParams) {
		// create a message to display in our view
		$http.get(domain_path+"/countries/"+$routeParams.country_id).success(function(data) {
  		$scope.states = data.states
  		$scope.country = data.country
  	})
	});
	distApp.controller('statesNewController', function($scope, $http) {
		$scope.state = {}
		$scope.submitForm = function(){
			$http({
				url: domain_path+"/states", 
				method: "post",
				data: $scope.state,
				transformRequest: form_transformation,
        headers: {'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8',"Accept":"application/x-www-form-urlencoded;charset=utf-8"}
      }).success(function(data) {
        $scope.message = data.message
      });
		}
	});

	distApp.controller('statesShowController', function($scope, $http, $routeParams) {
		// create a message to display in our view
		$http.get(domain_path+"/states/"+$routeParams.state_id).success(function(data) {
  		$scope.cities = data.cities
  		$scope.state = data.state
  		$scope.country = data.country
  	})
	});
	distApp.controller('citiesNewController', function($scope, $http) {
		$scope.city = {}
		$scope.submitForm = function(){
			$http({
				url: domain_path+"/cities", 
				method: "post",
				data: $scope.city,
				transformRequest: form_transformation,
        headers: {'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8',"Accept":"application/x-www-form-urlencoded;charset=utf-8"}
      }).success(function(data) {
        $scope.message = data.message
      });
		}
	});
	distApp.controller('distributorsIndexController', function($scope, $http) {
		$http.get(domain_path+"/distributors").success(function(data) {
  		$scope.distributors = data.distrubutors
  	})
	});
	distApp.controller('distributorsNewController', function($scope, $http) {
		$http.get(domain_path+"/countries").success(function(data) {
  		$scope.countries= data.countries
  	})
  	$http.get(domain_path+"/states").success(function(data) {
  		$scope.states= data.states
  	})
  	$http.get(domain_path+"/cities").success(function(data) {
  		$scope.cities= data.cities
  	})
		$scope.distributor = {}
		$scope.countryChange = function(){
			$scope.exstates = []
			angular.forEach($scope.distributor.included_countries,function(val){
				$http({
					url: domain_path+"/countries/"+val, 
					method: "get"
	      }).success(function(data) {
	        $scope.exstates = $scope.exstates.concat(data.states)
	      });
				
			})
		}
		$scope.stateChange = function(){
			$scope.excities = []
			angular.forEach($scope.distributor.included_states,function(val){
				$http({
					url: domain_path+"/states/"+val, 
					method: "get"
	      }).success(function(data) {
	        $scope.excities = $scope.excities.concat(data.cities)
	      });
				
			})
		}
		$scope.submitForm = function(){
			$http({
				url: domain_path+"/distributors", 
				method: "post",
				data: $scope.distributor,
				transformRequest: form_transformation,
        headers: {'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8',"Accept":"application/x-www-form-urlencoded;charset=utf-8"}
      }).success(function(data) {
        $scope.message = data.message
      });
		}
	});
