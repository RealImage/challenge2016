global.__base = __dirname + '/';

var express = require('express');
var index = require('./routes/index');
var http = require('http');
var bunyan = require('bunyan');
var app = express();
var log = bunyan.createLogger({name: 'distributor-permissions', level : 'info'});
var bodyParser = require('body-parser');

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.use(function(req, res, next) {
  req.log = log;
  next();
});

app.use('/', index);

app.set('port', '3000');
// catch 404 and forward to error handler

app.use(function(req, res, next) {
  var err = new Error('Not Found');
  err.status = 404;
  next(err);
});

app.use(function(err, req, res, next) {
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};
  req.log.info({err : err}, 'global exception handler :');
  res.sendStatus(err.status || 500);
});

http.createServer(app).on('error', function(err) {
  log.error(err);
  process.exit(1);
}).listen(app.get('port'), function() {
  log.info("distributor permissions server listening on port " + app.get('port') + ' in ' + app.get('env'));
});

module.exports = app;
