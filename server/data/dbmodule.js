const DATA_FILE		= './server/data/cities.json';
var fs				= require('fs');
var Data 			= JSON.parse(fs.readFileSync(DATA_FILE, 'utf8'));

var CsvParser 		= require('csv-to-json');

console.log('dir',process.cwd());

var writeToFile		= function (data) {

	var writeOptions		= {
		filename : DATA_FILE,
		json	 : data
	};
	return new Promise(function (resolve, reject) {
		CsvParser.writeJsonToFile(writeOptions, function (err) {
			if(err){
				return reject(err);
			}
			return resolve({status: true});
		});
	});
}


/**
 * makes find query on model
 * @param {String} model modelname
 * @param {Object} objQuery queryObject
 * @param {Number} skip will skip this much of first records
 * @param {Number} limit will limit the no of records
 */
exports.find	= function (model, objQuery, skip, limit) {
	return new Promise(function (resolve, reject) {
		var foundData	= Data[model].filter(function (objData) {
			var blStatus = true;
			for(key in objQuery){
				blStatus = blStatus && objQuery[key] == objData[key];
			}
			return blStatus;
		});
		if(!skip ){
			skip = 0;
		}
		if(!limit){
			limit = foundData.length;
		}
		return resolve(foundData.slice(skip, limit+1));

	});
}

/**
 * makes inserts documet
 * @param {String} model modelname
 * @param {Object} objData dataObject
 */
exports.insert	= function (model, objData) {
	if(!Data[model]){
		Data[model] = []
	}
	Data[model].push(objData);
	return writeToFile(Data);
}


/**
 * makes find and update  on model
 * @param {String} model modelname
 * @param {Object} objQuery queryObject
 * @param {Object} objUpdate update document
 */
exports.update	= function (model, objQuery, objUpdate) {
	var updatedData	= Data[model].map(function (objData) {
		var blStatus = true;
		for(key in objQuery){
			blStatus = blStatus && objQuery[key] == objData[key];
		}
		if(blStatus){
			for(key in objData['$set']){
				objData[key] = objUpdate[key];
			}
			for(key in objData['$push']){
				objData[key].push(objUpdate[key]);
			}
		}
		return objData
	});
	return writeToFile(updatedData);
}
