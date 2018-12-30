const { expect } = require("chai");
const DistributorClass = require("./DistributorClass");

describe("DistributorClass tests", () => {
  it("object should be created", () => {
    const obj = new DistributorClass("name");
    expect(obj).to.not.be.null;
    expect(obj.name).to.be.eq("name");
  });
});
