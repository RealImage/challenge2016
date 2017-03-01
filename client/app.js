var materialApp = angular
.module('materialApp', [
    'materialApp.routes',
    'ui.router',
    'ngMaterial',
 	'ngMessages',
    'appCtrl',
    'adminCtrl',
    'adminService',
	'editDistributorCtrl',
	'distributorCtrl',

]).config(function($mdThemingProvider) {
  $mdThemingProvider.theme('default')
    .primaryPalette('red')
    .accentPalette('pink');
});
