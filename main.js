const {
    set, get, listen
} = require('./store/store')

const {
    promptNewDistributor,
    promptQuestion
} = require('./util/prompt')

const initialize = async function () {
    try {
        await promptNewDistributor()
        console.log(`\nNow you can check whether distributor has access to a particular location\n`)
        await promptQuestion()
    } catch (error) {
        console.log(`Error while performing film distribution among the ditributors :P`, error)
    }
}

listen('data-loaded', (store) => {
    initialize()
})

process.once('SIGINT', function () {
    process.exit()
})