import Ember from 'ember';

export function ifLength(params/*, hash*/) {
    let v1 = params[0];
    let operator = params[1];
    let v2 = params[2];

    if(!v1){
        return false;
    }else{
    	v1 = v1.length;
    }

    switch (operator) {
        case '==':
            return (v1 == v2) ? true : false;
        case '===':
            return (v1 === v2) ? true : false;
        case '!=':
            return (v1 != v2) ? true : false;
        case '!==':
            return (v1 !== v2) ? true : false;
        case '<':
            return (v1 < v2) ? true : false;
        case '<=':
            return (v1 <= v2) ? true : false;
        case '>':
            return (v1 > v2) ? true : false;
        case '>=':
            return (v1 >= v2) ? true : false;
        case '&&':
            return (v1 && v2) ? true : false;
        case '||':
            return (v1 || v2) ? true : false;
        default:
            return false;
    }
}

export default Ember.Helper.helper(ifLength);
