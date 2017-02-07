var realImageApp = angular.module('realImgApp', ['ui.bootstrap', 'ui']);
realImageApp.controller('realImgAppCtrl', ['$scope', '$http', '$rootScope', '$modal', function ($scope, $http, $rootScope, $modal) {
		$rootScope.distributors = [];
		$rootScope.territories = {};

		$scope.onLoad = function() {
			$http({
					method : 'GET',
					url : '/territories'
			}).then(function(response) {
				$rootScope.territories = response.data.result;
				$scope.getAllDistributors();
			});
		};

		$scope.getAllDistributors = function() {
			$http({
					method : 'GET',
					url : '/distributors'
			}).then(function(response) {
				$rootScope.distributors = response.data.result;
				$scope.distributors = response.data.result;
			});
		};

		$scope.addDistributor = function() {
			$rootScope =
			$modal.open({
				backdrop : 'static',
				keyboard : false,
				backdropClick : false,
				templateUrl : 'distributorModal.html',
				controller : 'distributorModalCtrl',
				resolve : {
					object : function() {
						return {
							callback : $scope.onLoad
						};
					}
				}
			});
		};

		$scope.checkPermission = function() {
			$modal.open({
				backdrop : 'static',
				keyboard : false,
				backdropClick : false,
				templateUrl : 'checkPermissionModal.html',
				controller : 'checkPermissionModalCtrl'
			});
		};

		$scope.openDistributor = function(item) {
			$modal.open({
				backdrop : 'static',
				keyboard : false,
				backdropClick : false,
				templateUrl : 'distributorModal.html',
				controller : 'distributorModalCtrl',
				resolve : {
					object : function() {
						return {
							id : item.id,
							callback : $scope.onLoad
						};
					}
				}
			});
		};

		$scope.onLoad();
	}
]);

realImageApp.controller('distributorModalCtrl', ['$scope', '$http', '$rootScope', '$modalInstance', 'object', function ($scope, $http, $rootScope, $modalInstance, object) {
	$scope.resource = {
		rules : []
	};
	$scope.distributorLocation = "";
	$scope.countries =  $rootScope.territories["countries"];
	$scope.provinces = {};
	$scope.cities = {};

	$scope.onLoad = function() {
		if(object.id > 0) {
			$http({
					method : 'GET',
					url : '/distributors/' + object.id
			}).then(function(response) {
				if(response.data.message == "success") {
					$scope.resource = response.data.result;
					for(var i=0;i<$scope.resource.rules.length;i++) {
						$scope.countryOnChange($scope.resource.rules[i]);
						$scope.provinceOnChange($scope.resource.rules[i]);
					}
				} else
					alert(response.data.result.join(","));
			});
		}
	};

	$scope.distributorOnChange = function() {
		$scope.provinces = {};
		$scope.cities = {};
		$scope.countries = {};
		if($scope.resource.superDistributorId > 0) {
			for(var i=0;i<$rootScope.distributors.length;i++) {
				if($rootScope.distributors[i].id == $scope.resource.superDistributorId) {
					$scope.resource.superDistributor = $rootScope.distributors[i].distributorName;
					$scope.distributorLocation = $rootScope.distributors[i].territories;
					$scope.countries = $rootScope.distributors[i].territories["countries"];
					$scope.cities = $rootScope.distributors[i].territories["cities"];
					$scope.provinces = $rootScope.distributors[i].territories["provinces"];
					break;
				}
			}
		} else {
			$scope.distributorLocation = "";
			$scope.countries = $rootScope.territories["countries"];
		}
	};

	$scope.countryOnChange = function(item) {
		$scope.provinces = {};
		$scope.cities = {};
		provinces = ($scope.distributorLocation && Object.keys($scope.distributorLocation["provinces"]).length > 0) ? $scope.distributorLocation["provinces"] : $rootScope.territories["provinces"];
		for(var prop in provinces)
			if(prop.indexOf(item.countryCode) == 0)
				$scope.provinces[prop] = provinces[prop];
	};

	$scope.provinceOnChange = function(item) {
		$scope.cities = {};
		var cities = ($scope.distributorLocation && Object.keys($scope.distributorLocation["cities"]).length > 0) ? $scope.distributorLocation["cities"] : $rootScope.territories["cities"];
		for(var prop in cities)
			if(prop.indexOf(item.provinceCode) == 0)
				$scope.cities[prop] = cities[prop];
	};

	$scope.distributorOperation = function(actionVerb) {
		$scope.hide = true;
		$http({
			method : "POST",
			url : "/distributors",
			data : {
				actionVerb : actionVerb,
				data : $scope.resource
			}
		}).then(function(response) {
			if(response.data.message == "success") {
					object.callback();
					$modalInstance.dismiss('cancel');
			} else {
				alert(response.data.result.join('\n'));
				$scope.hide = false;
			}
		});
	};

	$scope.deleteRule = function(index) {
		$scope.resource.rules.splice(index,1);
	};

	$scope.addRule = function() {
		$scope.resource.rules.push({});
	};

	$scope.close = function() {
		$modalInstance.dismiss('cancel');
	};

	$scope.onLoad();
}]);

realImageApp.controller('checkPermissionModalCtrl', ['$scope', '$http', '$modalInstance', '$rootScope', function ($scope, $http, $modalInstance, $rootScope) {
	$scope.resource = {};

	$scope.onLoad = function () {
		$scope.countries =  $rootScope.territories["countries"];
		$scope.provinces = $rootScope.territories["provinces"];
		$scope.cities = $rootScope.territories["cities"];
	};

	$scope.countryOnChange = function() {
		$scope.provinces = {};
		$scope.cities = {};
		var provinces = $rootScope.territories["provinces"];
		for(var prop in provinces)
			if(prop.indexOf($scope.resource.countryCode) == 0)
				$scope.provinces[prop] = provinces[prop];
	};

	$scope.provinceOnChange = function() {
		$scope.cities = {};
		var cities = $rootScope.territories["cities"];
		for(var prop in cities)
			if(prop.indexOf($scope.resource.provinceCode) == 0)
				$scope.cities[prop] = cities[prop];
	};

	$scope.checkPermission = function() {
		$scope.hide = true;
		$http({
			method : "POST",
			url : "/checkPermission",
			data : {
				data : $scope.resource
			}
		}).then(function(response) {pp=response;
			if(response.data.message == "success") {
				if(response.data.result.length > 0)
					alert("NO \n Distributor don't have permission for this location\n Reasons : " + response.data.result.join('\n'));
				else
					alert("YES \n Distributor have permission for this location");
			} else
				alert(response.data.result.join('\n'));

			$scope.hide = false;
		});
	};

	$scope.close = function() {
		$modalInstance.dismiss('cancel');
	};

	$scope.onLoad();
}]);
