const _ = require('lodash'),
    debug = require('debug'),
    csv = require('csvtojson'),
    path = require('path'),
    jsonfile = require('jsonfile');
const csvFilePath = path.join(__dirname, '../', 'distributors', "input" + '.csv'),
    out = require('../distributors/out');

const log = debug('realImage.controller');

//let item = _.pickBy(jsonObj, _.identity);


//@todo write general function to Generate the Rules file
//=====This function will creates the Rules.json======//
function generateRules(fileName, callback) {
    let distrubutionDetails = [];
    let distributorfilepath = path.join(__dirname, '../', 'distributors', fileName + '.csv');
    csv().fromFile(distributorfilepath)
        .on('json', (jsonObj)=> {
            if (jsonObj['PARENT-ID'].length === 0) {
                let exclude = jsonObj.EXCLUDE.split(',').map(function (list) {
                    console.log("adf", list.split('-').length);
                    let bn = list.split('-');
                    if (bn.length == 1) {
                        console.log("India");
                        return {
                            parent_id: "",
                            id: jsonObj.Distrubutor,
                            permission: "NO",
                            country_name: list
                        };
                    }
                    else if (bn.length == 2) {
                        return {
                            parent_id: "",
                            id: jsonObj.Distrubutor,
                            permission: "NO",
                            country_name: bn[1],
                            state_name: bn[0]
                        };

                    }
                    else if (bn.length == 3) {
                        return {
                            parent_id: "",
                            id: jsonObj.Distrubutor,
                            permission: "NO",
                            city_name: bn[0],
                            state_name: bn[1],
                            country_name: bn[2]

                        };
                    }

                });
                distrubutionDetails.push.apply(distrubutionDetails, exclude);
                let include = jsonObj.INCLUDE.split(',').map(function (list2) {
                    return {
                        parent_id: "",
                        id: jsonObj.Distrubutor,
                        permission: "Yes",
                        country_name: list2
                    };
                });
                distrubutionDetails.push.apply(distrubutionDetails, include);

            }
            else {
                let exclude = jsonObj.EXCLUDE.split(',').map(function (list) {
                    console.log("adf", list.split('-').length);
                    let bn = list.split('-');
                    if (bn.length == 1) {
                        console.log("India");
                        return {
                            parent_id: jsonObj['PARENT-ID'],
                            id: jsonObj.Distrubutor,
                            permission: "NO",
                            country_name: list
                        };
                    }
                    else if (bn.length == 2) {
                        return {
                            parent_id: jsonObj['PARENT-ID'],
                            id: jsonObj.Distrubutor,
                            permission: "NO",
                            country_name: bn[1],
                            state_name: bn[0]
                        };

                    }
                    else if (bn.length == 3) {
                        return {
                            parent_id: jsonObj['PARENT-ID'],
                            id: jsonObj.Distrubutor,
                            permission: "NO",
                            city_name: bn[0],
                            state_name: bn[1],
                            country_name: bn[2]

                        };
                    }

                });
                distrubutionDetails.push.apply(distrubutionDetails, exclude);
                let include = jsonObj.INCLUDE.split(',').map(function (list2) {
                    let bn = list2.split('-');

                    if (bn.length == 1) {
                        return {
                            parent_id: jsonObj['PARENT-ID'],
                            id: jsonObj.Distrubutor,
                            permission: "Yes",
                            country_name: list2
                        };
                    }
                    else if (bn.length == 2) {
                        return {
                            parent_id: jsonObj['PARENT-ID'],
                            id: jsonObj.Distrubutor,
                            permission: "YES",
                            country_name: bn[1],
                            state_name: bn[0]
                        };

                    }
                    else if (bn.length == 3) {
                        return {
                            parent_id: jsonObj['PARENT-ID'],
                            id: jsonObj.Distrubutor,
                            permission: "YES",
                            city_name: bn[0],
                            state_name: bn[1],
                            country_name: bn[2]

                        };
                    }
                });
                distrubutionDetails.push.apply(distrubutionDetails, include);
            }
        }).on('done', (error)=> {
        callback(null, distrubutionDetails);
    })
}

exports.rules = function (req, res, next) {
    const rulesJson = path.join(__dirname, '../', 'distributors', "rules" + '.json');
    generateRules("Permissions", function (err, result) {
        if(err) next(err);
        jsonfile.spaces = 4;
        jsonfile.writeFile(rulesJson, {"permisions": result.filter(x => (x.country_name !== ''))}, function (err) {
            return res.status(201).end({ message: 'Succsessfully Generated the Rules Json'});
        })
    })
};


exports.getDistributionStatus = function (req, res) {
    let distributionObject = _.pick(req.body, "distributorName", "countryName", "stateName", "cityName");
    if (distributionObject.distributorName && distributionObject.countryName) {
        let filepath = path.join(__dirname, '../', 'distributors', "rules" + '.json');
        jsonfile.readFile(filepath, function (err, result) {
            let distributorData = _.filter(result.permisions, {"id": distributionObject.distributorName});
            if (distributorData.length > 0) {
                // Get the Parents and Child Distributor Data
                out.getData(distributionObject.distributorName, function (err, allData) {
                    distributorData = allData;

                    console.log('distributorData.length',distributorData.length);
                    //   distributorData.push.apply(distributorData, example);
                    let countryList = _.filter(distributorData, {"country_name": distributionObject.countryName});

                    //Now we have to start comparing  input data(distributionObject) with Distrubutor Permission data(rules.json)
                    if (countryList.length > 0) {
                        if (distributionObject.stateName && distributionObject.cityName) {

                            statecheckAndCityCheck(distributionObject, countryList, function (err, data) {
                                res.json({"message": data});
                            })
                        } else if (distributionObject.stateName) {
                            stateCheck(distributionObject, countryList, function (err, data) {
                                res.json({"message": data});
                            })
                        }
                        else {
                            res.json({message: "Yes,you are allowed to Distribute In This Country " + "===>" + distributionObject.countryName});
                        }

                    } else {
                        res.send({message: "No, You are Not allowed to Distribute In This Country " + "===>" + distributionObject.countryName});
                    }
                })
            }
            else {
                res.send("Invalid Distribuor Name");
            }
        })

    }
    else {
        res.send("distributorName or countryName Couldn't be empty");
    }

};

// Function returns the main data From the Input.csv
//@todo Wite a function that should Returns the data of Input.csv
function getSourceData(filename, callback) {
    let sourcefilepath = path.join(__dirname, '../', 'distributors', filename + '.csv'),
        sourceData = [];
    csv().fromFile(sourcefilepath)
        .on('json', (jsonObj)=> {
            let filterobject = _.omit(jsonObj, ['City Code', 'Province Code', 'Country Code']);
            sourceData.push(filterobject);
        }).on('done', (error)=> {
        callback(null, sourceData)
    })
}

// This Funtion will  Give only Permission for the  States and Cities Present in the Input.csv, Not allowed other Data
//@todo Write a function that accepts only city,state data present in the Input.csv ====> Input.csv is our  DataBase here
function statecheckAndCityCheck(distributionObject, countryList, maincallback) {
    //Fetching the Source data from the Input.csv
    getSourceData("input", (err, source)=> {
        let sourceData = _.filter(source, {"country_name": distributionObject.countryName});
        let stateAndCity = _.map(sourceData, (item)=> {
            return _.pick(item, ['city_name', 'state_name'])
        });

        // Fetching state and city only from the Source data
        let checkStateAndCity = _.find(stateAndCity, {
            'city_name': distributionObject.cityName,
            'state_name': distributionObject.stateName
        });

        // Checking the Permission from the rules file with NO
        let cityCheck = _.filter(countryList, {
            "state_name": distributionObject.stateName, "permission": "NO",
            "city_name": distributionObject.cityName
        }).length;
        let statecheck = _.filter(countryList, {
            "state_name": distributionObject.stateName, "permission": "NO"
        }).length;

        if (!cityCheck && !statecheck) {
            if (checkStateAndCity) {
                let message = "Yes, you are allowed to Distribute in the City" + " ===> " + distributionObject.cityName;
                maincallback(null, message);
            }
            else {
                let message = "NO, you are allowed to Distribute in the City" + " ===> " + distributionObject.cityName;
                maincallback(null, message);
            }
        }
        else {
            let message = "NO, you are allowed to Distribute in the City" + " ===> " + distributionObject.cityName;
            maincallback(null, message);
        }


    })
}
//Check Perticular City
function checkPerticularCity(distributionObject, countryList, callback) {
    let sourceCity = _.filter(countryList, {
        "state_name": distributionObject.stateName,
        "city_name": distributionObject.cityName,
        permission: 'YES'
    }).length;

    let count = countryList.filter(function (a) {
        return a.city_name && a.permission === 'YES'
    }).length;

    let result = {found: sourceCity, totalCity: count};
    console.log(result);
    callback(null, result);
}


// This Funtion will  Give only Permission for the  States Present in the Input.csv, Not allowed other Data
//@todo Write a function that accepts only state  present in the Input.csv ====> Input.csv is our  DataBase here
function stateCheck(distributionObject, countryList, maincallback1) {
    //Fetching the Source data from the Input.csv
    getSourceData("input", (err, source)=> {
        let sourceData = _.filter(source, {"country_name": distributionObject.countryName});
        /*  _.uniq(_.map(data, 'usernames'))*/
        let sourceStates = _.map(sourceData, (item)=> {
            return _.pick(item, ['state_name'])
        });

        // Fetching state and city only from the Source data
        let stateExisists = _.find(sourceStates, {
            'state_name': distributionObject.stateName
        });

        let statecheck = _.filter(countryList, {
            "state_name": distributionObject.stateName, "permission": "NO"
        }).length;

        if (!statecheck) {
            if (stateExisists) {
                let message = "Yes, you are allowed to Distribute in the State" + " ===> " + distributionObject.stateName;
                maincallback1(null, message);
            }
            else {
                let message = "NO, you are allowed to Distribute in the State" + " ===> " + distributionObject.stateName;
                maincallback1(null, message);
            }
        }
        else {
            let message = "NO, you are allowed to Distribute in the State" + " ===> " + distributionObject.stateName;
            maincallback1(null, message);
        }
    })
}