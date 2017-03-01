var CsvParser 			= require('csv-to-json');

const INPUT_FILE		= './server/data/cities.csv';
const OUTPUT_FILE		= './server/data/cities.json';


var parseOptions		= {filename : INPUT_FILE};
var writeOptions		= {filename : OUTPUT_FILE};

CsvParser.parse(parseOptions, function (err, json) {
	if(err){
		console.log(err);
		return
	}

	writeOptions.json = convertData(json);
	CsvParser.writeJsonToFile(writeOptions, function (err) {
		if(err){
			console.log(err);
		}
	})

})

var convertData = function (arrData) {
	var resultData			= new Object();
	var countyHash			= new Object();



	arrData.forEach(function (objData) {
		countyHash[objData['Country_Code']] ={
			'Country_Code' : objData['Country_Code'],
			'Country_Name' : objData['Country_Name']
		}
	})

	resultData.cities 		= arrData;
	resultData.countries 	= Object.keys(countyHash).map(function (key) {
		return countyHash[key];
	})
	return resultData;
}
