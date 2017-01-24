var fs = require('fs');
var path = require('path');
var csv = require('csv-parser');
var json2csv = require('json2csv');
var express = require('express');
var router = express.Router();

console.log("Server started");

function getuniqueResults(arr){
    var result = arr.filter(function (elem, pos, arr) {
        return arr.indexOf(elem) === pos;
    })
    return result;
}

router.get('/Home', function(req, res) {
    var dList = {select : 'Select'};
    fs.readdir(path.join(__dirname, '../', 'public', 'resource'), function(err, files){
        if(!err){
            files.forEach(function(FileName){
                FileName = FileName.split('.')[0];
                dList[FileName] = FileName;
            })
            res.send(dList);
        }else{
            process.stdout.write(`Error on reading /resource files {$err}`);
        }
    });
});

router.post('/getStatus', function(req, res) {
    var countryName = req.body.country;  
    var provinceName = req.body.province;
    var cityName = req.body.city;
    var fileName = req.body.fileName;
    var addNewFile = req.body.addNewFile;
    try{
        if(fileName && countryName && countryName!= 'Select'){
            var success = 'Yes, You are allowed to distribute in the given area';
            var failure = 'Sorry, Incorrect/Not Allowed to distribute in the given';
            var listOnCountry = [];
            fs.createReadStream(path.join(__dirname, '../', 'public', 'resource', fileName + '.csv'))
            .pipe(csv())
            .on('data', function (data) {
                if(data['Country Name'] == countryName){
                    listOnCountry.push(data);
                }
            }) 
            .on('end', function () {
                if(listOnCountry && listOnCountry.length>0){
                    if(provinceName && provinceName!= 'Select'){
                        var listOnProvince = [];
                        listOnCountry.forEach(function(data){
                            if(data['Province Name'] == provinceName){
                                listOnProvince.push(data);
                            }
                        })
                        if(listOnProvince && listOnProvince.length>0){
                            if(cityName && cityName!= 'Select'){
                                var listOnCity = [];
                                listOnProvince.forEach(function(data){
                                    if(data['City Name'] == cityName){
                                        listOnCity.push(data);
                                    }
                                })
                                if(listOnCity && listOnCity.length>0){
                                    if(addNewFile){
                                        res.send(listOnCity);}
                                     else{
                                        res.send(success);}
                                }else{
                                    res.send(failure + ' City');
                                }
                            }else{
                                if(addNewFile){
                                    res.send(listOnProvince);} 
                                 else{
                                    res.send(success);}
                            }
                        }else{
                            res.send(failure + ' Province');
                        }
                    }else{ 
                        if(addNewFile){
                            res.send(listOnCountry);}
                        else{
                            res.send(success);}
                    }
                }else{
                    res.send(failure + ' Country');
                }
            })
        }else{
            res.send('Sorry, No Country Name Received');
        }
    }catch(err){
        console.log('OOPS error on /getStatu! :'+err)
    }
});

router.post('/getCountry', function(req, res) {
    try{
        var fileName = req.body.nameOfFile;
        var countryNames = ['Select'];
        if(fileName){
            fs.createReadStream(path.join(__dirname, '../', 'public', 'resource', fileName + '.csv'))
            .pipe(csv())
            .on('data', function (data) {
                countryNames.push(data['Country Name']);
            }) 
            .on('end', function (err) {
                if(countryNames && countryNames.length>0){
                    countryNames = getuniqueResults(countryNames);
                    res.send(countryNames);
                }else if(err){
                    res.send('Error on Fetching Available County List');
                }
            })
        }
    }catch(err){
        console.log('OOPS error on /getCountry! :'+err);
    }
});

router.post('/getProvince', function(req, res) {
    try{
        var countryName = req.body.countryList;
        var fileName = req.body.nameOfFile;
        var provinceNames = ['Select'];
        if(fileName){
            fs.createReadStream(path.join(__dirname, '../', 'public', 'resource', fileName + '.csv'))
            .pipe(csv())
            .on('data', function (data) {
                if(data['Country Name'] == countryName){
                    provinceNames.push(data['Province Name']);
                }
            }) 
            .on('end', function (err) {
                if(provinceNames && provinceNames.length>0){
                    provinceNames = getuniqueResults(provinceNames);
                    res.send(provinceNames);
                }else if(err){
                    res.send('Error on Fetching Available Province List');
                }
            })
        }
    }catch(err){
        console.log('OOPS error on /getProvince! :'+err);
    }
});

router.post('/getCity', function(req, res) {
    try{
        var provinceName = req.body.provinceList;
        var fileName = req.body.nameOfFile;
        var setCityList = ['Select'];
        if(fileName){
            fs.createReadStream(path.join(__dirname, '../', 'public', 'resource', fileName + '.csv'))
            .pipe(csv())
            .on('data', function (data) {
                if(data['Province Name'] == provinceName){
                    setCityList.push(data['City Name']);
                }
            }) 
            .on('end', function (err) {
                if(setCityList && setCityList.length>0){
                    setCityList = getuniqueResults(setCityList);
                    res.send(setCityList);
                }else if(err){
                    res.send('Error on Fetching Available Province List');
                }
            })
        }
    }catch(err){
        console.log('OOPS error on /getCity! :'+err);
    }
});

router.post('/newDistributer', function(req, res) {
    var addDistributor = req.body.addDistributorList;
    var newDistribution = req.body.newDistribution;
    var parentDistributor = req.body.parentDistributor;
    var editMode = req.body.editMode;
    var fields = ['City Code', 'Province Code', 'Country Code', 'City Name','Province Name','Country Name'];
    try {
        var result = json2csv({ data: addDistributor, fields: fields });
        fs.writeFile(path.join(__dirname, '../', 'public', 'resource', newDistribution + '.csv'), result, 'utf8', function () {
          {
            if(!editMode){
                updateDetails(parentDistributor, newDistribution)
            }else{
                res.send(`Details of ${newDistribution} has been saved successfully.`);
            }
          } 
        });
        
    } catch (err) {
        console.error('Error in storing new distributor data');
    }
    function updateDetails(parentDistributor, newDistribution){
        var filePath = path.join(__dirname, '../', 'public', 'details' , 'details.json');
        var childAdded = false;
        fs.readdir(path.join(__dirname, '../', 'public', 'details'), (err, files) => {
          if(files.length > 0){
              files.forEach(function(file){
                  if(file == 'details.json'){
                      fs.readFile(filePath, 'utf-8', function(err,content){
                          var fileObj = content;   
                          fileObj = JSON.parse(fileObj);
                          for (i of fileObj.data) { 
                                if (i.parent == parentDistributor) { 
                                    childAdded = true;
                                    var nameExist = false;
                                    var checkChild = i.child.split(',');
                                    checkChild.forEach(function(child){
                                         if(child.toUpperCase() == newDistribution.toUpperCase()){
                                            nameExist = true;
                                            res.send('Distributor Name already exists.Please change');
                                         }
                                    })
                                    if(!nameExist){
                                    i.child = i.child + ',' + newDistribution;
                                    writeDetails(filePath, fileObj);
                                    }
                                }
                           }
                          if(!childAdded){
                              fileObj.data.push({parent:parentDistributor,child:newDistribution});
                              writeDetails(filePath, fileObj);
                             } 
                      })
                  }
              })
          }else{
              var detailsObj = {data:[{parent:parentDistributor,child:newDistribution}]};
              writeDetails(filePath, detailsObj);
          }
        })
    }
    function writeDetails(filePath, fileObj){
        var result = JSON.stringify(fileObj);
        fs.writeFile(filePath, result, 'utf8', function () {
                res.send(`Details of ${newDistribution} has been saved successfully.`);
        });
    }
});

router.post('/getChild', function(req, res) {
    var parentDistributor = req.body.parent;
    var child;
    fs.readdir(path.join(__dirname, '../', 'public', 'details'), (err, files) => {
          if(files.length > 0){
              files.forEach(function(file){
                  if(file == 'details.json'){
                    var filePath = path.join(__dirname, '../', 'public', 'details' , 'details.json');
					fs.readFile(filePath, 'utf-8', (err,content) => {
                        var fileObj = content;   
                        fileObj = JSON.parse(fileObj);
                        for (i of fileObj.data) { 
                            if (i.parent == parentDistributor) { 
                                child = i.child;
                            }
                        }
                        res.send(child);
                    })
				  }
			  })
	      }
    })
});

router.post('/EditChild', function(req, res) {
    var childDistributor = req.body.childDistributor;
    var result =[];
    if(childDistributor){
        fs.createReadStream(path.join(__dirname, '../', 'public', 'resource', childDistributor + '.csv'))
        .pipe(csv())
        .on('data', function (data) {
            result.push(data);
        }) 
         .on('end', function (err) {
            res.send(result);
            if(err){
                res.send("Sorry!!! Couldn't fetch details");
            }
        })
    }
});

module.exports = router;     