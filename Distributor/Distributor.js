const Cities = require("../EntitiesNoted/Cities");
const Provinces = require("../EntitiesNoted/Provinces");
const Countries = require("../EntitiesNoted/Countries");
const helper = require("../utils/helper");
const distributors = require("../EntitiesNoted/Distributors");
const DistributorClass = require("./class/DistributorClass");
class Distributor {
  /**
   * Create a new Distributor
   * @param {String} name The name of the distributor
   * @param {String[]} includes The list of places to include
   * @param {String[]} excludes The list of places to exclude
   * @returns {Boolean}
   */
  createDistributor(name, includes, excludes) {
    if (name in distributors) return false;
    for (let i = 0; i < includes.length; i++)
      if (
        helper.doesCodeExist(includes[i]) === false ||
        !helper.isHierarchyCorrect(includes[i].split("-"))
      )
        return false;

    for (let i = 0; i < excludes.length; i++)
      if (
        !helper.doesCodeExist(excludes[i]) ||
        !helper.isHierarchyCorrect(excludes[i])
      )
        return false;

    const dist = new DistributorClass();
    for (let i = 0; i < includes.length; i++)
      dist.addIncludes(helper.getObjFromSequence(includes[i].split("-")));
    for (let i = 0; i < excludes.length; i++)
      dist.addIncludes(helper.getObjFromSequence(excludes[i].split("-")));

    distributors[name] = dist;
    return true;
  }

  /**
   * List all the available distributors in the system
   */
  listDistributors() {
    const keys = Object.keys(distributors);
    if (keys.length === 0) {
      console.log("Sorry, there are no distributors in the system yet.");
    } else {
      for (let i = 0; i < keys.length; i++) {
        console.log(distributors[keys[i]]);
      }
    }
  }

  /**
   * Check if a distributor is present in the system
   * @param {String} name The name of the distributor
   */
  distributorPresent(name) {
    return name in distributors;
  }

  /**
   * Return distributor object
   * @param {String} name Get distributor object from his name
   */
  getDistributor(name) {
    return distributors[name];
  }
}

const distributor = new Distributor();

module.exports = distributor;
