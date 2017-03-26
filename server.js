var express = require('express');

var router = require('./routes/router');

var bodyParser = require('body-parser');

var app = express();

app.use(bodyParser.json());

app.use('/api', router);

app.use(express.static('public'));

app.listen(9090, serverStatus());

function serverStatus(){

	console.log('server running at http://localhost:9090');

}
