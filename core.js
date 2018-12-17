const csv = require('csvtojson')
const citiesFile = './cities.csv'

var stdin = process.openStdin()

const getUserInputs = () => {
  return new Promise((resolve, reject) => {
    stdin.once('data', function (d) {
      return resolve(d.toString().trim())
    })
  })
}

const getJsonFromCsv = () => {
  return csv()
    .fromFile(citiesFile)
}

class Cinema {
  getUserInputs () {
    return new Promise((resolve, reject) => {
      stdin.once('data', function (d) {
        return resolve(d.toString().trim())
      })
    })
  }
  /**
   * List the menu and get inputs from the user
   */
  async getMainMenuInputs () {
    console.log('1. Create a Distributor\n2.Create a Sub Distributor\n3.View all Distributors')
    const input = await getUserInputs()
    const choice = parseInt(input, 10)
    if ([1, 2, 3].indexOf(choice) > -1) {
      return choice
    } else {
      console.error('Invalid Input')
      return this.getMainMenuInputs()
    }
  }

  async includeRegions (regions = {}, db) {
    const cities = await getJsonFromCsv()
    const qns = 'Enter a Country? [press enter]'
    const excludeRegion = async () => {
      console.log('1.Exclude Region\n2.Go Back')
      const choice = await getUserInputs()
      switch (Number(choice)) {
        case 1:
          console.log(`Enter City Code to exclude in ${country}: (Please refer the location data)`)
          const exRegion = await getUserInputs()
          const region = cities.filter(i => (i['Country Name'] === country))
            .filter(i => (i['City Code'] === exRegion))
          if (region.length === 0) {
            console.log(`Please enter a valid code for ${country}`)
            return excludeRegion()
          }
          if (regions[country] && regions[country].status === 'included') {
            if (!regions[country].excludedRegions) {
              regions[country].excludedRegions = []
            }
            regions[country].excludedRegions.push(exRegion)
            regions[country] = Array.from(new Set(regions[country]).excludedRegions)
          } else if (regions[country] && regions[country].status === 'excluded') {
            console.log(`Currently distributor has no permission over ${country}`)
            return
          } else {
            regions[country] = {}
            regions[country].status = 'included'
            regions[country].excludedRegions = [exRegion]
          }
          return excludeRegion()
        default:
      }
    }
    console.log(qns)
    const country = await getUserInputs()
    console.log(db)
    if (db && db[country].excludedRegions && db[country].excludedRegions.length > 0) {
      regions[country] = {}
      regions[country].status = 'included'
      regions[country].excludedRegions = db[country].excludedRegions
    }
    return excludeRegion()
  }

  async excludeRegions (regions, db) {
    const cities = await getJsonFromCsv()
    const qns = 'Enter a Country? [press enter]'
    const includeRegion = async () => {
      console.log('1.include RegiIon\n2.Go Back')
      const choice = await getUserInputs()
      switch (Number(choice)) {
        case 1:
          console.log(`Enter City Code to include in ${country}: (Please refer the location data)`)
          const exRegion = await getUserInputs()
          const region = cities.filter(i => (i['Country Name'] === country))
            .filter(i => (i['City Code'] === exRegion))
          if (region.length === 0) {
            console.log(`Please enter a valid code for ${country}`)
            return includeRegion()
          }
          if (regions[country] && regions[country].status === 'included') {
            if (!regions[country].includedRegions) {
              regions[country].includedRegions = []
            }
            regions[country].includedRegions.push(exRegion)
            regions[country] = Array.from(new Set(regions[country]).includedRegions)
          } else if (regions[country] && regions[country].status === 'included') {
            console.log(`Currently distributor has full permission over ${country}`)
            return
          } else {
            regions[country] = {}
            regions[country].status = 'included'
            regions[country].includedRegions = [exRegion]
          }
          return includeRegion()
        default:
      }
    }
    console.log(qns)
    const country = await getUserInputs()
    if (db && db[country].excludedRegions && db[country].excludedRegions.length > 0) {
      regions[country] = {}
      regions[country].status = 'included'
      regions[country].excludedRegions = db[country].excludedRegions
    }
    return includeRegion()
  }

  /**
   * This method will create a distributor with regions
   * to include and exclude movie distribution
   * @param {Object} db
   */
  async createDistributor (db) {
    console.log(`Enter distributor's id? `)
    const id = await getUserInputs()
    if (db[id]) {
      throw new Error('Distributor already exists')
    }
    console.log(`Enter distributor's name? `)
    const name = await getUserInputs()
    db[id] = {
      id,
      name,
      regions: {}
    }
    const addRegions = async () => {
      console.log(`1. Include a Region\n2.Exclude a Region\n3.Main Menu`)
      const input = await getUserInputs()
      switch (Number(input)) {
        case 1:
          await this.includeRegions(db[id].regions, db.regions)
          console.log(JSON.stringify(db))
          return addRegions()
        case 2:
          await this.excludeRegions(db[id].regions, db.regions)
          return addRegions()
        default:
      }
    }
    return addRegions()
  }
}

module.exports = Cinema
