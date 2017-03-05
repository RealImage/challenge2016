var express = require('express');
var router = express.Router();
var distributor = require('./distributor');
var authorization = require('./authorization');

router.post('/distributor/add', distributor.get_parent, distributor.add);

router.get('/distributor/verify_auth', distributor.get, authorization.verify);

router.get('/distributor/all', distributor.all);

router.post('/authorization/add', distributor.get, authorization.add, distributor.add_authorization);

router.post('/authorization/update', authorization.get, authorization.update);

router.get('/authorization/all', authorization.all);


module.exports = router;
