/**
 * Created by vino on 3/21/17.
 */

var _ = require("underscore");
var conf = {};
var ENV =

    "dev";
//"test";
//"production";

var envconf = require("./conf/conf-" + ENV + ".js");
console.log(envconf);
//var sys_props = require('./conf/system-properties.js');
//conf = _.extend(envconf, sys_props);
//conf.project = require('./conf/project-properties.js');

conf = _.extend(envconf);

module.exports = conf;