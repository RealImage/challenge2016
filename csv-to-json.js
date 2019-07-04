const csvFilePath = "/home/devc271/lakshmi/challenge2016/cities.csv";
const jsonFilePath = "/home/devc271/lakshmi/challenge2016/cities.json";

const csv = require("csvtojson");

const readStream = require("fs").createReadStream(csvFilePath);

const writeStream = require("fs").createWriteStream(jsonFilePath);

readStream.pipe(csv()).pipe(writeStream);
