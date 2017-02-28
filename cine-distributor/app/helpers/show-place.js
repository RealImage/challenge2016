import Ember from 'ember';

export function showPlace(params/*, hash*/) {
    let place = params[0];
    
    let placeParts = place.trim().split(/\s*,\s*/);
    let placeType;
    console.log(placeParts)
    console.log(place)
    if(placeParts.length == 3){
        placeType = "City"
    }
    else if(placeParts.length == 2){
        placeType = "Province"
    }
    else{
        placeType = "Country"
    }
  
    return "<li class='mdl-list__item mdl-list__item--two-line'><span class='mdl-list__item-primary-content'><span class='place-name'>" + place + "</span><span class='mdl-list__item-sub-title'>" + placeType + "</span></span></li>"
    
}

export default Ember.Helper.helper(showPlace);
