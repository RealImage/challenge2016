'use strict'

var _ = require('lodash'),
	fs = require('fs'),
	authorizations = fs.readFileSync(global.__base + 'data/authorizations.js', 'utf8'),
	cities = JSON.parse(fs.readFileSync(global.__base + 'data/formatted_cities', 'utf8')),
	provinces = JSON.parse(fs.readFileSync(global.__base + 'data/formatted_provinces', 'utf8'));
	
try{	
	authorizations = JSON.parse(authorizations);
}catch(e){
	authorizations = {};
}

var authorization = {
	verify : function(auth, callback){
		
		var result = {
				allowed : 'Yes'	
		}, verification = true,
		available_authorizations = [],
		queried_for = {};

		if (auth.verify.cities){
			let city = cities[auth.verify.cities];
			queried_for.city_code = auth.verify.cities;
			if (!city){
				return callback({result : 'City Not found', queried_for, queried_for});
			}
			auth.verify.cities = city.city_code;
			auth.verify.countries = city.country_code;
			auth.verify.provinces = city.province_code;	
			queried_for.details = city;			
		}else if (auth.verify.provinces){
			let province = provinces[auth.verify.provinces];
			queried_for.province_code = auth.verify.provinces;
			if (!province){
				return callback({result : 'Province Not found', queried_for : queried_for});
			}
			auth.verify.provinces = province.province_code;
			auth.verify.countries = province.country_code;	
			queried_for.details = province;	
		}else{
			queried_for = {country : auth.verify.countries}
		}

		_.forEach(auth.parent_authorizations, function (permission){
			available_authorizations.push(authorizations[permission]);
		});		

		_.forEach(auth.authorized, function (permission){
			available_authorizations.push(authorizations[permission]);
		});		

		_.forEach(Object.keys(auth.verify), function (type){
			verify(type, auth.verify[type]);
		});


		function verify(type, value){
			(available_authorizations).forEach(function(auth){
				let includes = (auth.includes[type].length == 0) || check_includes (auth.includes[type], value);
				let excludes = (auth.excludes[type].length == 0) || check_excludes (auth.excludes[type], value);

				if (!(includes && excludes)){
					result.allowed = 'No'
					return;
				} 
			});

			function check_includes (available_values, value){
				return (available_values.indexOf(value) != -1)? true : false;				
			}

			function check_excludes (available_values, value){
				return (available_values.indexOf(value) != -1)?false : true;
			}			
		}
		result.queried_for = queried_for;
		return callback(null, result);
	},

	add : function(auth, callback){
		_.extend(authorizations, auth);
		fs.writeFileSync(global.__base + 'data/authorizations.js', JSON.stringify(authorizations));
		callback(null, {id : auth.id, result : 'Added successfully'});
	},

	all : function(callback){
		callback(null, {results : authorizations});
	},

	get : function(id, callback){
		let auth = authorizations[id];
		if (!auth){
			let err = {err : 'Authorization Not found! Please verify the id'};
			return callback(err);
		}
		callback(null, auth);
	},

	update : function(callback){

	}
};

module.exports = authorization;