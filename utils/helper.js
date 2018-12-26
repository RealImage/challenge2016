const LineByLineReader = require("line-by-line");
/**
 * A set of helper methods which is needed throught the project
 */
class Helper {
  /**
   * Reads a file line by line and does some operation
   * @param {String} path The path to the file to be read
   * @param {Function} fn The function to be performed after reading the file line by line
   * @param {Number} fn Lines to skip from top
   */
  readFile(path, fn, skip = 0) {
    return new Promise((resolve, reject) => {
      let count = 0;
      const lr = new LineByLineReader(path);
      lr.on("error", function(err) {
        reject(err);
      });
      lr.on("line", function(line) {
        if (count >= skip) fn(line);
        count++;
      });
      lr.on("end", function() {
        resolve();
      });
    });
  }
}

const helper = new Helper();
module.exports = helper;
