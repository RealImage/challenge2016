import DS from 'ember-data';
import attr from 'ember-data/attr';

export default DS.Model.extend({
	cityCode: attr('string'),
	provinceCode: attr('string'),
	countryCode: attr('string'),
	city: attr('string'),
	province: attr('string'),
	country: attr('string'),
	formattedName: attr('string'),
	formattedCode: attr('string')
});
