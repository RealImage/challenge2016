var express 		= require('express');
var router			= express.Router();
var cities			= require('./cities')
var distributors	= require('./distributors')


/* GET home page. */
router.use('/', cities);
router.use('/', distributors);

module.exports = router;
