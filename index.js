const Cinema = require('./core')
const cinema = new Cinema()
const db = {}; // maintaining the whole data as a json

/**
 * Initiates the program
 * Mutates the db object
 */
(async function mainLoader () {
  const input = await cinema.getMainMenuInputs()
  console.log(JSON.stringify(db))
  switch (input) {
    case 1:
      cinema.createDistributor(db)
        .then(mainLoader)
        .catch(e => {
          console.error(e.message)
          return mainLoader()
        })
      break
    case 2:
      console.log('Select a distributor: \n')
      Object.keys(db)
        .forEach(i => {
          console.log(`${i}\n`)
        })
      const choice = await cinema.getUserInputs()
      if (!db[choice].sub) {
        db[choice].sub = {}
      }
      cinema.createDistributor(db[choice].sub)
        .then(mainLoader)
        .catch(e => {
          console.log(e.message)
          return mainLoader()
        })
      break
    case 3:
      console.log('The distributors are\n')
      let dispTxt = ''
      const iterate = (db, ii) => {
        console.log(Object.keys(db))
        Object.keys(db)
          .forEach((i, idx) => {
            dispTxt += `${Array(ii).fill(' ').join('')}${idx + 1}.) ${i}\n`
            if (db.sub) {
              iterate(db.sub, ++ii)
            }
          })
      }
      iterate(db, 1)
      console.log(dispTxt)
      break
    default:
      break
  }
})()
