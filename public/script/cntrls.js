(function() {
    'use strict';
    
    moduleApp.controller('homeCtrl',['$scope','$state', '$http', 'myService', function($scope,$state,$http,myService){
        $scope.dNameList = 'Select';
        $http({ method: 'GET',
            url: '/Home'
         }).then(function (response){
            $scope.dNames = response.data;
         }, function (error) {
            console.log('No Distribution List found in Server!!!');
        });
        
        $scope.renderCheck=function(){
            myService.fileName = $scope.dNameList; 
            $state.go('check');
        }
	}]);
	
	moduleApp.controller('statusCtrl',['$scope', '$state', '$http', 'myService', function($scope, $state, $http, myService){
		$scope.status = 'Welcome !!!';
        $scope.fileName = myService.getFileName();
        if($scope.fileName){
            var req = {'parent' : $scope.fileName}
            $http.post('/getChild', req)
           .then( 
               function(response){
                 if(response.data){
                    $scope.myDistribution = response.data.split(',');
                    $scope.isChildAvailable = true;
                 }
               }, 
               function(err){
                 console.log('Error in Fetching Child Distributions');
               }
            ); 
        }
        $scope.submit = function(){
            $scope.dList = null;
            if($scope.country){
                $scope.dList = {'country':$scope.country}
                if($scope.province){
                    $scope.dList['province'] = $scope.province;
                    if($scope.city){
                        $scope.dList['city'] = $scope.city;  
                    }
                }
                if($scope.fileName){
                     $scope.dList['fileName'] = $scope.fileName;
                }
                $scope.sendDetails($scope.dList);
            }else{
                $scope.status = 'Atleast country of distribution must be mentioned';
            }
        }
        $scope.sendDetails = function(dList){
            $http.post('/getStatus', dList)
           .then( 
               function(response){
                 $scope.status = response.data;
               }, 
               function(err){
                 console.log('Sorry!!! error on /getStatus' + err);
               }
            ); 
        }
        $scope.createDistributor = function(){
             $state.go('create');
        }
        $scope.Edit = function(selectedChild){
           myService.editFile = selectedChild;
           $state.go('create');
        }
        $scope.openTab = function(){
            return myService.openCSV();
        }
	}]);
    
    moduleApp.controller('createCtrl',['$scope','$state', '$http', 'myService', function($scope,$state,$http,myService){
        $scope.status = 'Welcome !!!';
        $scope.countryList = 'Select';
        $scope.provinceList = 'Select';
        $scope.cityList = 'Select';
        $scope.addDistributor = [];
        var fileName = myService.getFileName(); ;
        $scope.fileToEdit = myService.getFileToEdit();
        if($scope.fileToEdit){
            myService.editFile='';
            $scope.newDistribution = $scope.fileToEdit;
            var req ={'childDistributor' : $scope.fileToEdit}
            $http.post('/EditChild', req)
             .then(function (response){
                 $scope.addList(response.data);
             }, function (error) {
                console.log('Error on /EditChild!!!');
            });
        }
        var req ={'nameOfFile' : fileName}
        if(fileName){
             $http.post('/getCountry', req)
             .then(function (response){
                $scope.countries = response.data;
             }, function (error) {
                console.log('Error on fetching country!!!');
            });
        }   
        $scope.onCountrySelect = function(){
            var req ={'countryList' : $scope.countryList,'nameOfFile' : fileName}
            $http.post('/getProvince', req)
             .then(function (response){
                $scope.provinces = response.data;
             }, function (error) {
                console.log('Error on fetching province!!!');
            });
        }
        $scope.onProvinceSelect = function(){
            var req ={'provinceList' : $scope.provinceList,'nameOfFile' : fileName}
            $http.post('/getCity', req)
             .then(function (response){
                $scope.cities = response.data;
             }, function (error) {
                console.log('Error on fetching cities!!!');
            });
        }
        $scope.add = function(){
            if($scope.countryList){
                var dList ={
                    'country' : $scope.countryList,
                    'province' : $scope.provinceList,
                    'city' : $scope.cityList,
                    'addNewFile' : true,
                    'fileName' : fileName 
                };
                $http.post('/getStatus', dList)
               .then( 
                   function(response){
                     $scope.addList(response.data);
                   }, 
                   function(response){
                     console.log('Error on ADD!' + response);
                   }
                );
            }else{
                $scope.status = 'Atleast country of distribution must be selected';
            }
        }
        $scope.addList = function(data){
            if(typeof(data)!='string'){
                $scope.dataAdded = true;
                data.forEach(function(item){
                    $scope.addDistributor.push(item);
                })
                $scope.addDistributor = $scope.getuniqueResults($scope.addDistributor,'City Name');
                $scope.status = 'Welcome!!!';
            }else{
                $scope.status = 'No data has been added';
            }
            $scope.clear('clean');
        }
        $scope.clear = function(field){
            if(field == 'clean'){
                $scope.countryList = 'Select';
                $scope.provinces = '';
                $scope.cities = '';                 
            }
            if(field == 'country')
                $scope.provinceList = 'Select';
            $scope.cityList = 'Select';
        }
        $scope.remove = function(){
            if($scope.countryList && $scope.countryList!='Select'){
                if($scope.provinceList && $scope.provinceList!='Select'){
                    if($scope.cityList && $scope.cityList!='Select'){
                        $scope.removeList('City Name',$scope.cityList);
                    }else{
                        $scope.removeList('Province Name',$scope.provinceList);
                    }
                }else{
                    $scope.removeList('Country Name',$scope.countryList);
                }
            }else{
                $scope.status = 'Atleast country of distribution must be selected';
            }
        }
        $scope.removeList = function(key, value){
            $scope.addDistributor=$scope.addDistributor.filter((elem) => {
                return elem[key] != value;
            })
            $scope.clear('clean');
        }
        $scope.submit = function(){
            if($scope.newDistribution){
                $scope.disableText = true;
                if($scope.addDistributor && $scope.addDistributor.length>0){
                    var req = {
                        'addDistributorList' : $scope.addDistributor,
                        'newDistribution' : $scope.newDistribution,
                        'parentDistributor' : fileName,
                        'editMode' : $scope.fileToEdit
                    };
                    $http.post('/newDistributer', req)
                     .then(function (response){
                        $scope.status = response.data;
                     }, function (error) {
                        $scope.status = 'Adding new distributor failed!!!';
                    });
                }else{
                    $scope.status = 'No data has been selected';
                }
            }else{
                $scope.status = 'Please ensure to mention name of the distributor';
            }
        }
        $scope.back =function(){
           $state.go('check'); 
        }
        $scope.getuniqueResults = function(arr, param){
            return arr.filter(function(item, pos, array){
                return array.map(function(mapItem){ 
                    return mapItem[param]; 
                }).indexOf(item[param]) === pos;
            })
        }
        $scope.openTab = function(){
            return myService.openCSV();
        }
	}]);
    
})();