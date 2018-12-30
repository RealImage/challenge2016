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
    if (includes.length === 0) {
      console.log(
        `The distributor ${name} must have access to some locations. None provided here.`
      );
      return false;
    }
    for (let i = 0; i < includes.length; i++) {
      if (!helper.doesCodeExist(includes[i])) {
        console.log(`The code ${includes[i]} does not exist`);
        return false;
      }
    }
    if (
      excludes.length > 0 &&
      (excludes.length === 1 && excludes[0].length > 0)
    ) {
      for (let i = 0; i < excludes.length; i++) {
        if (!helper.doesCodeExist(excludes[i])) {
          console.log(`The code ${excludes[i]} does not exist`);
          return false;
        }
      }
    }

    const dist = new DistributorClass(name);
    for (let i = 0; i < includes.length; i++)
      dist.addIncludes(helper.getObjFromSequence(includes[i]).code);
    for (let i = 0; i < excludes.length; i++) {
      if (helper.getObjFromSequence(excludes[i]))
        dist.addExcludes(helper.getObjFromSequence(excludes[i]).code);
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
      return 0;
    } else {
      for (let i = 0; i < keys.length; i++) {
        console.log(distributors[keys[i]]);
      }
      return distributors;
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
      console.log("NO");
      return "NO";
    }
    // check if all the code actually exist
    if (!helper.doesCodeExist(place)) {
      console.log(`${place} does not exist in system`);
      console.log("NO");
      return "NO";
    }

    const distributorObj = this.getDistributor(distributor);
    const includesObjects = Object.keys(distributorObj.includes);

    // check if code is <= includes of distributor
    if (!includesObjects.includes(place)) {
      let b = false;
      for (let i = 0; i < includesObjects.length; i++) {
        if (helper.isHierarchyCorrect(place, includesObjects[i])) b = true;
      }
      if (b === false) {
        console.log(
          "The place is not a direct or sub relation with distributor's includes"
        );
        console.log("NO");
        return "NO";
      }
    }

    // check the excludes
    let tmp = distributorObj;
    while (tmp != null) {
      if (place in tmp.excludes) {
        console.log(`Sorry the place is excluded by ${tmp.name}`);
        console.log("NO");
        return "NO";
      }
      tmp = tmp.parent;
    }

    console.log("YES");
    return "YES";
  }
}

// Singelton pattern
const distributor = new Distributor();

module.exports = distributor;
