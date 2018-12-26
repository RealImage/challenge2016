const Countries = require("../EntitiesNoted/Countries");
const Provinces = require("../EntitiesNoted/Provinces");
const Cities = require("../EntitiesNoted/Cities");

const LineByLineReader = require("line-by-line");
/**
 * A set of helper methods which is needed throught the project
 */
class Helper {
  /**
   * Check if the hierarchy between city province and country is correct or not.
   * Works for all combinations
   * @param {String[]} ar The array of relationships between places
   */
  isHierarchyCorrect(ar) {
    if (ar.length === 0) return false;
    else if (ar.length === 1) {
      return this.getEntityObject(ar[0]) != null ? true : false;
    } else {
      for (let i = 0; i < ar.length - 1; i++) {
        let obj = this.getEntityObject(ar[i]);
        let parent = ar[i + 1];
        if (!this.isParent(obj, parent)) return false;
      }
      return true;
    }
  }

  /**
   * Check if the parent is actually a parent of the element
   * @param {EntityObject} element The entity object whose parent is to be tracked
   * @param {String} parent The code for the parent
   */
  isParent(element, parent) {
    let parentNode = element.parent;
    while (parentNode != null) {
      if (parentNode.code === parent) return true;
      parentNode = parentNode.parent;
    }
    return false;
  }

  /**
   * Returns the object of the code,, for example if the code is of
   * city then this returns the city object for that code
   * @param {String} code The code for the place
   */
  getEntityObject(code) {
    let obj = null;
    if (code in Cities) {
      obj = Cities[code];
    } else if (code in Provinces) {
      obj = Provinces[code];
    } else if (code in Countries) {
      obj = Countries[code];
    }
    return obj;
  }

  /**
   * Reads a file line by line and does some operation
   * @param {String} path The path to the file to be read
   * @param {Function} fn The function to be performed after reading the file line by line
   * @param {Number} fn Lines to skip from top
   */
  readFile(path, fn, skip = 0) {
    return new Promise((resolve, reject) => {
      let count = 0;
      const lr = new LineByLineReader(path);
      lr.on("error", function(err) {
        reject(err);
      });
      lr.on("line", function(line) {
        if (count >= skip) fn(line);
        count++;
      });
      lr.on("end", function() {
        resolve();
      });
    });
  }
}

const helper = new Helper();
module.exports = helper;
