var routerApp = angular.module('routerApp', ['ui.router']);

routerApp.config(function($stateProvider, $urlRouterProvider) {
    
    $urlRouterProvider.otherwise('/home');
    
    $stateProvider
        
        .state('home', {

            url: '/home',

            templateUrl: 'views/home.html',

            controller : 'HomeController'
        });
    
});

routerApp.controller('HomeController', function($scope, $rootScope, $http){

	$scope.searchDetails = function(){

		var destributer = $scope.destributer;

		if(destributer == "" || destributer == undefined){

			alert('please select destributer');

			return;
		}

		var city = $scope.city;

		var province = $scope.province;

		var country = $scope.country;

		var searchObject = {};

		searchObject.destributer = destributer;

		if(city != undefined)searchObject.city = city;

		if(province != undefined)searchObject.province = province;

		if(country != undefined)searchObject.country = country;

		var url = "http://localhost:9090/api/checkAuthorization";

		$http.post(url, searchObject).then(

			function(successResponse){

				$scope.result = "This destributer is allow to access this region";

				$scope.infoclass = "bg-success";
			},

			function(failureResponse){

				$scope.result = "This destributer is not allow to access this region";

				$scope.infoclass = "bg-danger";
			}
		);
	}

	$scope.init = function(){

		$http.get('http://localhost:9090/api/getDistributersList').then(

			function(successResponse){

				$scope.destributers = successResponse.data.data;
			},

			function(failureResponse){

				alert('failed to get destributers list error');
			}
		);
	}	

	$scope.init();
});