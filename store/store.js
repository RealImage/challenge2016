const csv = require('csvtojson')
const _cliProgress = require('cli-progress')
const loader = new _cliProgress.SingleBar({}, _cliProgress.Presets.shades_classic);

const store = {
    distributors: [],
    meta: {},
    processedData: {},
    listeners: {}
}

const fetchFileAndProcess = async () => {
    const data = await csv()
        .fromFile('./data/cities.csv')
    const length = data.length
    loader.start(length, 0)    
    for( let i=0; i < length; i++) {
        const row = data[i]
        const cityName = row["City Name"].toLowerCase()
        const provinceName = row["Province Name"].toLowerCase()
        const countryName = row["Country Name"].toLowerCase()
        const cityCode = row["City Code"].toLowerCase()
        const provinceCode = row["Province Code"].toLowerCase()
        const countryCode = row["Country Code"].toLowerCase()
        
        store.processedData[`${cityName}-${provinceName}-${countryName}`] = `${cityCode}-${provinceCode}-${countryCode}`
        store.processedData[`${provinceName}-${countryName}`] = `${provinceCode}-${countryCode}`
        store.processedData[`${countryName}`] = `${countryCode}`

        if (!store.meta[`${countryName}`]) {
            store.meta[`${countryName}`] = {}
        }
        if (!store.meta[`${countryName}`][`${provinceName}`]) {
            store.meta[`${countryName}`][`${provinceName}`] = []
        }
        store.meta[`${countryName}`][`${provinceName}`].push(cityName)
        loader.update(i)
    }
    trigger('data-loaded')
    loader.stop()    
}

fetchFileAndProcess()

const set = (key, value) => {

}

const get = (key) => {
    return store[key]
}

const listen = (eventName, callback) => {
    store.listeners[eventName] = callback
}

const trigger = (eventName) => {
    if (store.listeners[eventName]) {
        store.listeners[eventName](store)
    }
}

module.exports = {
    set,
    get,
    listen,
    trigger
}

