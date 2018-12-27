var expect = require("chai").expect;
describe("Array", function() {
  describe("#indexOf()", function() {
    it("should return -1 when the value is not present", function() {
      [1, 2, 3].indexOf(1).should.equal(0);
      [1, 2, 3].indexOf(2).should.equal(1);
    });
  });
});
