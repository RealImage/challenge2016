const express = require('express');
const csvFilePath = 'cities.csv'
const csv = require('csvtojson')
const bodyParser = require('body-parser')
const path = require('path');
const app = express();

app.use(bodyParser.json())
const PORT = 3000;

const geographyC = require('./geopgraphy.js')

const distributorC = require('./distributer.js');
const { copyFileSync } = require('fs');

console.log("distributrs", distributorC)


let parsedCsv = []
let geographyTree = {}
let distributiontTree = {}
csv()
    .fromFile(csvFilePath)
    .then((jsonObj) => {
        parsedCsv = jsonObj
        //  console.log(parsedCsv);


        for (let i = 0; i < parsedCsv.length; i++) {

            const cityName = parsedCsv[i]['City Name']
            const provinceName = parsedCsv[i]['Province Name']
            const countryName = parsedCsv[i]['Country Name']
            if (!geographyTree[countryName]) {
                geographyTree[countryName] = {
                    'includes': [],
                    'excludes': [],
                    'childs': {
                    }
                }


            }

            if (!geographyTree[countryName].childs[provinceName]) {
                geographyTree[countryName].childs[provinceName] = {
                    'includes': [],
                    'excludes': [],
                    'childs': {
                    }
                }

            }

            if (!geographyTree[countryName].childs[provinceName].childs[cityName]) {
                geographyTree[countryName].childs[provinceName].childs[cityName] = {
                    'includes': [],
                    'excludes': [],
                    'childs': null
                }

            }

        }

        //  console.log(geographyTree)



    })


app.post('/api/distributer', async (req, res) => {

    try {
        const distributer = req.body.distributer;  //  dist1<dist2
        const includes = req.body.includes;  //['karnataka', 'india']
        const excludes = req.body.excludes; //['iii-dddd', 'yyyyy-india']

        const distributers = distributer.split('<');
        console.log(distributer)

        const validationDistribution = distributorC.validateDistributers(distributers, distributers.length - 1, distributiontTree)

        if (!validationDistribution.status) {
            return res.status(400).json({ status: false, message: validationDistribution.message })

        }

        //console.log(geographyTree[])

        const validationGeography = geographyC.validateGeography(distributers, geographyTree, includes, excludes)


        if (!validationGeography.status) {

            return res.status(400).json({ status: false, message: validationGeography.message })
        }

        distributorC.insertDistributer(distributers, distributers.length - 1, distributiontTree, includes, excludes);

        geographyC.insertIncluesDistributorGeographyTree(distributers, includes, geographyTree)

        geographyC.insertExcludedDistributorGeographyTree(distributers, excludes, geographyTree)

        res.json({ status: true, message: "Successfully created" })


    } catch (err) {
        return res.status(500).json({ status: false, message: err.message })
    }


});



app.post('/api/distributer/fetch', (req, res) => {


    const distributer = req.body.distributer;
    const geography = req.body.geography;

    const distributerArray = distributer.split("<");


    const geographuStatus=geographyC.ifGeographyExist(distributerArray, geographyTree, geography)

    if(geographuStatus.status){
        res.send('YES')
    }else{
        res.send('NO')
    }





});






app.listen(PORT, (error) => {
    if (!error)
        console.log("Server is Successfully Running, and App is listening on port " + PORT)
    else
        console.log("Error occurred, server can't start", error);
}
);