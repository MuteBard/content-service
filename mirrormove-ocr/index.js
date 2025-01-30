
// const { executeMove } = require('./src/runMove');
const {

    isDesiredEncounter,
    isDesiredHorde,
    isMoveUsable,
    isCaptchaVisible,
    hasRanAway,
    hasFrisked,
    hasDefeated,
    getNewEncounterLogs


} = require("./src/service/pokemonService");


async function getReport() {
    let desiredPokemon = ["Stunky"]
    let chosenMove = 4
    let logs = await getNewEncounterLogs()

    let report = {
        isfavorableE: await isDesiredEncounter(desiredPokemon),
        isfavorableHorde: await isDesiredHorde(desiredPokemon),
        isMoveUsable: await isMoveUsable(chosenMove),
        isCaptchaVisible: await isCaptchaVisible(),
        hasRanAway: await hasRanAway(logs),
        hasFrisked: await hasFrisked(logs),
        hasDefeated: await hasDefeated(logs)
    }
    console.log(report)
    return report
}

setInterval(getReport, 2500)


