class DistributorClass {
  constructor(parent) {
    this.parent = parent;
    this.includes = [];
    this.excludes = [];
    this.children = [];
  }
  addIncludes(code) {
    this.includes.push(code);
  }
  addExcludes(code) {
    this.excludes.push(code);
  }
  getParent() {
    return this.parent;
  }
  addChild(distributor) {
    this.children.push(distributor);
  }
  listChildren() {
    return this.children;
  }
}

module.exports = DistributorClass;
