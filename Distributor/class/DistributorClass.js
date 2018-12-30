class DistributorClass {
  constructor(name, parent = null) {
    this.name = name;
    this.parent = parent;
    this.includes = {};
    this.excludes = {};
    this.children = [];
  }
  /**
   * Add a code to the includes array
   * @param {String} code The code for the place
   */
  addIncludes(code) {
    this.includes[code] = true;
    delete this.excludes[code];
  }
  /**
   * Add a code to the excludes array
   * @param {String} code The code for the place
   */
  addExcludes(code) {
    delete this.includes[code];
    this.excludes[code] = true;
  }
  /**
   * Get parent director
   */
  getParent() {
    return this.parent;
  }
  /**
   * Add director to children
   * @param {Object} distributor Add director to the children array
   */
  addChild(distributor) {
    this.children.push(distributor);
  }
  /**
   * @returns {Object[]} The list of children
   */
  listChildren() {
    return this.children;
  }
}

module.exports = DistributorClass;
