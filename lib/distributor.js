/**
 * Created by hmspl on 17/3/17.
 */

const async = require('async');
const _     = require('underscore');

let distributorData    = require('../data/distributor.json');
let subDistributorData = require('../data/sub_distributor.json');

let distributor = {};

distributor.findDistributionPermit = function (request, callback) {
    let payload       = request.payload;
    let params        = request.params;
    let distributorId = params.distributorId;
    let data          = [];
    let tasks         = [];

    //Get Data & combine distributor data & Sub Distributor data of the distributor
    tasks.push(function (innerCb) {
        let distData    = _.where(distributorData, {distributorId: distributorId});
        let subDistData = _.where(subDistributorData, {distributorId: distributorId});
        data.push.apply(data, distData);
        data.push.apply(data, subDistData);
        if (data && data.length) {
            return innerCb();
        } else {
            return innerCb({code: 425, message: "Distributor Not Exists"});
        }
    });

    //Get All includes data of distributor and form an array with matched data
    tasks.push(function (innerCb) {
        let includesArr = _.flatten(_.pluck(data, "includeData"));
        doCheckInIncludesData(includesArr, payload, function (err, resp) {
            if (err) {
                return innerCb(err);
            } else {
                console.log('INC DATA ', resp);
                return innerCb(null, resp);
            }
        });
    });

    //compares matched data with exclude array and form final array
    tasks.push(function (matchedArr, innerCb) {
        let excludesArr = _.flatten(_.pluck(data, "excludeData"));
        doCheckInExcludedData(matchedArr, excludesArr, function (err, resp) {
            if (err) {
                return innerCb(err);
            } else {
                console.log('FINAL DATA ', resp);
                return innerCb(null, resp);
            }
        });
    });

    async.waterfall(tasks, function (err, resp) {
        if (err) {
            return callback(err);
        } else {
            let resObj = {
                statusCode: 200,
                message   : "Success",
                hasRights : false
            };
            if (resp && resp.length) {
                resObj.hasRights = true;
            }
            return callback(null, resObj);
        }
    });
};

/**
 *
 * @param includesArr - Array - all include array
 * @param payload - Object - payload
 * @param callback function
 */
function doCheckInIncludesData(includesArr, payload, callback) {
    let resArr = [];
    console.log('ILL ', includesArr, payload);
    async.forEachSeries(includesArr, function (oneData, cback) {

        let innerTasks = [];
        innerTasks.push(function (innerCb) {
            if (payload.countryCode) {
                checkMatches(oneData, "countryCode", payload.countryCode, function (err, isOkay) {
                    if (isOkay) {
                        return innerCb(null, true);
                    } else {
                        return innerCb(true);
                    }
                });
            } else {
                return innerCb(null, true);
            }
        });

        innerTasks.push(function (innerCb) {
            if (payload.provinceCode) {
                checkMatches(oneData, "provinceCode", payload.provinceCode, function (err, isOkay) {
                    if (isOkay) {
                        return innerCb(null, true);
                    } else {
                        return innerCb(true);
                    }
                });
            } else {
                return innerCb(null, true);
            }
        });

        innerTasks.push(function (innerCb) {
            if (payload.cityCode) {
                checkMatches(oneData, "cityCode", payload.cityCode, function (err, isOkay) {
                    if (isOkay) {
                        return innerCb(null, true);
                    } else {
                        return innerCb(true);
                    }
                });
            } else {
                return innerCb(null, true);
            }
        });

        async.parallel(innerTasks, function (err, resp) {
            if (err) {
                return cback();
            } else {
                resArr.push(oneData);
                return cback();
            }
        });

    }, function (err, resp) {
        return callback(null, resArr);
    });
}

/**
 *
 * @param includeArr - Array - matched Arr
 * @param excludesArr - Array - Excluded Arr
 * @param callback - function
 */
function doCheckInExcludedData(includeArr, excludesArr, callback) {
    let finalArr = [];

    async.forEachSeries(excludesArr, function (excData, cbk) {
        async.forEachSeries(includeArr, function (incData, cback) {

            let innerTasks     = [];
            let countryMatches = false, provinceMatches = false, cityMatches = false;
            innerTasks.push(function (innerCb) {
                checkMatches(excData, "countryCode", incData.countryCode, function (err, isOkay) {
                    if (isOkay) {
                        countryMatches = true;
                        return innerCb(null, true);
                    } else {
                        return innerCb(true);
                    }
                });
            });

            innerTasks.push(function (innerCb) {
                checkMatches(excData, "provinceCode", incData.provinceCode, function (err, isOkay) {
                    if (isOkay) {
                        provinceMatches = true;
                        return innerCb(null, true);
                    } else {
                        return innerCb(true);
                    }
                });
            });

            innerTasks.push(function (innerCb) {
                checkMatches(excData, "cityCode", incData.cityCode, function (err, isOkay) {
                    if (isOkay) {
                        cityMatches = true;
                        return innerCb(null, true);
                    } else {
                        return innerCb(true);
                    }
                });
            });

            async.parallel(innerTasks, function (err, resp) {
                if (!(countryMatches && provinceMatches && cityMatches)) {
                    finalArr.push(incData);
                }
                return cback();
            });
        }, function (err, resp) {
            return cbk();
        });
    }, function (err, resp) {
        return callback(null, finalArr);
    });
}

/**
 *
 * @param data - Object - Include or Exclude data
 * @param key - String - countryCode or provinceCode or cityCode
 * @param value - String - Value to compare
 * @param callback - function
 * @returns {*}
 */
function checkMatches(data, key, value, callback) {
    if (data[key] && (data[key] == value || data[key] == "ALL")) {
        return callback(null, true);
    } else {
        return callback(null, false);
    }
}

module.exports = distributor;