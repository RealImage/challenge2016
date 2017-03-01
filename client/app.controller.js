angular.module('appCtrl', ['adminService'])
.controller('appCtrl', function(Admin, $stateParams, $rootScope) {


    self = this;

	self.updateTitle = function() {
        $rootScope.title = $stateParams.title;
    }

    // Run updateTitle on each state change
    $rootScope.$on('$stateChangeSuccess', self.updateTitle);

	Admin.getAllDistributors().then(function (result) {
		if(result.data.status){
			self.distributors = result.data.data.filter(function (distributor) {
				return distributor.name;
			})
		}
	})
})
