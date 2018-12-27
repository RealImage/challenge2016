const path = require("path");
const helper = require("./utils/helper");

const Entity = require("./Entity/Entity");

const countries = require("./EntitiesNoted/Countries");
const provinces = require("./EntitiesNoted/Provinces");
const cities = require("./EntitiesNoted/Cities");
const World = require("./EntitiesNoted/World");

const csvPath = path.join(__dirname, "cities.csv");

/**
 * Do some operation on the line of the file being read
 * @param {String} line One line in the file being read
 */
function processLine(line) {
  const parts = line.split(",");
  const cityCode = parts[0];
  const provinceCode = parts[1];
  const countryCode = parts[2];
  const cityName = parts[3];
  const provinceName = parts[4];
  const countryName = parts[5];

  let countryObj = null;
  let provinceObj = null;
  let cityObj = null;

  countryObj = createNewEntry(
    countries,
    countryCode,
    Entity,
    countryName,
    World
  );

  provinceObj = createNewEntry(
    provinces,
    provinceCode,
    Entity,
    provinceName,
    countryObj
  );

  cityObj = createNewEntry(cities, cityCode, Entity, cityName, provinceObj);
}

/**
 * Create a new object and mark it in memory or retrieve the object from
 * list of marked objects
 * @param {Object} history The noted hashmap for entity say country or province or city
 * @param {String} code The string code for a country name
 * @param {Class} ClassName The class who is going to create object
 * @param {String} name The name of the place (say country name or state name)
 * @param {Object} parent The parent entity of the entity
 */
function createNewEntry(history, code, ClassName, name, parent) {
  if (!history[code]) {
    history[code] = new ClassName(name, code, parent);
    parent.children[code] = history[code];
  }
  return history[code];
}

module.exports = () => {
  return new Promise((resolve, reject) => {
    helper
      .readFile(csvPath, processLine, 1)
      .then(() => {
        resolve(World);
      })
      .catch(err => {
        reject(err);
      });
  });
};
