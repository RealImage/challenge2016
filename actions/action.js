/************************************************************************
 * Required Modules
 * **********************************************************************/

var _ = require('lodash');
require("underscore-query")(_)
var Promise = require('bluebird');
var data = require('../data.json');



/************************************************************************
 * Class Declaration
 * **********************************************************************/
var APIAction = function (app,conf) {
    this.app = app;
    this.conf = conf;

};

module.exports = APIAction;

APIAction.prototype.create= function(req,res){

var self=this
    var resObj={
    status:false,
        data:[],
        message:'FAILED'
    }
    req.query.include=req.query.include.map(function (x) {
        return x.toUpperCase()
    })
    req.query.exclude=req.query.exclude.map(function (x) {
        return x.toUpperCase()
    })
    if(self.conf.distributor && self.conf.distributor.length!=0 && req.query.parentDistributor!=(false || 'false') ){
       var exist=_.query(self.conf.distributor, {
            'distributorName': {$in : [req.query.distributorName]},
        });
        if(exist.length > 0){
           resObj.message='Distributor Already Exist'
            res.send(resObj)
        }else if(req.query.parentDistributor){
            self.checking(req.query).then(function (rep) {
                res.send(rep)
            })
        }

    }else{
        req.query.path=req.query.distributorName
        var exist=_.query(self.conf.distributor, {
            'distributorName': {$in : [req.query.distributorName]},
        });
        if(exist.length > 0){
            resObj.message='Distributor Already Exist'
            res.send(resObj)
        }else{
            self.conf.distributor.push(req.query)
            resObj.status=true
            resObj.data=self.conf.distributor
            resObj.message='CREATED SUCCESSFULLY'
            res.send(resObj)
        }

}



}
APIAction.prototype.checking=function (input) {
    var self=this;
    var parent=null
    return new Promise(function (resolve, reject) {
        console.log('inside checking');
        self.getParents(input).then(function (res) {
            parent=res
            return self.getAvilabePlace(res,input)
        }).then(function (res) {
            console.log('res in getAvilabePlace', res);
            if(res.status=='true' || res.status==true){
                input.path=parent[0].path+'.'+input.distributorName
                self.conf.distributor.push(input)
                resolve(self.conf.distributor)
            }else{
                resolve(res)
            }
        })
    })


}

APIAction.prototype.getParents=function (input) {
    var self=this;
    console.log('inside getParents');


    return new Promise.resolve(_.query(self.conf.distributor, {
        'distributorName': input.parentDistributor ,

    }))

}
APIAction.prototype.getAvilabePlace=function (input,req) {
    var self=this;
    console.log('get AvilabePlace checking');
    var temp=[]
    var status=true

    var tempAV=null
    var exClude=null
    var inClude=null
    var temp1,temp2,temp3=null
    var ex1,ex2,ex3=null
    return new Promise(function (resolve, reject) {
        self.getExclude(input).then(function (res) {
            exClude=res

            return self.getInclude(input)
        }).then(function (res) {

            inClude=res
            var condition={

            }
            var condition1={

            }
            var condition2={

            }


            if((exClude.exCountry.length > 0) && (inClude.inCountry.length == 0)){
                ex3=1
                // condition2['countryName']= {$nin: exClude.exCountry}
            }
            if((exClude.exProvince.length > 0) && (inClude.inProvince.length == 0)){
                ex2=1
                // condition1['provinceName']= {$nin: exClude.exCountry}
            }
            if((exClude.excity.length > 0) && (inClude.incity.length == 0)){
                ex1=1
            }



            if((exClude.exCountry.length > 0) && (inClude.inCountry.length > 0)){
                condition2['countryName']=
                        {$nin: exClude.exCountry,$in: inClude.inCountry}
                if(ex1==1){
                    condition2['cityName']= {$nin:exClude.excity}
                }
                if(ex2==1){
                    condition2['provinceName']= {$nin:exClude.exProvince}
                }

            }

            if((exClude.exCountry.length == 0) && (inClude.inCountry.length > 0)){
                condition2['countryName']= {$in: inClude.inCountry}
                if(ex1==1){
                    condition2['cityName']= {$nin:exClude.excity}
                }
                if(ex2==1){
                    condition2['provinceName']= {$nin:exClude.exProvince}
                }
            }





            if((inClude.inProvince.length > 0) && (exClude.exProvince.length > 0)){
                condition1['provinceName']={$nin: exClude.exProvince,$in: inClude.inProvince}
                if(ex1==1){
                    condition1['cityName']= {$nin:exClude.excity}
                }
                if(ex3==1){
                    condition1['countryName']= {$nin:exClude.exCountry}
                }
            }

            if((exClude.exProvince.length == 0) && (inClude.inProvince.length > 0)){
                condition1['provinceName']= {$in: inClude.inProvince}
                if(ex1==1){
                    condition1['cityName']= {$nin:exClude.excity}
                }
                if(ex3==1){
                    condition1['countryName']= {$nin:exClude.exCountry}
                }
            }



            if((inClude.incity.length > 0) && (exClude.excity.length > 0)){
                condition['cityName']={$nin: exClude.excity,$in: inClude.incity}
                if(ex2==1){
                    condition['provinceName']= {$nin:exClude.exProvince}
                }
                if(ex3==1){
                    condition['countryName']= {$nin:exClude.exCountry}
                }
            }

            if((exClude.excity.length == 0) && (inClude.incity.length > 0)){
                condition['cityName']= {$in: inClude.incity}
                if(ex2==1){
                    condition2['provinceName']= {$nin:exClude.exProvince}
                }
                if(ex3==1){
                    condition1['countryName']= {$nin:exClude.exCountry}
                }
            }


            if(!_.isEmpty(condition)){
                temp1=_.query(data,condition)
            }
            if(!_.isEmpty(condition1)){
                temp2=_.query(data,condition1)
            }
            if(!_.isEmpty(condition2)){
                temp3=_.query(data,condition2)
            }
            // console.log(temp1, temp2, temp3);
            return _.union(temp1,temp2,temp3)

        }).then(function (res) {
            tempAV=res
            temp[0]=req
            return self.getInclude(temp)

        }).then(function (final) {
            var resId=0
            var conditions={}

            if((final.incity.length > 0)){
                for(i=0 ; i<final.incity.length ; i++ ){
                    if(_.query(tempAV, {cityName: {$in: [final.incity[i]]}}).length > 0){
                        resId=1
                    }
                    else{
                        status=false
                        resolve({status:false,message: 'CITY: '+final.incity[i]+' Not Avilable For This Distributor'})
                    }
                }
            }
            if((final.inProvince.length > 0)){
                for(i=0 ; i<final.inProvince.length ; i++ ){
                    if(_.query(tempAV, {provinceName: {$in: [final.inProvince[i]]}}).length > 0){
                        resId=1
                    }
                    else{
                        status=false
                        resolve({status:false,message:'PROVINCE: '+final.inProvince[i]+' Not Avilable For This Distributor'})
                    }
                }
            }
            if((final.inCountry.length > 0)){
                for(i=0 ; i<final.inCountry.length ; i++ ){
                    if(_.query(tempAV, {countryName: {$in: [final.inCountry[i]]}}).length > 0){
                        resId=1
                    }
                    else{
                        status=false
                        resolve({status:false,message:'COUNTRY: '+final.inCountry[i]+' Not Avilable For This Distributor'})
                    }
                }
            }
            if(status=true && resId==1){
                resolve({status:true})
            }

        })

    })


}

APIAction.prototype.getExclude=function (input) {
 var self=this;
    var excity=[]
    var exProvince=[]
    var exCountry=[]
    var res={
        excity:[],
        exProvince:[],
        exCountry:[]
    }
    for(i=0 ; i< input[0].exclude.length ; i++){
        if(input[0].exclude[i].split('-').length==3){
            res.excity.push(input[0].exclude[i].split('-')[0])
        }else if(input[0].exclude[i].split('-').length==2){
            res.exProvince.push(input[0].exclude[i].split('-')[0])
        }else if(input[0].exclude[i]){
            res.exCountry.push(input[0].exclude[i])
        }
        if(input[0].exclude.length-1 == i){
            return new Promise.resolve(res)

        }



    }


}

APIAction.prototype.getInclude=function (input) {
 var self=this;


    var res={
        incity:[],
        inProvince:[],
        inCountry:[]
    }
    for(i=0 ; i< input[0].include.length ; i++){
        if(input[0].include[i].split('-').length==3){
            res.incity.push(input[0].include[i].split('-')[0])
        }else if(input[0].include[i].split('-').length==2){
            res.inProvince.push(input[0].include[i].split('-')[0])
        }else if(input[0].include[i]){
            res.inCountry.push(input[0].include[i])
        }
        if(input[0].include.length-1 == i){
            return new Promise.resolve(res)

        }



    }


}

APIAction.prototype.list=function (req,res) {
    var self=this
    res.send(self.conf.distributor)
}
