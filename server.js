/**
 * Created by vino on 3/21/17.
 */
var express = require('express')
var app = express()


var ENV = process.env.ENV ? process.env.ENV : 'dev';
var envconf = require("./conf/conf-" + ENV + ".js");
var data = require('./data.json');

_ = require("lodash");
require("underscore-query")(_);




app.listen(envconf.web.port, function () {
    console.log('app listening on port'+envconf.web.port+' !')
})

var APIRouts = require("./routs.js");
var Routs = new APIRouts(app,envconf);
Routs.init();



// var csv2json = require('csv2json');
// var fs = require('fs');
//
// fs.createReadStream('cities.csv')
//     .pipe(csv2json({
//         // Defaults to comma.
//         // separator: ';'
//     }))
//     .pipe(fs.createWriteStream('data.json'));

