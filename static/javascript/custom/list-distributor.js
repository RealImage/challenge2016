myApp.controller('AppCtrl', ['$scope', function($scope) {
  if(vm.List == null){
    $scope.list = [];
  } else {
    $scope.list = vm.List;
  }

  $scope.gotoView = function(item) {
    window.location = "/view?name=" + item;
    return false;
  }

}]);
