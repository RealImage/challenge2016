exports.sendSuccessResponse = function (res, data) {
	res.json({
		status: true,
		data: data
	})
}

exports.sendErrorResponse = function (res, error) {
	console.log(error.stack);
	res.json({
		status: false,
		message : error.message
	});
}
