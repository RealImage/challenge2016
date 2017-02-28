import DS from 'ember-data';
import ApplicationAdapter from './application';

export default ApplicationAdapter.extend({
	query: function(store, type, query) {
        let url = this.buildURL(type.modelName, null, null, 'query', query);
        url = url + "/" + query.id + "/" + query.type + "/" + query.code + "/" + query.name;
        
        return this.ajax(url, 'GET');
    }
});
