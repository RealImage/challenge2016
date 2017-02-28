import DS from 'ember-data';

export default DS.JSONAPISerializer.extend({
	payloadKeyFromModelName: function(modelName) {
        return Ember.String.dasherize(modelName);
    },
    keyForAttribute: function(attr, method) {
        return Ember.String.camelize(attr);
    },
    serializeIntoHash: function(hash, type, record, options) {
        Ember.merge(hash, this.serialize(record, options));
    }
});
