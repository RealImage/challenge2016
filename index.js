const chalk = require("chalk");
const createWorld = require("./createWorld");
const helper = require("./utils/helper");
const config = require("./config");
const logger = require("./utils/logger");
const distributor = require("./Distributor/Distributor");

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
    logger.info(`User selected ${input} in main menu`);
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
async function createNewDistributor() {
  console.log(chalk.yellow(config.add_distributor));
  console.log(`Enter distributor name`);
  const name = await helper.getUserInput();
  console.log(`INCLUDES`);
  const includes = (await helper.getUserInput()).split(",").map(e => e.trim());
  console.log(`EXCLUDES`);
  const excludes = (await helper.getUserInput()).split(",").map(e => e.trim());
  logger.info(
    `User selected ${name}, ${includes}, ${excludes} in createNewDistributor menu`
  );

  const result = distributor.createDistributor(name, includes, excludes);
  if (result) {
    console.log("Success");
    logger.info(`Result was success`);
    mainMenu();
  } else {
    logger.info(`Result was failure`);
    console.log(
      "Sorry something went wrong, please make sure the codes match the codes from csv and the distributor does not exist already."
    );
    createNewDistributor();
  }
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
        createNewDistributor();
        break;
    }
  } else {
    console.log(chalk.red(config.menu_error));
    createDistributorMenu();
  }
}
async function relateDistributorMenu() {
  console.log(chalk.yellow(config.relate_distributor));
  console.log("Enter relationship");
  const input = await helper.getUserInput();
  const ar = input.split("<").map(a => a.trim());
  if (ar.length !== 2) {
    console.log("Please make sure the number of distributors is 2");
    relateDistributorMenu();
  } else {
    const result = distributor.relateDistributors(ar[0], ar[1]);
    if (!result) {
      console.log("Sorry the relation could not be established");
    }
    mainMenu();
  }
}
async function listDistributorMenu() {
  distributor.listDistributors();
  mainMenu();
}
async function queryDistributorMenu() {
  console.log(chalk.yellow(config.query_distributor));
  console.log("Enter distributor name");
  const distributorName = await helper.getUserInput();
  console.log("Enter place to query");
  const place = await helper.getUserInput();
  distributor.queryDistributor(distributorName, place);
  mainMenu();
}
