import Ember from 'ember';

export default Ember.Route.extend({
	model: function() {
        return Ember.RSVP.hash({
            distributors: this.store.findAll('distributor'),
            places: this.store.peekAll('place')
        });
	},
	setupController(controller, model) {
		controller.set('distributors', model.distributors);
		controller.set('places', model.places);
		controller.set('searchType', "City");
		controller.set('distributor', this.store.createRecord('distributor'));		
	},	
	actions: {
		showDistributorDialog: function(parentDistributorId, index) {
			if(typeof(parentDistributorId) == "object"){
				parentDistributorId = parentDistributorId[index].id
			}

			this.controller.get('distributor').parentDistributorId = parentDistributorId;
			var dialog = document.querySelector('#distributorDialog');
			dialog.showModal();
		},
		hideDistributorDialog: function() {
			var dialog = document.querySelector('#distributorDialog');
			dialog.close();	
		},
		saveDistributor: function(distributor) {
			const flashMessages = Ember.get(this, 'flashMessages');
			var _this = this;
			distributor.save().then(function() {
				var dialog = document.querySelector('#distributorDialog');
				dialog.close();	
				flashMessages.success('Distributor created');
				_this.controller.set('distributor', _this.store.createRecord('distributor'))
				_this.controllerFor("distributors").send('refresh');
            }).catch(function(){
                flashMessages.danger('Distributor creation failed');
            });
		},
		showPlaceDialog: function(parentDistributorId, index, type) {
			this.controller.set("operationType", type)
			if(typeof(parentDistributorId) == "object"){
				parentDistributorId = parentDistributorId[index].id
			}

			this.controller.get('distributor').parentDistributorId = parentDistributorId;

			var dialog = document.querySelector('#placeDialog');
			dialog.showModal();
		},
		hidePlaceDialog: function() {
			var dialog = document.querySelector('#placeDialog');
			dialog.close();	
		},
		showPlaceTypeDropdown: function() {
			jQuery("#placeDialog .searchTypeMenu").toggleClass("hide");
		},
		setSearchType: function(type) {
			this.controller.set('searchType', type);
			jQuery("#placeDialog .searchTypeMenu").toggleClass("hide");
		},
		searchPlaces: function(param) {
			if(param){
				this.get('store').unloadAll('place');
				return this.get('store').query('place', {query: param, type: this.controller.get("searchType")});
			}	      
	    },
	    addPlace: function(code, name){
	    	this.get('store').unloadAll('place');
	    	this.controller.set('selectedPlaceCode', code)
	    	this.controller.set('selectedPlaceName', name)
	    	jQuery(".placeTypeInputBox").val(name);
	    },
	    savePlace: function(place){
	    	const flashMessages = Ember.get(this, 'flashMessages');
	    	let _this = this;
	    	if(this.controller.get("operationType") == "include"){
	    		this.get('store').query('distributor', {type:"include", id: this.controller.get("distributor").parentDistributorId, code: this.controller.get("selectedPlaceCode"), name: this.controller.get("selectedPlaceName")}).then(function() {
                    flashMessages.success('Added successfully');                    
                    _this.controllerFor("distributors").send('refresh');
                }).catch(function(reason) {
                    flashMessages.danger('Not allowed to save');
                });                
	    	}
	    	else{
	    		this.get('store').query('distributor', {type:"exclude", id: this.controller.get("distributor").parentDistributorId, code: this.controller.get("selectedPlaceCode"), name: this.controller.get("selectedPlaceName")}).then(function() {
                    flashMessages.success('Added successfully');
                    _this.controllerFor("distributors").send('refresh');
                }).catch(function(reason) {
                    flashMessages.danger('Not allowed to save');
                });
	    	}

	    	var dialog = document.querySelector('#placeDialog');
			dialog.close();	
	    }
	}
});
