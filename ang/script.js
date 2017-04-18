	
	var distApp = angular.module('distApp', ['ngRoute']);
  var domain_path = "http://localhost:3000"
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
		.when('/countries/create', {
				templateUrl : 'pages/countries-new.html',
				controller  : 'countriesCreateController'
			})
		.when('/countries/:country_id', {
				templateUrl : 'pages/country-show.html',
				controller  : 'countriesShowController'
			})
		.when('/states/:state_id', {
				templateUrl : 'pages/state-show.html',
				controller  : 'statesShowController'
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
				data: {name:$scope.country.name, code:$scope.country.code} ,
				transformRequest: function(obj) {
	        var str = [];
	        for(var p in obj)
	        str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
	        return str.join("&");
	    	},
        headers: {'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8',"Accept":"application/x-www-form-urlencoded;charset=utf-8"}
      }).success(function(data) {
        $scope.message = data.message
      });
		}
	});
	distApp.controller('countriesCreateController', function($scope, $http) {
		console.log($scope)
	});
	distApp.controller('countriesShowController', function($scope, $http, $routeParams) {
		// create a message to display in our view
		$http.get(domain_path+"/countries/"+$routeParams.country_id).success(function(data) {
  		$scope.states = data.states
  		$scope.country = data.country
  	})
	});

	distApp.controller('statesShowController', function($scope, $http, $routeParams) {
		// create a message to display in our view
		$http.get(domain_path+"/states/"+$routeParams.state_id).success(function(data) {
  		$scope.cities = data.cities
  		$scope.state = data.state
  		$scope.country = data.country
  	})
	});
