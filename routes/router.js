var router = require('express').Router();

var fs = require('fs');

var csv = require('fast-csv');

var dataResource;

var destributersList  = new Array();

router.get('/', welcomeMessage);

router.post('/checkAuthorization', ceheckAuthorization);

router.get('/getDistributersList', getDistributersList);

function welcomeMessage(req, res){

	res.send('welcome to api');

	readDataFromRawFile();
}

function ceheckAuthorization(req, res){

	var reqData = req.body;

	var destributer = reqData.destributer;

	var city = reqData.city;

	var province = reqData.province;

	var country = reqData.country;	

	var includeCountry = new Array();

	var includeProvince = new Array();

	var includeCity = new Array();

	var excludeCountry = new Array();

	var excludeProvince = new Array();

	var excludeCity = new Array();

	var includesObj = {};

	var excludesObj = {};

	if(dataResource == undefined) readDataFromRawFile();

	dataResource.forEach(function(destributerDetails){

		if(destributerDetails.id.toLowerCase() == destributer.toLowerCase()){

			include = destributerDetails.include;

			include.forEach(function(singleInclude){

				if(includeCity.indexOf(singleInclude.city) == -1)includeCity.push(singleInclude.city);

				if(includeProvince.indexOf(singleInclude.provision) == -1)includeProvince.push(singleInclude.provision);

				if(includeCountry.indexOf(singleInclude.countryname) == -1)includeCountry.push(singleInclude.countryname);
			
			});

			exclude = destributerDetails.exclude;

			exclude.forEach(function(singleExclude){

				if(excludeCity.indexOf(singleExclude.city) == -1)excludeCity.push(singleExclude.city);

				if(excludeProvince.indexOf(singleExclude.provision) == -1)excludeProvince.push(singleExclude.provision);

				if(excludeCountry.indexOf(singleExclude.countryname) == -1)excludeCountry.push(singleExclude.countryname);
			});

		}else{

			if(destributerDetails.hasOwnProperty('sub-dest')){

				var destDetails = destributerDetails['sub-dest'];

				destDetails.forEach(function(oneDest){

					if(oneDest.id.toLowerCase() == destributer.toLowerCase()){

						parentInclude = destributerDetails.include;

						parentInclude.forEach(function(singleInclude){

							if(includeCity.indexOf(singleInclude.city) == -1)includeCity.push(singleInclude.city);

							if(includeProvince.indexOf(singleInclude.provision) == -1)includeProvince.push(singleInclude.provision);

							if(includeCountry.indexOf(singleInclude.countryname) == -1)includeCountry.push(singleInclude.countryname);
						});

						parentExclude = destributerDetails.exclude;

						parentExclude.forEach(function(singleExclude){

							if(excludeCity.indexOf(singleExclude.city) == -1)excludeCity.push(singleExclude.city);

							if(excludeProvince.indexOf(singleExclude.provision) == -1)excludeProvince.push(singleExclude.provision);

							if(excludeCountry.indexOf(singleExclude.countryname) == -1)excludeCountry.push(singleExclude.countryname);
						});

						include = oneDest.include;

						include.forEach(function(singleInclude){

							if(includeCity.indexOf(singleInclude.city) == -1)includeCity.push(singleInclude.city);

							if(includeProvince.indexOf(singleInclude.provision) == -1)includeProvince.push(singleInclude.provision);

							if(includeCountry.indexOf(singleInclude.countryname) == -1)includeCountry.push(singleInclude.countryname);
						});

						exclude = oneDest.exclude;

						exclude.forEach(function(singleExclude){

							if(excludeCity.indexOf(singleExclude.city) == -1)excludeCity.push(singleExclude.city);

							if(excludeProvince.indexOf(singleExclude.provision) == -1)excludeProvince.push(singleExclude.provision);

							if(excludeCountry.indexOf(singleExclude.countryname) == -1)excludeCountry.push(singleExclude.countryname);
						});

					}
				});
			}
		}
	});

	var responseObect = {};

	if(includeCountry.indexOf(country) != -1){

		if(includeProvince.indexOf(province) != -1 || province == ""){

			if(((excludeCity.indexOf(city) == -1)) && ((includeCity.indexOf(city) != -1) || city == "" || city == undefined)){

				sendSuccessResponse(res);
			}else{
				sendFailureResponse(res)
			}
		}else{
			sendFailureResponse(res)
		}
	}else{
		sendFailureResponse(res)
	}
}

function sendSuccessResponse(res){
	var resobj = {};
	res.statusCode = 200;
	resobj.message = "success";
	resobj.data = "YES";
	res.send(resobj);
}

function sendFailureResponse(res){
	var resobj = {};
	res.statusCode = 402;
	resobj.message = "failure";
	resobj.data = "NO";
	res.send(resobj);
}


function getDistributersList(req, res){

	if(dataResource == undefined){

		readDataFromRawFile();
	}

	destributersList = new Array();

	prepareDestributersList(dataResource);	

	var responseObject = {};

	responseObject.status = 200;

	responseObject.data = destributersList;

	responseObject.message = "Operation completed successfully";

	res.status = 200;

	res.send(responseObject);
}

function prepareDestributersList(arrayOfDestributers){

	arrayOfDestributers.forEach(function(destributer){

		destributersList.push(destributer.id);

		if(destributer.hasOwnProperty('sub-dest')){

			prepareDestributersList(destributer['sub-dest']);
		}
	});
}

function readDataFromRawFile(){

	var data = fs.readFileSync(__dirname+"/../raw/input_data.json");

	dataResource = JSON.parse(data.toString());
}

module.exports = router;