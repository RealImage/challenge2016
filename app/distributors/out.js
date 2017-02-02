/**
 * Created by Chethan H R on 2/2/2017.
 */
var _ = require('lodash');
var data = require('./rules.json');
exports.getData = function (distributorName, callback) {
    var rules = data.permisions;
    var input = distributorName;
    var out = [];

    function recursion(rule) {
        out.push(rule);
        if (rule.parent_id === "") {

            return [rule]
        }

        var current = _.filter(rules, {id: rule.parent_id});
        return current.map(x => {
            rule.parents = recursion(x);
            return x;
        });
    }

    var currentRules = _.filter(rules, {id: input});

    final = currentRules.map(rule => {
        rule.parents = recursion(rule);
        return rule;
    });

    return callback(null, _.uniq(out));

};
//console.log('indexed', indexed);
