const createWorld = require("./createWorld");

createWorld()
  .then(world => {
    console.log(world);
  })
  .catch(err => {
    console.log(err);
  });
