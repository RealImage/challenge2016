app.controller('permissionsCtrl',['$scope','utilityService',function($scope,utilityService){
  
  $scope.dataSet=[];
   $scope.selectedStateArray=null;
    $scope.selectedCityArray=null;
    $scope.selectedCountryArray=null;
    $scope.operationType="include";
    $scope.includeButton=true;
    $scope.excludeButton=false;
    $scope.operationChange=function(){
      if($scope.operationType=="include"){
         $scope.includeButton=true;
         $scope.excludeButton=false;
         $scope.setupInclude();
      }else{
         $scope.includeButton=false;
         $scope.excludeButton=true;
         $scope.setupExclude();
      }
      $scope.selectedStateArray=null;
    $scope.selectedCityArray=null;
    $scope.selectedCountryArray=null;
    }
    $scope.setupInclude=function(){
      $scope.finalDataSet=$scope.finalDataSetCopy;
    };
    $scope.setupExclude=function(){
      $scope.finalDataSetCopy=angular.copy($scope.finalDataSet);
      $scope.excludeArray=utilityService.getRegionArray();
      for(let i=0;i<$scope.excludeArray.length;i++){
        if($scope.excludeArray[i].name==$scope.distributorName){
          $scope.finalDataSet=$scope.excludeArray[i].include;
        }
      }
    }
  utilityService.getData('data.json').then(function(data){
    var myArray=[];
    myArray=data;
    $scope.countryArray=[];
    $scope.stateArray=[];
    $scope.cityArray=[];
    for(let i=0;i<myArray.length;i++){
      var objCountry={};
      var objState={};
      var objCity={};
      objCountry.countryCode=myArray[i].countryCode;
      objCountry.countryName=myArray[i].countryName;
      objState.provinceCode=myArray[i].provinceCode;
      objState.provinceName=myArray[i].provinceName;
      objState.countryCode=myArray[i].countryCode;
      objCity.cityCode=myArray[i].cityCode;
      objCity.cityName=myArray[i].cityName;
      objCity.provinceCode=myArray[i].provinceCode;
      $scope.countryArray.push(objCountry);
      $scope.stateArray.push(objState);
      $scope.cityArray.push(objCity);
    }
    
    function removeDuplicate(arr,prop){
        var newArr = [];
        angular.forEach(arr, function(value, key) {
            var exists = false;
            angular.forEach(newArr, function(val2, key) {
                if(angular.equals(value[prop], val2[prop])){ exists = true }
            });
            if(exists === false && value[prop] !== "") { newArr.push(value); }
        });
        return newArr;
    }
    $scope.finalCountryArray=removeDuplicate($scope.countryArray,'countryCode');
    $scope.finalStateArray=removeDuplicate($scope.stateArray,'provinceCode');
    $scope.finalCityArray=removeDuplicate($scope.cityArray,'cityCode');
    $scope.shareArray={};
     $scope.shareArray.countryArray=$scope.finalCountryArray;
     $scope.shareArray.stateArray=$scope.finalStateArray;
     $scope.shareArray.cityArray=$scope.finalCityArray;
    utilityService.setRegionObject($scope.shareArray);
     console.log($scope.shareArray);
     $scope.finalDataSet=[];
    console.log("HIII");
    console.log($scope.finalCountryArray);
   
    angular.forEach($scope.finalCountryArray,function(x){
       var countryObj={};
      countryObj.countryCode=x.countryCode;
      countryObj.countryName=x.countryName;
      countryObj.provinces=[];
      angular.forEach($scope.finalStateArray,function(y){
        var stateObj={};
        stateObj.provinceName=y.provinceName;
        stateObj.provinceCode=y.provinceCode;
        stateObj.cities=[];
        angular.forEach($scope.finalCityArray,function(z){
          if(y.provinceCode==z.provinceCode){
            stateObj.cities.push(z);
          }
        });
        if(x.countryCode==y.countryCode){
        countryObj.provinces.push(stateObj);}
      });
      $scope.finalDataSet.push(countryObj);
    });
    //console.log(angular.toJson($scope.finalDataSet));
    utilityService.setRegionArray($scope.finalDataSet);
    
  
  },function(){
    
  });
  $scope.changeCountry=function(){
  
    $scope.stateDisplay=[];
    angular.forEach($scope.selectedCountryArray,function(x){
      
      angular.forEach(x.provinces,function(y){
        y.countryCode=x.countryCode;
        $scope.stateDisplay.push(y);
      })
      
      console.log($scope.stateDisplay);
    });
    $scope.cityDisplay=[];
    
  }
  $scope.changeState=function(){
    
    $scope.cityDisplay=[];
    angular.forEach($scope.selectedStateArray,function(x){
      
      angular.forEach(x.cities,function(y){
        y.provinceCode=x.provinceCode;
        $scope.cityDisplay.push(y);
      })
      console.log($scope.cityDisplay);
    });
  }
  $scope.changeCity=function(){
    $scope.selectedCityArray=$scope.selectedCityArray;
  }
  
  $scope.distributorArray=[];
  $scope.distributorName=null;
  $scope.includeRegion=function(){
    
    $scope.selectedCountryArrayCopy=angular.copy($scope.selectedCountryArray);
    $scope.selectedStateArrayCopy=angular.copy($scope.selectedStateArray);
    $scope.selectedCityArrayCopy=angular.copy($scope.selectedCityArray);
    var distObj={};
    distObj.include=[];
   
    distObj.name=$scope.distributorName;
    angular.forEach($scope.selectedCountryArrayCopy,function(x){
       var countryObj={};
      countryObj.countryCode=x.countryCode;
      countryObj.countryName=x.countryName;
      countryObj.provinces=[];
      angular.forEach($scope.selectedStateArrayCopy,function(y){
        var stateObj={};
        stateObj.provinceName=y.provinceName;
        stateObj.provinceCode=y.provinceCode;
        stateObj.cities=[];
        angular.forEach($scope.selectedCityArrayCopy,function(z){
          if(y.provinceCode==z.provinceCode){
            stateObj.cities.push(z);
          }
        });
        if(x.countryCode==y.countryCode){
        countryObj.provinces.push(stateObj);}
      });
      distObj.include.push(countryObj);
    });
    $scope.distributorArray.push(distObj);
    utilityService.setRegionArray( $scope.distributorArray);
    console.log(angular.toJson(distObj));
  }
  
  
  
}]);

