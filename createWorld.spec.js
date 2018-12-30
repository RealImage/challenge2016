var { expect } = require("chai");
const createWorld = require("./createWorld");

describe("Testing of world object", () => {
  let world = null;
  before(done => {
    createWorld().then(data => {
      world = data;
      done();
    });
  });

  it("should create the world", () => {
    expect(world).to.not.be.null;
  });

  it("should create the world object", () => {
    expect(world.name).to.equal("World");
    expect(world.code).to.equal("WORLD");
    expect(world.parent).to.be.null;
    expect(Object.keys(world.children).length).to.be.greaterThan(0);
  });

  it("should have India as country", () => {
    expect(world.children["IN"]).to.not.be.null;
    expect(world.children["IN"].name).to.be.equal("India");
    expect(world.children["IN"].code).to.be.equal("IN");
    expect(Object.keys(world.children["IN"].children).length).to.be.greaterThan(
      0
    );
  });

  it("India should have Tamil Nadu", () => {
    expect(world.children["IN"].children["TN-IN"]).to.not.be.null;
    expect(world.children["IN"].children["TN-IN"].name).to.be.equal(
      "Tamil Nadu"
    );
    expect(world.children["IN"].children["TN-IN"].code).to.be.equal("TN-IN");
    expect(
      Object.keys(world.children["IN"].children["TN-IN"].children).length
    ).to.be.greaterThan(0);
  });

  it("Tamil Nadu should have Chennai", () => {
    expect(world.children["IN"].children["TN-IN"].children["CENAI-TN-IN"]).to
      .not.be.null;
    expect(
      world.children["IN"].children["TN-IN"].children["CENAI-TN-IN"].name
    ).to.be.equal("Chennai");
    expect(
      world.children["IN"].children["TN-IN"].children["CENAI-TN-IN"].code
    ).to.be.equal("CENAI-TN-IN");
    expect(
      Object.keys(
        world.children["IN"].children["TN-IN"].children["CENAI-TN-IN"].children
      ).length
    ).to.be.equal(0);
  });

  it("India should be able to access World through its parent", () => {
    expect(world.children["IN"].parent.name).to.be.equal("World");
    expect(world.children["IN"].parent.code).to.be.equal("WORLD");
  });
  after(() => {
    world = null;
  });
});
