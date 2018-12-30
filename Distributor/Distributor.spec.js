var { expect } = require("chai");
const createWorld = require("../createWorld");
const distributor = require("./Distributor");
const distributors = require("../EntitiesNoted/Distributors");

describe("Create distributor tests", () => {
  let world = null;
  before(done => {
    createWorld().then(data => {
      world = data;
      done();
    });
  });

  describe("Create new distributor", () => {
    it("createDistributor should be true after successful creation", () => {
      const name = "d1";
      const includes = "IN,US,KA-IN".split(",").map(e => e.trim());
      const excludes = "CENAI-TN-IN".split(",").map(e => e.trim());
      expect(distributor.createDistributor(name, includes, excludes)).to.be
        .true;
    });

    it("createDistributor should be false after failed creation", () => {
      const name = "d1";
      const includes = "IN,US,KA-FOO-BAR".split(",").map(e => e.trim());
      const excludes = "CENAI-TN-IN".split(",").map(e => e.trim());
      expect(distributor.createDistributor(name, includes, excludes)).to.be
        .false;
    });

    it("After successful creation of a distributor the distributor list should have the new distributor", () => {
      const name = "d10";
      const includes = "IN,US,KA-IN".split(",").map(e => e.trim());
      const excludes = "CENAI-TN-IN".split(",").map(e => e.trim());
      expect(distributor.createDistributor(name, includes, excludes)).to.be
        .true;
      expect(distributors[name]).to.not.be.undefined;
    });
  });

  describe("List distributors", () => {
    it("Should have more than 0 distributors", () => {
      expect(
        Object.keys(distributor.listDistributors()).length
      ).to.be.greaterThan(0);
    });
  });

  describe("Check distributor", () => {
    it("Should return true if distributor is present", () => {
      expect(distributor.distributorPresent("d1")).to.be.true;
    });
    it("Should return false if distributor is notpresent", () => {
      expect(distributor.distributorPresent("foo-bar")).to.be.false;
    });
  });
  describe("Relate distributor", () => {
    before(() => {
      const name = "d2";
      const includes = "IN".split(",").map(e => e.trim());
      const excludes = "TN-IN".split(",").map(e => e.trim());
      distributor.createDistributor(name, includes, excludes);
    });
    it("should relate d2 and d1 as d2<d1", () => {
      const d2 = "d2";
      const d1 = "d1";
      expect(distributor.relateDistributors(d2, d1)).to.be.true;
    });
    it("should have parent of d2 as d1", () => {
      const d2 = "d2";
      const d1 = "d1";
      expect(distributors[d2].parent.name).to.be.equal("d1");
    });
  });
  describe("Query distributor", () => {
    it("CHICAGO-ILLINOIS-UNITEDSTATES should be YES for d1", () => {
      const name = "d1";
      const code = "CHIAO-IL-US";
      expect(distributor.queryDistributor(name, code)).to.equal("YES");
    });
    it("CHENNAI-TAMILNADU-INDIA should be NO for d1", () => {
      const name = "d1";
      const code = "CENAI-TN-IN";
      expect(distributor.queryDistributor(name, code)).to.equal("NO");
    });
  });
});
