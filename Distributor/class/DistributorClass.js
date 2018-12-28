class DistributorClass {
  constructor(parent = null) {
    this.parent = parent;
    this.includes = {};
    this.excludes = {};
    this.children = [];
  }
  addIncludes(code) {
    this.includes[code] = true;
  }
  addExcludes(code) {
    delete this.excludes[code];
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
