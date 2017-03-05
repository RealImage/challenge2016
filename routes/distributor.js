'use strict'

var _ = require('lodash');
var distributor_api = require('../api/distributor');
var uuid = require('uuid/v1');

var distributor = {

	add : function(req, res, next){
		let id = uuid();
		var distributor = {
			[id] : {
				name : _.get(req, 'body.name'),
				auth : [],
				parent_auth : req.parent_distributor && 
					(req.parent_distributor[req.body.parent].parent_auth.concat(req.parent_distributor[req.body.parent].auth)) || [],
			}						
		};

		distributor_api.add(distributor, function(err, results){
			if (err){
				return res.sendStatus(500);
			}
			res.send(results);
		});

	},

	get : function(req, res, next){		
		req.log.info ({id : req.query.id}, 'request recieved to get distributor details for :');

		req.distributor_id = req.query.id || req.body.id;
		distributor_api.get(req.distributor_id, function(err, results){
			if (err){
				return res.sendStatus(500);
			}else if (!results){
				return res.send(404).json({result : 'Distributor information not found'});
			}
			req.log.info (results, 'successfully found distributor');
			req.distributor = results;
			next();
		});
	},

	get_parent : function(req, res, next){
		let parent = _.get(req, 'body.parent');
		if (!parent){
			return next();
		}
		distributor_api.get(parent, function(err, results){
			if (err){
				return res.status(err.status).json({error : err.msg});
			}
			req.log.info (results, 'successfully found parent');
			req.parent_distributor = results;
			next();
		});
	},

	add_authorization : function(req, res, next){
		let distributor = req.distributor;
		if (distributor[req.distributor_id].auth.length !=0){
			return res.status(400).json({error : 'Authorization already exists! Please update existing Authorization', id :distributor[req.distributor_id].auth});
		}
		distributor[req.distributor_id].auth.push(req.auth_id);
		distributor_api.add_authorization(distributor, function(err, results){
			if (err){
				return res.sendStatus(500);
			}
			res.send(results);
		});
	},

	all : function(req, res, next){
		distributor_api.all(function(err, results){
			if (err){
				return res.send(500).json({error : err});
			}else if (!results){
				res.status(404).json({result :'No information found for distributors'});
			}
			res.send(results);
		});
	}
}

module.exports = distributor;