const DataProcessor = require('./DataProcessor')
const fs = require('fs')


const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
})



async function Stratprocess() {

    // console.log("Extracting data from csv")
    let rows = await DataProcessor.extract()
        .catch((err) => {
            console.log("caught")
        })

    // console.log("Loading Data into memory")
    let StructuredData = await DataProcessor.formatdata(rows).catch((err) => {
        console.error("Error Handled")
    })


    // console.log("Data is loaded and ready to take instructions")

    // console.log("Reading Input Fields From input.JSON")

    const permitedCities = []

    try {
        let FileData = fs.readFileSync('input.JSON')

        let permissionJSON = JSON.parse(FileData.toString())
        // console.log(` Distributors Name: ${permissionJSON.DistributorName} \n INCLUDE: ${permissionJSON.Include} \n Exclude: ${JSON.stringify(permissionJSON.Exclude)}`)

        for (let ind in permissionJSON.Include) {


            if (permissionJSON.Exclude.ExcludeCountries.includes(permissionJSON.Include[ind])) {
                break
            } else {

                for (let [provinceId, pValue] of StructuredData.get(permissionJSON.Include[ind]).get('Provinces')) {
                    if (permissionJSON.Exclude.ExcludeProvinces.includes(provinceId)) {
                        // console.log(`Skipping ${provinceId}`)
                        break
                    } else {
                        pValue.get('cities').map((city) => {
                            if (!permissionJSON.Exclude.ExcludeCities.includes(city.CityCode)) {

                                permitedCities.push(city.CityCode)
                            }

                        })

                    }

                }
            }

        }
        console.log(`Permited cities identified \n Number of permited cities ${permitedCities.length}`)
        // console.log(permitedCities)


        readline.question(`Enter a City Code to check permissions \n`, (city) => {
            console.log(`Checking Permissions for ${city}`)

            if (permitedCities.includes(city.trim())) {
                console.log(`${city} has permissions`)
                process.exit(0)
            } else {
                console.log(`${city} do not have permissions`)
                process.exit(0)
            }
            readline.close()
        })


    } catch (error) {
        console.log("err", error)
    }


}

Stratprocess()