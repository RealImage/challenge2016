var express = require('express');
var router = express.Router();

var Proxy = require('./distributors.server.proxy.js');

//gets the list of distributors
router.get('/api/distributors/getAllDistributors/',  Proxy.getAllDistributors);

// adds distributor
router.get('/api/distributors/addDistributor/',  Proxy.addDistributor);


module.exports = router;
