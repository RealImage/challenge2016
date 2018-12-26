const path = require("path");
const helper = require("./utils/helper");

helper
  .readFile(path.join(__dirname, "cities.csv"), line => {
    console.log(line);
  })
  .then(() => {
    console.log("COMPLETE");
  })
  .catch(err => {
    console.log({ err });
  });
