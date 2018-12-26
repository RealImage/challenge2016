const createWorld = require("./createWorld");
const helper = require("./utils/helper");

createWorld()
  .then(world => {
    console.log(helper.isHierarchyCorrect("TN-KA".split("-")));
    console.log(helper.doesCodeExist("KANA"));
  })
  .catch(err => {
    console.log(err);
  });

/**
 * Ask user for
 * 1. Create distributor(with includes and excludes)
 * 2. Relate distrubutors
 * 3. List distributors
 * 4. Query distributor to entity code
 */
