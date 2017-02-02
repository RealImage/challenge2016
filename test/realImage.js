
var should = require('should'),
    request = require('supertest'),
    app = require('../../../server'),

    agent = request.agent(app);


describe('Distributors Rules Generation', function(req,res) {
    it('should generate rules', function(done) {
        agent.post('/rules')
            .send()
            .expect(201)
            .end(function(err, res) {
                if (err) return done(err);
                done();
            });
    })
});