const { expect } = require("chai");
const Entity = require("./Entity");

describe("Entity tests", () => {
  it("object should be created", () => {
    const obj = new Entity("name");
    expect(obj).to.not.be.null;
    expect(obj.name).to.be.eq("name");
  });
});
