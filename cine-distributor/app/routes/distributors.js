import Ember from 'ember';

export default Ember.Route.extend({
	queryParams: {
		distributorIds: {
            refreshModel: true
        }
    },
    distributorIds: null,
    model: function(params) {
    	if(params.distributorIds == null || params.distributorIds.trim() == ""){
    		this.transitionTo("/");
    	}

	    let distributors = this.store.peekAll('distributor');
	    let _this = this;
	    let currentDistributors = Ember.A();
	    let childDistributors = Ember.A();

	    distributors.filterBy("isActive", true).forEach(function(distributor){
	    	distributor.set("isActive", false);
	    });

	    params.distributorIds.split(",").forEach(function(distributorId){
	    	var distributor = _this.store.peekRecord('distributor', Number(distributorId));	    	
	    	let currentChildDistributors = Ember.A();

	    	distributor.set("isActive", true);
	    	currentDistributors.addObject(distributor);
	    	
	    	distributors.forEach(function(d){
	    		if(Number(d.get("parentDistributorId")) == Number(distributorId)){
	    			currentChildDistributors.addObject(d);
	    		}
	    	})
	    	childDistributors.addObject(currentChildDistributors);
	    });

	    return Ember.RSVP.hash({
            currentDistributors: currentDistributors,
            childDistributors: childDistributors
        });
	},
	setupController(controller, model) {
		controller.set('currentDistributors', model.currentDistributors);
		controller.set('childDistributors', model.childDistributors);
	},
	actions:{
		refresh: function() {
            this.refresh();
        }
	}
});
