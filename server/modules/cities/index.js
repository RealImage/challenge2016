var express = require('express');
var router = express.Router();

var Proxy = require('./cities.server.proxy.js');

//gets the list of cities
router.get('/api/cities/getCities/',  Proxy.getCities);

// gets the list of all countries
router.get('/api/cities/getAllCountries/',  Proxy.getAllCountries);


module.exports = router;
