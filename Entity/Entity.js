class Entity {
  constructor(name, code, parent) {
    this.name = name;
    this.code = code;
    this.parent = parent;
    this.children = {};
  }

  getParent() {
    return this.parent;
  }
  getChildren() {
    return this.children;
  }
  getName() {
    return this.name;
  }
  getCode() {
    return this.code;
  }
  display() {
    console.log(`Name: ${this.name} Code: ${this.code}`);
  }
  addChild(code, entity) {
    this.children[code] = entity;
  }
}

module.exports = Entity;
