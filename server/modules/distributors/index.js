var express = require('express');
var router = express.Router();

var Proxy = require('./distributors.server.proxy.js');

//gets the list of distributors
router.get('/api/distributors/getAllDistributors/',  Proxy.getAllDistributors);

// adds distributor
router.post('/api/distributors/addDistributor/',  Proxy.addDistributor);

// updates distributor
router.post('/api/distributors/updateDistributor/',  Proxy.updateDistributor);

router.get('/api/distributors/getDistributorById/:id',  Proxy.getDistributorById);

router.post('/api/distributors/saveSharedLocations',  Proxy.saveSharedLocations);



module.exports = router;
