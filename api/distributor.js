'use strict'

var _ = require('lodash'),
	fs = require('fs'),
	distributors = fs.readFileSync(global.__base + 'data/distributors.js', 'utf8');
try{	
	distributors = JSON.parse(distributors);
}catch(e){
	distributors = {};
}

var distributor = {

	add : function(distributor, callback){
		_.extend(distributors, distributor);
		fs.writeFileSync(global.__base + '/data/distributors.js', JSON.stringify(distributors));
		callback(null, {distributor : distributor, result : 'Added successfully, Please add authorization for the distributor'});
	},

	get : function(id, callback){
		if (!distributors[id]){
			let err = {status : 400, msg : 'Parent not found! Please check the parent id Or create distributor without passing parent info'}
			return callback(err, null);
		}
		callback(null, {[id] : distributors[id]});
	},

	add_authorization : function(distributor, callback){
		_.extend(distributors, distributor);
		fs.writeFileSync(global.__base + '/data/distributors.js', JSON.stringify(distributors));
		callback(null, {distributor : distributor, result : 'Added authorzation successfully'});
	},

	all : function(callback){
		callback(null, {results : distributors});
	}
};

module.exports = distributor;