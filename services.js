app.factory('utilityService',function($http,$q){
  var region = {};
  var data=[];
  return{
    getRegionObject:function(){
      return region;
    },
    setRegionObject:function(reg){
      region=reg;
    },
    getRegionArray:function(){
      return data;
    },
    setRegionArray:function(dt){
      data=dt;
    },
    getData:function(url){
       var deferred = $q.defer();
      $http.get(url,{headers: {'Content-type': 'application/json'}})
        .success(function (data) {
            deferred.resolve(data);
        })
        .error(function (error) {
            deferred.reject(error);
        });

      return deferred.promise;
    }
  }
  
});