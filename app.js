var express = require('express');
var path = require('path');
var bodyParser = require('body-parser');
//var logger = require('morgon');

var app = express();

//app.set('views', path.join(__dirname , 'public'));
//app.set('view engine', 'jade');

var dSelect = require(path.join(__dirname, 'routes', 'server.js'));

app.use(express.static(path.join(__dirname, 'public')));
app.use(express.static(path.join(__dirname, 'public' , 'view')));

//app.use(logger('dev'));

/** bodyParser.urlencoded(options)
 * Parses the text as URL encoded data (which is how browsers tend to send form data from regular forms set to POST)
 * and exposes the resulting object (containing the keys and values) on req.body
 */
app.use(bodyParser.urlencoded({
    extended: true
}));

/**bodyParser.json(options)
 * Parses the text as JSON and exposes the resulting object on req.body.
 */
app.use(bodyParser.json());

app.use('/', dSelect);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
  var err = new Error('Not Found');
  err.status = 404;
  next(err);
});

// error handler
// no stacktraces leaked to user unless in development environment
app.use(function(err, req, res, next) {
  res.status(err.status || 500);
  res.render('error', {
    message: err.message,
    error: (app.get('env') === 'development') ? err : {}
  });
});

/*app.listen(3000);
console.log("Running at port 3000");*/

module.exports = app;