(function(){
     	moduleApp.service('myService',['$window', function($window){
            this.fileName;
            this.editFile;
            this.getFileName = function(){
                if(this.fileName){
                    return this.fileName;
                }
            }
            this.getFileToEdit = function(){
                if(this.editFile){
                    return this.editFile;
                }
            }
            this.openCSV = function(){
                $window.open("https://raw.githubusercontent.com/RealImage/challenge2016/master/cities.csv");
            }
	}]);
})();