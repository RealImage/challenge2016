var express 		= require('express');
var path 			= require('path');
var bodyParser		= require('body-parser');

// loading routes
var index = require('./server/modules');

var app = express();

// view engine setup
app.set('views', path.join(__dirname, 'server/views'));
app.set('view engine', 'ejs');

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.static(path.join(__dirname, 'client'), { redirect: false }));
app.use('/app/*', function (req, res) {
   res.sendFile(__dirname + '/client/index.html');
});

// attaching routes to the application
app.use(index);




module.exports = app;
