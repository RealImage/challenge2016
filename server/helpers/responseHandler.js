exports.sendSuccessResponse = function (res, data) {
	res.json({
		status: true,
		data: data
	})
}

exports.sendErrorResponse = function (res, error) {
	res.json({
		status: false,
		message : error.message
	});
}
