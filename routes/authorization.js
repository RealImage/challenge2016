'use strict'
var _ = require('lodash');
var uuid = require('uuid/v1');

var authorization_api = require('../api/authorization');

var authorization = {

	add : function(req, res, next){

		var id = uuid(); 
		let auth = {
			[id]: {
				includes : {
					countries : (req.body.include_countries && req.body.include_countries.split(',')) || [],
					provinces : (req.body.include_provinves && req.body.include_provinves.split(',')) || [],
					cities : (req.body.include_cities && req.body.include_cities.split(',')) || []
				},
				excludes : {
					countries : req.body.exclude_countries || [],
					provinces : req.body.exclude_provinves || [],
					cities : req.body.exclude_cities || []
				}
			}			
		};
		req.log.info('Request recieved to add new authorization', {id : id, auth : JSON.stringify(auth)}); 
		
		authorization_api.add(auth, function(err, results){
			if (err){
				return res.sendStatus(500);
			}
			req.auth_id = id;
			next();
		});
	},

	verify : function(req, res, next){
		req.log.info({query : req.query}, 'Request recieved to verify permissions')
		

		let to_verify = {};

		if (req.distributor[req.distributor_id].auth.length == 0){
			return res.status(400).json({
				distributor_id : req.distributor_id,			
				distributor : req.distributor[req.distributor_id], 
				error : 'No valid authorization exists for the user, Please add authorization before verifying!'}
			);
		}

		if (req.query.city){
			to_verify.cities = req.query.city + '-' + req.query.province;
			if (!req.query.province){
				return res.status(400).json({err : 'Please send province as well, as there are same city code in diff provinces'});
			}
		}else if (req.query.province){
			if (!req.query.country){
				return res.status(400).json({err : 'Please send country as well, as there are same province code in diff countries'});
			}
			to_verify.provinces = req.query.province + '-' + req.query.country;
		}else if (req.query.country){
			to_verify.countries = req.query.country;
		}

		let auth = {
			authorized : req.distributor[req.distributor_id].auth,
			parent_authorizations : req.distributor[req.distributor_id].parent_auth,
			verify : to_verify
		}
		authorization_api.verify(auth, function(err, results){
			if (err){
				return res.status(500).json(err);
			}
			res.send(results);
		});
	},

	all : function(req, res, next){
		authorization_api.all(function(err, results){
			if (err){
				return res.sendStatus(500);
			}
			res.send(results);
		});
	},

	update : function(req, res, next){
		var id = req.body.id; 
		let auth = {
			[id]: {
				includes : {
					countries : (req.body.include_countries && req.body.include_countries.split(',')) || [],
					provinces : (req.body.include_provinves && req.body.include_provinves.split(',')) || [],
					cities : (req.body.include_cities && req.body.include_cities.split(',')) || []
				},
				excludes : {
					countries : (req.body.exclude_countries && req.body.exclude_countries.split(',')) ||  [],
					provinces : (req.body.exclude_provinces && req.body.exclude_provinces.split(',')) || [],
					cities : (req.body.exclude_cities && req.body.exclude_cities.split(',')) || []
				}
			}			
		};
		req.log.info('Request recieved to update new authorization', {id : id, auth : JSON.stringify(auth)}); 
		
		authorization_api.add(auth, function(err, results){
			if (err){
				return res.sendStatus(500);
			}
			res.status(200).json({result : 'Successfully updated authorization'});	
		});
	},

	get : function(req, res, next){
		let id = req.body.id;
		authorization_api.get(id, function(err, results){
			if (err){
				return res.send(500).json({err : err});
			}
			req.auth = results;
			next();
		});
	}


}

module.exports = authorization;