'use strict'

var fs = require('fs');

var cities = fs.readFileSync('/Users/jaijin/clones/distributor-permissions/data/cities.json', 'utf8');

cities = JSON.parse(cities);

var cities_formatted = {}, provinces_formatted = {};

cities.forEach(function(city){
	cities_formatted[city.FIELD1 + '-' + city.FIELD2] = {
		city_code : city.FIELD1,
		province_code	: city.FIELD2,
		province_name	: city.FIELD5,
		country_code	: city.FIELD3,
		country_name	: city.FIELD6,
		city_name 		: city.FIELD4
	}; 

	provinces_formatted[city.FIELD2 + '-' + city.FIELD3] = {
		province_code : city.FIELD2,
		province_name	: city.FIELD5,
		country_code	: city.FIELD3,
		country_name	: city.FIELD6,
	};
});	

fs.writeFileSync('/Users/jaijin/clones/distributor-permissions/data/formatted_cities', JSON.stringify(cities_formatted));

fs.writeFileSync('/Users/jaijin/clones/distributor-permissions/data/formatted_provinces', JSON.stringify(provinces_formatted));

