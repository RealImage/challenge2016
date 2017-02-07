var api = require('./api');
module.exports = function (app) {
	var dataBase = {},
	config = {
		id : 0
	},
	territories = {
		countries : {},
		provinces : {},
		cities : {}
	};

	api['getTerritories'](function (err, result) {
		if(err) {
			console.log("******************************Territory File Read Error***********************************");
			console.log(result);
		} else
			territories = result;
	});

	app.get('/', function (req, res, next) {
		res.render('main');
	});

	app.get('/territories', function (req, res, next) {
		res.json({
			result : territories
		});
	});

	app.get('/distributors', function (req, res, next) {
		var result = [];
		for (var prop in dataBase)
			result.push({
				id : prop,
				distributorName : dataBase[prop]["distributorName"],
				territories : api['getPermittedLocations'](dataBase, territories, dataBase[prop]["distributorName"])
			});

		res.json({
			result : result
		});
	});

	app.get('/distributors/:id', function (req, res, next) {
		if (dataBase[Number(req.params.id)])
			res.json({
				message : 'success',
				result : dataBase[Number(req.params.id)]
			});
		else
			res.json({
				message : 'failure',
				result : ["Distributor does not exists."]
			});
	});

	app.post('/distributors', function (req, res, next) {
		api['distributors'](dataBase, config, territories, req, res, next);
	});

	app.post('/checkPermission', function (req, res, next) {
		api['checkPermission'](dataBase, territories, req, res, next);
	});

	app.get('*', function (req, res) {
		console.log('*');
		res.render('main');
	});
};
