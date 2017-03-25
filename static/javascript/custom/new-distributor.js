myApp.controller('AppCtrl', ['$scope', function($scope) {
  var allPlaces = vm.AllCities;
  //console.log(vm);
  $scope.countries = vm.UniqueCountries;
  var uniqueProvinces = vm.UniqueProvinces

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

  $scope.selectAllCountries = function() {
    if (allCountryFlag == false) {
      $scope.selectedCountries = [];
      $scope.selectedCountries = vm.UniqueCountries;
      allCountryFlag = true;
    } else {
      $scope.selectedCountries = [];
      allCountryFlag = false;
    }
  }

  $scope.countryChecked = function(country) {
    var index = $scope.selectedCountries.indexOf(country);
    if(index === -1) {
      $scope.selectedCountries.push(country);
    } else {
      $scope.selectedCountries.splice(index, 1);
    }
  }

  $scope.gotoProvince = function() {
    $scope.countryButton = false;
    $scope.countrySelection = false;
    $scope.provinceSelection = true;
    for (var i = 0; i < $scope.selectedCountries.length; i++) {
      for (var j = 0; j < uniqueProvinces.length; j++) {
        if (uniqueProvinces[j][1] == $scope.selectedCountries[i]) {
          $scope.provincesByCountry.push(uniqueProvinces[j][0]);
        }
      }
    }
  }

  $scope.selectAllProvinces = function() {
    if (allProvinceFlag == false) {
      $scope.selectedProvinces = [];
      $scope.selectedProvinces = $scope.provincesByCountry;
      allProvinceFlag = true;
    } else {
      $scope.selectedProvinces = [];
      allProvinceFlag = false;
    }
  }

  $scope.provinceChecked = function(province) {
    var index = $scope.selectedProvinces.indexOf(province);
    if(index === -1) {
      $scope.selectedProvinces.push(province);
    } else {
      $scope.selectedProvinces.splice(index, 1);
    }
  }


  $scope.gotoCity = function() {
    $scope.provinceButton = false;
    $scope.provinceSelection = false;
    $scope.citySelection = true;
    for (var i = 0; i < $scope.selectedProvinces.length; i++) {
      for (var j = 0; j < allPlaces.length; j++) {
        if (allPlaces[j][4] == $scope.selectedProvinces[i]) {
          $scope.citiesByProvince.push(allPlaces[j][3]);
        }
      }
    }
  }

  $scope.selectAllCities = function() {
    if (allCityFlag == false) {
      $scope.selectedCities = [];
      $scope.selectedCities = $scope.citiesByProvince;
      allCityFlag = true;
    } else {
      $scope.selectedCities = [];
      allCityFlag = false;
    }
  }

  $scope.cityChecked = function(city) {
    var index = $scope.selectedCities.indexOf(city);
    if(index === -1) {
      $scope.selectedCities.push(city);
    } else {
      $scope.selectedCities.splice(index, 1);
    }
    console.log($scope.selectedCities);
  }

}]);
