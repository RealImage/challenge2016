/*jshint esversion: 6 */
'use strict';
const express = require('express'),
    bodyParser = require('body-parser'),
    routes = require('../app/routes'),
    path = require('path'),
    config = require('./env'),
    app = express();

// Middleware to require login/auth
app.use(bodyParser.json());  // jshint ignore:line
app.use(bodyParser.urlencoded({ extended: true }));


app.all('/*', function (req, res, next) {
    // CORS headers
    res.header("Access-Control-Allow-Origin", "*");
   /* res.header('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE,OPTIONS');
    res.header('Access-Control-Allow-Headers', 'Content-type,Accept,mimeType');*/
    if (req.method == 'OPTIONS') {
        res.status(200).end();
    } else {
        next();
    }
});


app.use('/', routes);


// catch 404
app.use((req, res, next) => {
    return res.status(404).send({
        "message": "API Not Found",
        "statusCode": "404",
        "data": null
    });
});


app.use(function (err, req, res, next) {
    console.error(err.stack)
    // db logger comes here
    next();
    //res.status(500).send('Something broke!')
})
app.set('x-powered-by', false);

module.exports = app;