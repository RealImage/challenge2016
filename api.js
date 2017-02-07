exports.getTerritories = function (callback) {
	var csv = require('csvtojson'),
	Converter = require("csvtojson").Converter,
	converter = new Converter({});

	converter.fromFile('./cities.csv', function (err, result) {
		if (err)
			callback(true, result);
		else {
			var countries = {},
			provinces = {},
			cities = {};
			for (var i = 0; i < result.length; i++) {
				if (!countries[result[i]["Country Code"]])
					countries[result[i]["Country Code"]] = result[i]["Country Name"];

				if (!provinces[result[i]["Country Code"] + "_" + result[i]["Province Code"]])
					provinces[result[i]["Country Code"] + "_" + result[i]["Province Code"]] = result[i]["Province Name"];

				if (!cities[result[i]["Country Code"] + "_" + result[i]["Province Code"] + "_" + result[i]["City Code"]])
					cities[result[i]["Country Code"] + "_" + result[i]["Province Code"] + "_" + result[i]["City Code"]] = result[i]["City Name"];
			}

			console.log("Territory Loaded Successful");
			callback(null, {
				countries : countries,
				provinces : provinces,
				cities : cities
			});
		}
	});
};

exports.distributors = function (dataBase, config, territories, req, res, next) {
	function validateDetails() {
		//Flat validation for readability
		var errorArray = [], superDistributor;
		if (!req.body.data)
			errorArray.push("Invalid distributor details.");

		if (req.body.data && !req.body.data["distributorName"])
			errorArray.push("Distributor Name is required.");

		if (req.body.data && req.body.data["distributorName"] && !req.body.data["id"]) {
			for (var prop in dataBase) {
				if (dataBase[prop]["distributorName"] == req.body.data["distributorName"]) {
					errorArray.push("Distributor Name is already exists.");
					break;
				}
			}
		}

		if (req.body.data && req.body.data["superDistributor"]) {
			if (req.body.data && req.body.data["distributorName"] && req.body.data["superDistributor"] == req.body.data["distributorName"])
				errorArray.push("Distributor Name and Super Distributor can't be same.");

			var itemFound = false;
			for (var prop in dataBase) {
				if (dataBase[prop]["distributorName"] == req.body.data["superDistributor"]) {
					itemFound = true;
					break;
				}
			}

			if (itemFound)
				superDistributor = exports.getPermittedLocations(dataBase, territories, req.body.data["superDistributor"]);
			else
				errorArray.push("Super Distributor Name is invalid.");
		}

		if (req.body.data && (!req.body.data["rules"] || (req.body.data["rules"] && req.body.data["rules"].length == 0)))
			errorArray.push("Atleast one rule details is required.");

		if (req.body.data && req.body.data["rules"] && req.body.data["rules"].length > 0) {
			var rules = req.body.data["rules"];
			for (var i = 0; i < rules.length; i++) {
				if (!rules[i]["countryCode"])
					errorArray.push("Country Name is required for rule-" + (i + 1) + ".");

				if (rules[i]["countryCode"] && !territories.countries[rules[i]["countryCode"]])
					errorArray.push("Country Name is invalid for rule-" + (i + 1) + ".");

				if (rules[i]["provinceCode"] && !territories.provinces[rules[i]["provinceCode"]])
					errorArray.push("Province Name is invalid for rule-" + (i + 1) + ".");

				if (rules[i]["cityCode"] && !territories.cities[rules[i]["cityCode"]])
					errorArray.push("City Name is invalid for rule-" + (i + 1) + ".");

				if (!rules[i]["permissionType"])
					errorArray.push("Permission Type is required for rule-" + (i + 1) + ".");

				if (rules[i]["permissionType"] && ["INCLUDE", "EXCLUDE"].indexOf(rules[i]["permissionType"]) < 0)
					errorArray.push("Permission Type is invalid for rule-" + (i + 1) + ".");

				if(superDistributor) {
					if (rules[i]["countryCode"] && !territories.countries[rules[i]["countryCode"]] && !superDistributor.countries[rules[i]["countryCode"]])
						errorArray.push("You don't have permission to this country rule-" + (i + 1) + ".");

					if (rules[i]["provinceCode"] && !territories.provinces[rules[i]["provinceCode"]] && !superDistributor.provinces[rules[i]["provinceCode"]])
						errorArray.push("You don't have permission to this province rule-" + (i + 1) + ".");

					if (rules[i]["cityCode"] && !territories.cities[rules[i]["cityCode"]] && !superDistributor.cities[rules[i]["cityCode"]])
						errorArray.push("You don't have permission to this city rule-" + (i + 1) + ".");
				}

				var ruleCount=0;
				for(var j=0;j<rules.length;j++) {
					if(rules[i]["countryCode"] == rules[j]["countryCode"] && rules[i]["provinceCode"] == rules[j]["provinceCode"] && rules[i]["cityCode"] == rules[j]["cityCode"] && rules[i]["permissionType"] == rules[j]["permissionType"]) {
						ruleCount++;
					}
				}

				if(ruleCount > 1 && errorArray.indexOf("Same rules exists multiple times.") < 0)
					errorArray.push("Same rules exists multiple times.");
			}
		}

		if (errorArray.length > 0)
			res.json({
				message : "failure",
				result : errorArray
			});
		else
			saveDistributors();
	}

	function saveDistributors() {
		if (!req.body.data.id) {
			(config.id++);
			req.body.data["id"] = config.id;
		}

		dataBase[req.body.data.id] = req.body.data;
		res.json({
			message : "success"
		});
	}

	function deleteDistributor() {
		if(dataBase[req.body.data.id]) {
			delete dataBase[req.body.data.id];
			res.json({
				message : "success"
			});
		} else
			res.json({
				message : "failure",
				result : ["Can't delete distributor not found."]
			});
	}

	if(req.body.actionVerb == "Save")
		validateDetails();
	else
		deleteDistributor();
};

exports.getPermittedLocations = function (dataBase, territories, distributorName) {
	var result = {
		countries : {},
		provinces : {},
		cities : {}
	},
	inculdeArray = [],
	exculdeArray = [];

	function getPermissionsDetails(distributorName) {
		for (var prop in dataBase) {
			if (dataBase[prop]["distributorName"] == distributorName) {
				var rules = dataBase[prop]["rules"];
				for (var i = 0; i < rules.length; i++) {
					if (rules[i]["permissionType"] == "INCLUDE")
						inculdeArray.push(rules[i]);
					else if (rules[i]["permissionType"] == "EXCLUDE")
						exculdeArray.push(rules[i]);
				}

				if (dataBase[prop]["superDistributor"])
					getPermissionsDetails(dataBase[prop]["superDistributor"]);
			}
		}
	}

	function includeDetails(rule) {
		if(rule["cityCode"]) {
			result.countries[rule["countryCode"]] = territories.countries[rule["countryCode"]];
			result.provinces[rule["provinceCode"]] = territories.provinces[rule["provinceCode"]];
			result.cities[rule["cityCode"]] = territories.cities[rule["cityCode"]];
		} else if(rule["provinceCode"]) {
			result.countries[rule["countryCode"]] = territories.countries[rule["countryCode"]];
			result.provinces[rule["provinceCode"]] = territories.provinces[rule["provinceCode"]];
			for (var prop in territories.cities)
				if (territories.cities[prop].indexOf(rule["provinceCode"]) == 0)
					result.cities[prop] = territories.cities[prop];
		} else if(rule["countryCode"]) {
			result.countries[rule["countryCode"]] = territories.countries[rule["countryCode"]];
			for (var prop in territories.provinces)
				if (territories.provinces[prop].indexOf(rule["countryCode"]) == 0)
					result.provinces[prop] = territories.provinces[prop];

			for (var prop in territories.cities)
				if (territories.cities[prop].indexOf(rule["countryCode"]) == 0)
					result.cities[prop] = territories.cities[prop];
		}
	}

	function exculdeDetails(rule) {
		if (rule["cityCode"]) {
			if (result.cities[rule["cityCode"]])
				delete result.cities[rule["cityCode"]];
		} else if (rule["provinceCode"]) {
			if (result.provinces[rule["provinceCode"]])
				delete result.provinces[rule["provinceCode"]];

			for (var prop in result.cities)
				if (result.cities[prop].indexOf(rule["provinceCode"]) == 0)
					delete result.cities[prop];
		} else if (rule["countryCode"]) {
			if (result.countries[rule["countryCode"]])
				delete result.countries[rule["countryCode"]];

			for (var prop in result.provinces)
				if (result.provinces[prop].indexOf(rule["countryCode"]) == 0)
					delete result.provinces[prop];

			for (var prop in result.cities)
				if (result.cities[prop].indexOf(rule["countryCode"]) == 0)
					delete result.cities[prop];
		}
	}

	getPermissionsDetails(distributorName);

	for (var i = 0; i < inculdeArray.length; i++)
		includeDetails(inculdeArray[i]);

	for (var i = 0; i < exculdeArray.length; i++)
		exculdeDetails(exculdeArray[i]);

	return result;
};

exports.checkPermission = function (dataBase, territories, req, res, next) {
	var errorArray = [],
	permissionError = [];
	if (!req.body.data)
		errorArray.push("Invalid details provided.");

	if (req.body.data && !req.body.data["distributorName"])
		errorArray.push("Distributor Name is required.");

	if (req.body.data && !req.body.data["countryCode"])
		errorArray.push("Country Name is required.");

	if (req.body.data && req.body.data["distributorName"]) {
		var itemFound = false;
		for (var prop in dataBase) {
			if (dataBase[prop]["distributorName"] == req.body.data["distributorName"]) {
				itemFound = true;
				break;
			}
		}

		if (itemFound) {
			var result = exports.getPermittedLocations(dataBase, territories, req.body.data["distributorName"]);
			if (req.body.data && req.body.data["countryCode"] && !territories.countries[req.body.data["countryCode"]])
				errorArray.push("Country Name is invalid.");

			if (req.body.data && req.body.data["countryCode"] && territories.countries[req.body.data["countryCode"]] && !result.countries[req.body.data["countryCode"]])
				permissionError.push("You don't have permission to this country.");

			if (req.body.data && req.body.data["provinceCode"] && !territories.provinces[req.body.data["provinceCode"]])
				errorArray.push("Province Name is invalid.");

			if (req.body.data && req.body.data["provinceCode"] && territories.provinces[req.body.data["provinceCode"]] && !result.provinces[req.body.data["provinceCode"]])
				permissionError.push("You don't have permission to this province.");

			if (req.body.data && req.body.data["cityCode"] && !territories.cities[req.body.data["cityCode"]])
				errorArray.push("Province Name is invalid.");

			if (req.body.data && req.body.data["cityCode"] && territories.cities[req.body.data["cityCode"]] && !result.cities[req.body.data["cityCode"]])
				permissionError.push("You don't have permission to this city.");
		} else
			errorArray.push("Distributor doesn't exists.");
	}

	if (errorArray.length > 0)
		res.json({
			message : "failure",
			result : errorArray
		});
	else
		res.json({
			message : "success",
			result : permissionError
		});
};
