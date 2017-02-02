/*jshint esversion: 6 */
const express = require('express'),
    UserController = require('./controllers/qubeChallenge.controller'),
    RealImage = require('./controllers/realImage.controller'),
    path = require('path');


router = express.Router();

router.post('/', function(req, res){
    res.json({ message: 'try POST: http://localhost:9000/rules & POST: http://localhost:9000/getDistributionStatus ' })
});
router.post('/rules', RealImage.rules);
router.post('/getDistributionStatus', RealImage.getDistributionStatus);

module.exports = router;