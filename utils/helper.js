const Countries = require("../EntitiesNoted/Countries");
const Provinces = require("../EntitiesNoted/Provinces");
const Cities = require("../EntitiesNoted/Cities");

const LineByLineReader = require("line-by-line");

const stdin = process.openStdin();

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
   * Get a user input from the user via command line
   */
  getUserInput() {
    return new Promise(resolve => {
      stdin.once("data", function(d) {
        return resolve(d.toString().trim());
      });
    });
  }

  /**
   * Check if the code for a place does actually exist in the system
   * or not
   * @param {String} code The code for a place
   */
  doesCodeExist(code) {
    return code in Cities || code in Provinces || code in Countries;
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
   *
   * @param {String[]} ar The sequence of hierarchy like [CHENI,TN,IN]
   */
  getObjFromSequence(ar) {
    switch (ar.length) {
      case 3:
        return Cities[ar[0]];
      case 2:
        return Provinces[ar[0]] || Cities[ar[0]];
      case 1:
        return Countries[ar[0]] || Provinces[ar[0]] || Cities[ar[0]];
      default:
        return {};
    }
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
