const path = require("path");
const helper = require("./utils/helper");

const Entity = require("./Entity/Entity");

const countries = require("./EntitiesNoted/Countries");
const provinces = require("./EntitiesNoted/Provinces");
const cities = require("./EntitiesNoted/Cities");
const World = require("./EntitiesNoted/World");

const csvPath = path.join(__dirname, "cities.csv");

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
