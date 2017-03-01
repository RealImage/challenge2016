const DATA_FILE		= './server/data/cities.json';
var fs				= require('fs');
var Data 			= JSON.parse(fs.readFileSync(DATA_FILE, 'utf8'));

var CsvParser 		= require('csv-to-json');

const OBJECT		= 'object';


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

var nextId = function () {
	if(Data.id){
		a = parseInt(Data.id, 16) + 1;
		a = a.toString(16)
		Data.id = a;
		writeToFile(Data);
		return a;
	}
	Data.id = '111111111';
	writeToFile(Data);
	return '111111111';
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
				if(typeof(objQuery[key]) == OBJECT){
					for(subKey in objQuery[key]){
						if(subKey == '$in'){
							if(!isArray(objQuery[key][subKey])){
								return reject(new Error('Type error expected array at $.'+ key +'.'+subKey));
							}
							blStatus = blStatus && objQuery[key][subKey].indexOf(objData[key]) != -1;
						}
						else if(subKey == '$nin'){
							if(!isArray(objQuery[key][subKey])){
								return reject(new Error('Type error expected array at $.'+ key +'.'+subKey));
							}
							blStatus = blStatus && objQuery[key][subKey].indexOf(objData[key]) == -1;
						}
						else {
							blStatus = false
						}
					}
				}
				else {
					blStatus = blStatus && objQuery[key] == objData[key];
				}
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
	objData._id = nextId();
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
	Data[model]	= Data[model].map(function (objData) {
		var blStatus = true;
		for(key in objQuery){
			blStatus = blStatus && objQuery[key] == objData[key];
		}
		if(blStatus){
			for(key in objUpdate['$set']){
				objData[key] = objUpdate['$set'][key];
			}
			for(key in objUpdate['$push']){
				objData[key].push(objUpdate['$push'][key]);
			}
		}
		return objData
	});
	return writeToFile(Data);
}

var isArray = function (obj) {
	return !!obj.push;
}
