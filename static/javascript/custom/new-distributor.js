myApp.controller('AppCtrl', ['$scope', function($scope) {
  var allPlaces = vm.AllCities;
  $scope.countries = vm.UniqueCountries;
  var uniqueProvinces = vm.UniqueProvinces;
  $scope.existingDistributors = [];
  var detailsSaveLevel = 0;

  $scope.nameError = false;
  $scope.saveValidate = "";

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
    if($scope.selectedCountries.length > 0){
      $scope.saveValidate = "true";
    } else {
      $scope.saveValidate = "";
    }
  }

    /*Function for creating and displaying province array*/
  $scope.gotoProvince = function() {
    document.getElementById("distributorSelect").disabled = true;
    $scope.saveValidate = "";
    detailsSaveLevel = 1;
    $scope.countryButton = false;
    $scope.countrySelection = false;
    $scope.provinceSelection = true;
    if($scope.superDistributor == "") {
      for (var i = 0; i < $scope.selectedCountries.length; i++) {
        for (var j = 0; j < uniqueProvinces.length; j++) {
          if (uniqueProvinces[j][0] == $scope.selectedCountries[i]) {
            $scope.provincesByCountry.push(uniqueProvinces[j]);
          }
        }
      }
    } else {

      var citiesToServer = "";

      for (var i = 0; i < $scope.selectedCountries.length; i++) {
        citiesToServer = citiesToServer + "&selectedCountries=" + $scope.selectedCountries[i];
      }
      var tempArray = [];
      $.ajax({
        url: '/new/distributor-provinces',
        type: 'POST',
        dataType: 'json',
        data : "&superDistributor=" + $scope.superDistributor + citiesToServer,
        success : function(data) {
          $scope.$apply(function(){
              $scope.provincesByCountry = data;
          });
        }
      });
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
    if($scope.selectedProvinces.length > 0){
      $scope.saveValidate = "true";
    } else {
      $scope.saveValidate = "";
    }
  }

  /*Function for creating and displaying city array*/
  $scope.gotoCity = function() {
    $scope.saveValidate = "";
    detailsSaveLevel = 2;
    $scope.provinceButton = false;
    $scope.provinceSelection = false;
    $scope.citySelection = true;
    if($scope.superDistributor == "") {
      for (var i = 0; i < $scope.selectedProvinces.length; i++) {
        for (var j = 0; j < allPlaces.length; j++) {
          if (allPlaces[j][1] == $scope.selectedProvinces[i][2] && allPlaces[j][2] == $scope.selectedProvinces[i][0]) {
            $scope.citiesByProvince.push(allPlaces[j]);
          }
        }
      }
    } else {

      console.log($scope.selectedProvinces);


      var citiesToServer = "";

      for (var i = 0; i < $scope.selectedProvinces.length; i++) {
        citiesToServer = citiesToServer + "&selectedProvinces=" + $scope.selectedProvinces[i][0] + "," + $scope.selectedProvinces[i][2];
      }
      var tempArray = [];
      $.ajax({
        url: '/new/distributor-cities',
        type: 'POST',
        dataType: 'json',
        data : "&superDistributor=" + $scope.superDistributor + citiesToServer,
        success : function(data) {
          $scope.$apply(function(){
              $scope.citiesByProvince = data;
          });
        }
      });
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
    if($scope.selectedCities.length > 0){
      $scope.saveValidate = "true";
    } else {
      $scope.saveValidate = "";
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
        citiesToServer = citiesToServer + "&selectedCities=" + $scope.selectedProvinces[i][2] + "," + $scope.selectedProvinces[i][0];
      }
    } else {
      for (var i = 0; i < $scope.selectedCities.length; i++) {
        citiesToServer = citiesToServer + "&selectedCities=" + $scope.selectedCities[i][0] + "," + $scope.selectedCities[i][1] + "," + $scope.selectedCities[i][2];
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
      $scope.countries = vm.DistributorCountries[$scope.superDistributor];
    }

  }

  $scope.validateName = function(){
    var index = vm.DistributorNames.indexOf($scope.name);
    if(index != -1) {
      $scope.nameError = true;
    } else {
      $scope.nameError = false;
    }
  }

}]);
