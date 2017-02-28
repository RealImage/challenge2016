import Ember from 'ember';

export default Ember.Component.extend({
	store: Ember.inject.service('store'),

	actions :{
		goToDistributor(distributorId){
			let distributorIds = [];
			let store = this.get('store');

			let distributor = store.peekRecord("distributor", distributorId);
			distributorIds.push(distributor.id);

			while(distributor.get("parentDistributorId") != 0){
				distributor = store.peekRecord("distributor", distributor.get("parentDistributorId"));
				distributorIds.unshift(distributor.id);				
			}

			this.get('router').transitionTo('distributors', { queryParams: { distributorIds: distributorIds } });
		}
	}
});
