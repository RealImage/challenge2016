const LineByLineReader = require("line-by-line");
/**
 * A set of helper methods which is needed throught the project
 */
class Helper {
  /**
   * Reads a file line by line and does some operation
   * @param {String} path The path to the file to be read
   * @param {*} fn The function to be performed after reading the file line by line
   */
  readFile(path, fn) {
    return new Promise((resolve, reject) => {
      const lr = new LineByLineReader(path);
      lr.on("error", function(err) {
        reject(err);
      });
      lr.on("line", function(line) {
        fn(line);
      });
      lr.on("end", function() {
        resolve();
      });
    });
  }
}

const helper = new Helper();
module.exports = helper;
