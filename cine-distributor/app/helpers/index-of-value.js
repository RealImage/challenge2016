import Ember from 'ember';

export function indexOfValue(params/*, hash*/) {
	return params[0][params[1]].get(params[2]);
}

export default Ember.Helper.helper(indexOfValue);
