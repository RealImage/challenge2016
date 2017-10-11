	
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
				controller  : 'subDistributorsNewController'
			})
		.when('/distributors/permision/:distributor_id', {
				templateUrl : 'pages/distributors-permision.html',
				controller  : 'distributorsPermisionController'
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
		$http.get(domain_path+"/states/"+$routeParams.state_id)(+function(data) {
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
	distApp.controller('subDistributorsNewController', function($scope, $http, $routeParams) {
		$scope.distributor = {distributor_id: $routeParams.distributor_id}
		$scope.states = []
		$scope.cities = []
		$scope.exstates = []
		$scope.excities = []
  	$http.get(domain_path+"/distributors/"+$routeParams.distributor_id).success(function(data) {
  		$scope.countries = data.in_countries
			var in_states = data.in_states
			var in_cities = data.in_cities
			var excluded_states = data.ex_states
			var excluded_cities = data.ex_cities

			angular.forEach($scope.countries,function(country){
				$http({
					url: domain_path+"/countries/"+country.Id, 
					method: "get"
	      }).success(function(data) {
	        $scope.exstates = $scope.exstates.concat(data.states)
	        angular.forEach(data.states,function(state){
	        	var is_exc_state = false
	        	angular.forEach(excluded_states,function(ex_state){
		        	if(ex_state.Id == state.Id){
		        		is_exc_state = true
		        	}
		        });
		        if(!is_exc_state){
		        	$scope.states.push(state)
		        	$http({
								url: domain_path+"/states/"+state.Id, 
								method: "get"
				      }).success(function(data) {
	       				$scope.excities = $scope.excities.concat(data.cities)
	       				angular.forEach(data.cities,function(city){
		       				var is_exc_city = false
		       				angular.forEach(excluded_cities,function(ex_city){
					        	if(ex_city.Id == city.Id){
					        		is_exc_state = true
					        	}
					        });
					        if(!is_exc_city){
					        	$scope.cities.push(city)
					        }
		       			});
				      });
		        }
	        });
	      });
			})

  	})
		
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
	distApp.controller('distributorsPermisionController', function($scope, $http, $routeParams) {
				$scope.distributor = {distributor_id: $routeParams.distributor_id}

		$http.get(domain_path+"/countries").success(function(data) {
  		$scope.countries= data.countries
  	})
  	$scope.countryChange = function(){
  		$scope.cities = []
			$http({
				url: domain_path+"/countries/"+$scope.distributor.country_ids, 
				method: "get"
      }).success(function(data) {
        $scope.states = data.states
      });
		}
		$scope.stateChange = function(){
			$http({
				url: domain_path+"/states/"+$scope.distributor.state_ids, 
				method: "get"
      }).success(function(data) {
        $scope.cities = data.cities
      });
		}
		$scope.submitForm = function(){
			$http({
				url: domain_path+"/permisions", 
				method: "post",
				data: $scope.distributor,
				transformRequest: form_transformation,
        headers: {'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8',"Accept":"application/x-www-form-urlencoded;charset=utf-8"}
      }).success(function(data) {
        $scope.message = data.message
      });
		}
	});

