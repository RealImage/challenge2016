import DS from 'ember-data';
import attr from 'ember-data/attr';

export default DS.Model.extend({
	name: attr('string'),
	parentDistributorId: attr('number'),
	includes: attr('array'),
	excludes: attr('array'),
	formattedIncludes: attr('array'),
	formattedExcludes: attr('array'),
	isActive: attr('boolean' , { defaultValue: false }),
	searchFor: attr('string' , { defaultValue: "City" }),
});
