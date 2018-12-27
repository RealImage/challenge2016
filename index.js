const chalk = require("chalk");
const createWorld = require("./createWorld");
const helper = require("./utils/helper");
const config = require("./config");
const logger = require("./utils/logger");

createWorld()
  .then(async world => {
    mainMenu();
  })
  .catch(err => {
    console.log(
      "There system encountered an error. Please restart the software."
    );
    console.log(err);
    logger.error(err.message);
  });

/**
 * The main menu for the program
 */
async function mainMenu() {
  console.log(config.main_menu);
  const input = await helper.getUserInput();
  if (!["0", "1", "2", "3", "4"].includes(input)) {
    console.log(chalk.red(config.menu_error));
    mainMenu();
  } else {
    console.log(chalk.green(`User selected ${input}`));
    switch (input) {
      case "0":
        process.exit(0);
      case "1":
        createDistributorMenu();
        break;
      case "2":
        relateDistributorMenu();
        break;
      case "3":
        listDistributorMenu();
        break;
      case "4":
        queryDistributorMenu();
        break;
    }
  }
}

/**
 * Add a new distributor into the system
 */
async function createNewDistributorWrapper() {
  return new Promise(async (resolve, reject) => {
    console.log(chalk.yellow(config.add_distributor));
    console.log(`Enter distributor name`);
    const name = await helper.getUserInput();
    console.log(`INCLUDES`);
    const includes = (await helper.getUserInput())
      .split(",")
      .map(e => e.trim());
    console.log(`EXCLUDES`);
    const excludes = (await helper.getUserInput())
      .split(",")
      .map(e => e.trim());

    console.log({ name, includes, excludes });
    resolve();
  });
}

/**
 * Menu for the creation of new distributor
 */
async function createDistributorMenu() {
  console.log(config.distributor_menu);
  const input = await helper.getUserInput();
  if (["0", "1"].includes(input)) {
    switch (input) {
      case "0":
        mainMenu();
        return;
      case "1":
        await createNewDistributorWrapper();
        mainMenu();
        break;
    }
  } else {
    console.log(chalk.red(config.menu_error));
    createDistributorMenu();
  }
}
async function relateDistributorMenu() {}
async function listDistributorMenu() {}
async function queryDistributorMenu() {}

/**
 * Ask user for
 * 1. Create distributor(with includes and excludes)
 * 2. Relate distrubutors
 * 3. List distributors
 * 4. Query distributor to entity code
 */
