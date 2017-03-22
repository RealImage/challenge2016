app.controller('checkCtrl',['$scope','utilityService','$timeout',function($scope,utilityService,$timeout){
  $scope.countryArray=[];
  $scope.stateArray=[];
  $scope.cityArray=[];
  $scope.callMe=function(){
    $scope.dataArray=utilityService.getRegionArray();
  $scope.sharedArray=utilityService.getRegionObject();
  console.log( $scope.dataArray);
  console.log( $scope.sharedArray);
  }
  $scope.callMe();  
  
}]);