myApp.controller('AppCtrl', ['$scope', function($scope) {
  var allPlaces = vm.AllCities;
  $scope.countries = vm.UniqueCountries;
  var uniqueProvinces = vm.UniqueProvinces;
  $scope.existingDistributors = [];

  var detailsSaveLevel = 0;

  /*Condition for checking whether there are any already created distributors*/
  if (vm.DistributorCities != null) {
    $scope.existingDistributors = Object.keys(vm.DistributorCities);
  }

  $scope.superDistributor = "";
  $scope.name = "";

  var allCountryFlag = false;
  var allProvinceFlag = false;
  var allCityFlag = false;

  $scope.countrySelection = true;
  $scope.provinceSelection = false;
  $scope.citySelection = false;

  $scope.selectedCountries = [];
  $scope.provincesByCountry = [];
  $scope.citiesByProvince = [];
  $scope.selectedProvinces = [];
  $scope.selectedCities = [];

  /*Function for selecting and deselecting a country*/
  $scope.countryChecked = function(country) {
    var index = $scope.selectedCountries.indexOf(country);
    if(index === -1) {
      $scope.selectedCountries.push(country);
    } else {
      $scope.selectedCountries.splice(index, 1);
    }
  }

  /*Function for saving from country level*/
  $scope.saveAllCountryDetails = function() {

  }

  /*Function for creating and displaying province array*/
  $scope.gotoProvince = function() {
    detailsSaveLevel = 1;
    $scope.countryButton = false;
    $scope.countrySelection = false;
    $scope.provinceSelection = true;
    if($scope.superDistributor == "") {
      for (var i = 0; i < $scope.selectedCountries.length; i++) {
        for (var j = 0; j < uniqueProvinces.length; j++) {
          if (uniqueProvinces[j][1] == $scope.selectedCountries[i]) {
            $scope.provincesByCountry.push(uniqueProvinces[j][0]);
          }
        }
      }
    } else {
      var distributorPlaces = vm.DistributorCities[$scope.superDistributor];
      var tempArray = [];
      for (var i = 0; i < distributorPlaces.length; i++) {
        for (var j = 0; j < $scope.selectedCountries.length; j++) {
          if(distributorPlaces[i][5] == $scope.selectedCountries[j]) {
            var index = tempArray.indexOf(distributorPlaces[i][4]);
            if(index == -1) {
              tempArray.push(distributorPlaces[i][4]);
            }
          }
        }
      }
      $scope.provincesByCountry = tempArray;
    }
  }

  /*Function for selecting and deselecting provinces*/
  $scope.provinceChecked = function(province) {
    var index = $scope.selectedProvinces.indexOf(province);
    if(index === -1) {
      $scope.selectedProvinces.push(province);
    } else {
      $scope.selectedProvinces.splice(index, 1);
    }
  }

  /*Function for creating and displaying city array*/
  $scope.gotoCity = function() {
    detailsSaveLevel = 2;
    $scope.provinceButton = false;
    $scope.provinceSelection = false;
    $scope.citySelection = true;
    if($scope.superDistributor == "") {
      for (var i = 0; i < $scope.selectedProvinces.length; i++) {
        for (var j = 0; j < allPlaces.length; j++) {
          if (allPlaces[j][4] == $scope.selectedProvinces[i]) {
            var index = $scope.citiesByProvince.indexOf(allPlaces[j][3]);
            if(index == -1){
              $scope.citiesByProvince.push(allPlaces[j][3]);
            }
          }
        }
      }
    } else {
      var distributorPlaces = vm.DistributorCities[$scope.superDistributor];
      var tempArray = [];
      for (var i = 0; i < distributorPlaces.length; i++) {
        for (var j = 0; j < $scope.selectedProvinces.length; j++) {
          if(distributorPlaces[i][4] == $scope.selectedProvinces[j]) {
            var index = tempArray.indexOf(distributorPlaces[i][3]);
            if(index == -1) {
              tempArray.push(distributorPlaces[i][3]);
            }
          }
        }
      }
      $scope.citiesByProvince = tempArray;
    }
  }

  /*Function for selecting and deselecting cities*/
  $scope.cityChecked = function(city) {
    var index = $scope.selectedCities.indexOf(city);
    if(index === -1) {
      $scope.selectedCities.push(city);
    } else {
      $scope.selectedCities.splice(index, 1);
    }
  }

  /*Function for collecting filled data and save them*/
  $scope.saveDetails = function() {
    var citiesToServer = "";
    if(detailsSaveLevel == 0) {
      for (var i = 0; i < $scope.selectedCountries.length; i++) {
        citiesToServer = citiesToServer + "&selectedCities=" + $scope.selectedCountries[i];
      }
    } else if (detailsSaveLevel == 1) {
      for (var i = 0; i < $scope.selectedProvinces.length; i++) {
        citiesToServer = citiesToServer + "&selectedCities=" + $scope.selectedProvinces[i];
      }
    } else {
      for (var i = 0; i < $scope.selectedCities.length; i++) {
        citiesToServer = citiesToServer + "&selectedCities=" + $scope.selectedCities[i];
      }
    }

    $.ajax({
      url: '/new',
      type: 'POST',
      dataType: 'json',
      data : "name=" + $scope.name + "&mode=" + detailsSaveLevel + citiesToServer,
      success : function(data) {
        window.location = "/";
      }
    });
  }

  /*Function for selecting existing distributors and adding them as super distributors fro newly creating distributors*/
  $scope.selectDistributor = function(){
    if($scope.superDistributor != ""){
      var distributorPlaces = vm.DistributorCities[$scope.superDistributor];
      $scope.countries = [];
      for (var i = 0; i < distributorPlaces.length; i++) {
        var index = $scope.countries.indexOf(distributorPlaces[i][5]);
        if(index == -1) {
          $scope.countries.push(distributorPlaces[i][5]);
        }
      }
    }

  }

}]);
