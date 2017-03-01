var router = angular.module('materialApp.routes', ['ui.router']);
router.config(function($stateProvider, $urlRouterProvider, $locationProvider) {

    $urlRouterProvider.otherwise('/app/');

    // UI Router States
    // Inserting Page title as State Param
    $stateProvider
        .state('home', {
            url: '/app/',
            templateUrl: 'home.html',
			controller: 'appCtrl',
			controllerAs: 'App',

            params: {
                title: "Distributors Portal"
            }
        })
        .state('admin', {
            url: '/app/admin',
            templateUrl: '/modules/admin/views/admin.html',
            controller: 'adminCtrl',
            controllerAs: 'Admin',
            params: {
                title: "Admin Portal"
            }
        })
		.state('editDistributor', {
            url: '/app/editDistributor/:id',
            templateUrl: '/modules/editDistributor/views/editDistributor.html',
            controller: 'editDistributorCtrl',
            controllerAs: 'EditDistributor',
            params: {
                title: "Admin Portal"
            }
        })
		.state('distributor', {
            url: '/app/distributor/:id',
            templateUrl: '/modules/distributor/views/distributor.html',
            controller: 'distributorCtrl',
            controllerAs: 'Distributor',
            params: {
                title: "Distributor Portal"
            }
        })



    $locationProvider.html5Mode(true);

});
