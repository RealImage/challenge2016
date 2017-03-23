/**
 * Created by vino on 3/22/17.
 */

var APIAction = require("./actions/action.js");
var eventObj = require("./data.json");
// var fs =require("fs");
var _ = require("underscore");

var APIRouts = function(app,conf){

    this.app = app;
    this.conf = conf;
    this.action = new APIAction(app,conf);

};
module.exports = APIRouts;


APIRouts.prototype.init = function(){
    var self=this

    self.app.post('/create', function (req, res) {
        req.query=req.body ?req.body : (req.payload ? req.payload : req.query)
            self.action.create(req,res)

    })
    self.app.get('/list', function (req, res) {
        req.query=req.body ?req.body : (req.payload ? req.payload : req.query)
            self.action.list(req,res)

    })

}
