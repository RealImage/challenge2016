var csvFilePath = './cities.csv'
var csv = require('csvtojson');
var readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
});
var dist1 = [],
    dist2 = [],
    dist3 = [];


//Get all inputs for three distributors.
console.log("INCLUDE/EXCLUDE regions for distributors. If no EXCLUDE just press enter.");
readline.question(`Enter Distributor 1 regions to be INCLUDED: `, function (distributor1_include) {
    readline.question(`Enter Distributor 1 regions to be EXCLUDED: `, function (distributor1_exclude) {
        readline.question(`Enter Distributor 2 regions to be INCLUDED: `, function (distributor2_include) {
            readline.question(`Enter Distributor 2 regions to be EXCLUDED: `, function (distributor2_exclude) {
                readline.question(`Enter Distributor 3 regions to be INCLUDED: `, function (distributor3_include) {
                    readline.question(`Enter Distributor 3 regions to be EXCLUDED: `, async function (distributor3_exclude) {
                        let ans = await getResult(distributor1_include, distributor1_exclude, distributor2_include, distributor2_exclude, distributor3_include, distributor3_exclude);
                        console.log('\n',"************* ", ans, " *************");
                        readline.close()
                    });
                });
            });
        });
    });
});

function getDistributorRegions(distRegions, excludedRegions, citiesList) {
    let distList = [],
        status = false;
    for (let j = 0; j < distRegions.length; j++) {
        for (let i = 0; i < citiesList.length; i++) {
            //conditions to check if the entered name is a country/state/city and it's not an excluded one
            if (citiesList[i].country_name.toLowerCase() == distRegions[j].trim() && excludedRegions.indexOf(citiesList[i].country_name.toLowerCase()) == -1) {
                distList.push(citiesList[i]);
                status = true;
            } else if (citiesList[i].province_name.toLowerCase() == distRegions[j].trim() && excludedRegions.indexOf(citiesList[i].province_name.toLowerCase()) == -1) {
                distList.push(citiesList[i]);
                status = true;
            } else if (citiesList[i].city_name.toLowerCase() == distRegions[j].trim() && excludedRegions.indexOf(citiesList[i].city_name.toLowerCase()) == -1) {
                distList.push(citiesList[i]);
                status = true;
            }
        }
    }
    return {
        distList: distList,
        status: status
    };
}

async function getResult(distributor1_include, distributor1_exclude, distributor2_include, distributor2_exclude, distributor3_include, distributor3_exclude) {
    let citiesList = await csv().fromFile(csvFilePath);
    let resp = "Successfully regions distributed!!!";
    distributor1_include = distributor1_include.toLowerCase().split(',');
    distributor2_include = distributor2_include.toLowerCase().split(',');
    distributor3_include = distributor3_include.toLowerCase().split(',');
    distributor1_exclude = distributor1_exclude.toLowerCase().split(',');
    distributor2_exclude = distributor2_exclude.toLowerCase().split(',');
    distributor3_exclude = distributor3_exclude.toLowerCase().split(',');
    dist1 = getDistributorRegions(distributor1_include, distributor1_exclude, citiesList);
    if (dist1.status) {
        dist2 = getDistributorRegions(distributor2_include, distributor2_exclude, dist1.distList);
        if (dist2.status) {
            dist3 = getDistributorRegions(distributor3_include, distributor3_exclude, dist2.distList);
            if (dist3.status) {} else {
                resp = "Sorry can't distribute regions for distributor3, because the selected region might be excluded.";
            }
        } else {
            resp = "Sorry can't distribute regions for distributor2, because the selected region might be excluded. So distributor3 also don't have any regions.";
        }
    } else {
        resp = "Sorry can't distribute regions for distributor1,  because there is no such region in our data set. So distributor2 and distributor3 also don't have any regions.";
    }
    return resp;
}