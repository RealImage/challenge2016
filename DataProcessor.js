const csvParser = require('csv-parser')
const fs = require('fs')


class DataProcessor {


    extract() {

        var rows = []
        return new Promise((resolve, reject) => {

            fs.createReadStream('./cities.csv')
                .pipe(csvParser())
                .on('error', (err) => {
                    reject(err)
                })
                .on('data', (data) => rows.push(data))
                .on('end', () => {
                    // console.log(rows);
                    resolve(rows);
                });

        })

    }


    formatdata(rows) {

        let StructuredData = new Map()


        return new Promise((resolve, reject) => {
            try {

                rows.map((row) => {
                    let CountryCode = row['Country Code']
                    let ProvinceCode = row['Province Code']
                    let ProvinceName = row['Province Name']
                    let CityCode = row['City Code']
                    let CityName = row['City Name']
                    if (StructuredData.has(CountryCode)) {
                        // console.log("Found Country Code",StructuredData.get(CountryCode).get('CountryName'))
                        if (StructuredData.get(CountryCode).get('Provinces').has(ProvinceCode)) {

                            StructuredData.get(CountryCode).get('Provinces').get(ProvinceCode).get('cities').push({ 'CityCode': CityCode, 'CityName': CityName })
                        } else {
                            StructuredData.get(CountryCode).get('Provinces').set(ProvinceCode, new Map().set('ProvinceName', ProvinceName).set('cities', new Array({ 'CityCode': CityCode, 'CityName': CityName })))
                        }


                    } else {
                        // console.log("Adding new country")
                        StructuredData.set(row['Country Code'], new Map().set("CountryName", row['Country Name']).set("Provinces", new Map().set(ProvinceCode, new Map().set('ProvinceName', ProvinceName).set('cities', new Array({ 'CityCode': CityCode, 'CityName': CityName })))))
                    }

                })

                // console.log(StructuredData.get('IN').get('Provinces').get('OR').get('cities'));
                // console.log("Data is formated and arranged")
                resolve(StructuredData)
            } catch (error) {
                console.error("ERROR IN FORMATING DATA")
                reject(error)
            }

        })


    }

}

module.exports= new DataProcessor();