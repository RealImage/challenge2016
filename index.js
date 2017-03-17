/**
 * Created by hmspl on 16/3/17.
 */
'use strict';

const Hapi = require('hapi');

const distributorCtrl = require('./lib/distributor');

// Create a server with a host and port
const server = new Hapi.Server();
server.connection({
    host: 'localhost',
    port: 8000
});

// Add the route
server.route({
    method : 'POST',
    path   : '/distributor/{distributorId}',
    handler: function (request, reply) {
        distributorCtrl.findDistributionPermit(request, function (err, response) {
            if (err) {
                return reply(err);
            } else {
                return reply(response);
            }
        });
    }
});

// Start the server
server.start((err) => {

    if (err) {
        throw err;
    }
    console.log('Server running at:', server.info.uri);
});