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
    for (let i = 0; i < includes.length; i++) {
      let tmp = includes[i].split("-");
      for (let j = 0; j < tmp.length; j++) {
        if (!helper.doesCodeExist(tmp[j])) {
          return false;
        }
      }
      if (!helper.isHierarchyCorrect(includes[i].split("-"))) {
        return false;
      }
    }

    for (let i = 0; i < excludes.length; i++) {
      let tmp = excludes[i].split("-");
      for (let j = 0; j < tmp.length; j++) {
        if (!helper.doesCodeExist(tmp[i])) {
          //   console.log("here1");
          return false;
        }
      }
      if (!helper.isHierarchyCorrect(excludes[i].split("-"))) {
        {
          //   console.log("here2");
          return false;
        }
      }
    }

    const dist = new DistributorClass(name);
    for (let i = 0; i < includes.length; i++)
      dist.addIncludes(helper.getObjFromSequence(includes[i].split("-")).code);
    for (let i = 0; i < excludes.length; i++) {
      dist.addExcludes(helper.getObjFromSequence(excludes[i].split("-")).code);
    }

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

  /**
   * Create a relationship between two distributors
   * @param {String} d1 The name of distributor1
   * @param {String} d2 The name of distributor2
   */
  relateDistributors(d2, d1) {
    if (!this.distributorPresent(d1) || !this.distributorPresent(d2)) {
      console.log(`Either ${d2} or ${d1} does not exist`);
      return false;
    }
    const distrib1 = this.getDistributor(d1);
    const distrib2 = this.getDistributor(d2);

    // if distrib2 already has parent then return false
    if (distrib2.parent != null) {
      console.log(`${d2} already has a parent`);
      return false;
    }

    const includesObjectd1 = Object.keys(distrib1.includes);
    const includesObjectd2 = Object.keys(distrib2.includes);

    // each include of distributor2 <= distributor1 include
    for (let i = 0; i < includesObjectd2.length; i++) {
      if (!(includesObjectd2[i] in distrib1.includes)) {
        let b = false;
        for (let j = 0; j < includesObjectd1.length; j++) {
          if (
            helper.isHierarchyCorrect(includesObjectd2[i], includesObjectd1[j])
          )
            b = true;
        }
        if (b === false) {
          console.log("Each include of distributor2 <= distributor1 include");
          return false;
        }
      }
    }

    // each include of distributor2 not in exclude of distributor1 or its parents
    let tmp = distrib1;
    while (tmp != null) {
      for (let i = 0; i < includesObjectd2.length; i++) {
        if (includesObjectd2[i] in tmp.excludes) {
          return false;
        } else {
          let tmpAr = Object.keys(tmp.excludes);
          for (let j = 0; j < tmpAr.length; j++) {
            if (helper.isHierarchyCorrect(includesObjectd2[i], tmpAr[j])) {
              console.log(
                "Each include of distributor2 not in exclude of distributor1 or its parents"
              );
              return false;
            }
          }
        }
      }
      tmp = tmp.parent;
    } // while!

    distrib2.parent = distrib1;
    distrib1.children.push(distrib2);
    return true;
  }

  /**
   * Check if the distributor is allowed to sell in the place
   * @param {String} distributor The distributor name to query
   * @param {String} place The name of the place to query
   */
  queryDistributor(distributor, place) {
    if (!this.distributorPresent(distributor)) {
      console.log(`The distributor does not exist ${distributor}`);
      return;
    }
    // check if all the code actually exist
    let tmp = place.split("-");
    for (let j = 0; j < tmp.length; j++) {
      if (!helper.doesCodeExist(tmp[j])) {
        {
          console.log(`${tmp[j]} does not exist in system`);
          return;
        }
      }
    }
    // check if the hierarchy is correct
    if (!helper.isHierarchyCorrect(place.split("-"))) {
      console.log(`The hierarchy of code is not correct ${place}`);
      return;
    }

    const placeName = place.split("-");

    const distributorObj = this.getDistributor(distributor);
    const includesObjects = Object.keys(distributorObj.includes);
    console.log(includesObjects, placeName);
    // check if code is <= includes of distributor
    if (!includesObjects.includes(placeName[0])) {
      let b = false;
      for (let i = 0; i < includesObjects.length; i++) {
        if (helper.isHierarchyCorrect(placeName[0], includesObjects[i]))
          b = true;
      }
      if (b === false) {
        console.log(
          "The place is not a direct or sub relation with distributor's includes"
        );
        return;
      }
    }

    // check the excludes
    tmp = distributorObj;
    while (tmp != null) {
      if (placeName[0] in tmp.excludes) {
        console.log(`Sorry the place is excluded by ${tmp.name}`);
        return;
      }
      tmp = tmp.parent;
    }

    console.log("YES");
  }
}

// Singelton pattern
const distributor = new Distributor();

module.exports = distributor;
