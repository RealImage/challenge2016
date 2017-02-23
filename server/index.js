var express 		= require('express');
var path 			= require('path');
var bodyParser		= require('body-parser');

// loading routes
var index = require('./modules');

var app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.static(path.join(__dirname, 'client')));


// attaching routes to the application
app.use(index);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
    var err = new Error('Not Found');
    err.status = 404;
    next(err);
});




module.exports = app;
