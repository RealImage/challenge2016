(function() {
    'use strict';

    moduleApp.config(function($stateProvider,$urlRouterProvider){
        $urlRouterProvider.otherwise("/Home");

        $stateProvider
        .state('home', {
          url: "/Home",
          templateUrl: "home.html",
          controller:"homeCtrl"
        })
        .state('check', {
          url: "/StatusCheck",
          templateUrl: "checkStatus.html",
          controller:"statusCtrl"
        })
        .state('create', {
          url: "/CreateDistributor",
          templateUrl: "createDistributor.html",
          controller:"createCtrl"
        });
    });
})();

