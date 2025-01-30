const rgpokedb = require('rgpokedb');
const Fuse = require("fuse.js");
const { getHorde, getSingle, getMoves, getCaptcha, getLogs} = require('./ocr/pokemonOcr')


function textFuzzyCompare(texts, searchTerm, isCleaned, threshold) {
    if (!searchTerm) return
    const cleanTerm = isCleaned ? searchTerm : searchTerm.replace(/[^A-Za-z]+/g, '');
    const allText = texts.map(text => {return {text}})


    const options = {
        keys: ['text'],
        includeScore: true,
        threshold: threshold > 0 ? threshold : .35
    }

    const fuse = new Fuse(allText, options);
    const searchResults = fuse.search(cleanTerm);

    return !!searchResults[0];
}


async function isDesiredEncounter(expectedPokemon) {
    const pokemonName = await getSingle();
    return textFuzzyCompare(expectedPokemon, pokemonName)
}

async function isDesiredHorde(expectedPokemon) {
    const horde5 = await getHorde();
    return horde5.map((pkmn) => {
      return {
        id: pkmn.id,
        isDesired: textFuzzyCompare(expectedPokemon, pkmn.value),
      };
    });

}

async function isMoveUsable(id){
    let moves = await getMoves()
    const pattern = /(\d+)\/(\d+)/;
    try {
        const updatedMoves = moves.map(move => {
           
            const pp = move.value.match(pattern)[0];
            const pp_split = pp.split("/")
            console.log(pp_split)
            // if (Number(pp_split[0] == NaN)) throw new Error("Move pp not recognized")
            return  {id : move.id, value: Number(pp_split[0]) > 0}
        }).reduce((acc, { id, value }) => {
            acc[id] = value;
            return acc;
        }, {});

        return updatedMoves[id.toString()]
    } catch (e) {
        return null
    }
}

async function getLatestLog(){
    let log = await getLogs()

    let cleanLog = log.split('\n').map(_ => _.substring(8,).trim()).filter(_ => _)
    let newEncounterLog = []
    let i = cleanLog.length - 1;
    while (i >= 0) {
        let isBattleIntro = textFuzzyCompare(['Awild', 'A wild'], cleanLog[i], true, .8)
        newEncounterLog.unshift(cleanLog[i])
        if(isBattleIntro) {
            break;
        }
        i-- 
    }
    return newEncounterLog
}

async function isCaptchaVisible(){
    const captcha = await getCaptcha();
    let texts = ["Please solve the following CAPTCHA code to continue", "CAPTCHA"]
    let response = textFuzzyCompare(texts, captcha)
    return response;
}


async function logSearch(terms) {
    let log = await getLogs()
    let cleanLog = log?.split('\n')?.map(_ => _.substring(8,)?.trim())?.filter(_ => _) || null
    if (!cleanLog) return null
    let trimmedLog = []
    let i = cleanLog.length - 1;
    while (i >= 0) {
        let isDesiredTerm = textFuzzyCompare(terms, cleanLog[i], true, .8)
        trimmedLog.unshift(cleanLog[i])
        if( isDesiredTerm ) {
            break;
        }
        i-- 
    }
    return trimmedLog
}

async function getNewEncounterLogs() {
    const terms = ['Awild', 'A wild']
    const response = await logSearch(terms)
    return response
}

async function hasRanAway(logs){
    const terms = ['Got away safely!'];
    // const logs = await getNewEncounterLogs();
    let found = logs?.map((log) => textFuzzyCompare(terms, log, true)).find(i => i) || false;
    return found;
}


async function hasDefeated(logs){
    const terms = ['fainted'];
    // const logs = await getNewEncounterLogs();
    let found = logs?.map((log) => textFuzzyCompare(terms, log, true, .7)).find(i => i) || false;
    return found;
}

async function hasFrisked(logs){
    const terms = ['Banette risked', 'Banette frisked', 'found an item!'];
    // const logs = await getNewEncounterLogs();
    let found = logs?.map((log) => textFuzzyCompare(terms, log, true, .6)).find(i => i) || false;
    return found;
}



exports.isDesiredEncounter = isDesiredEncounter;
exports.isDesiredHorde = isDesiredHorde;
exports.isMoveUsable = isMoveUsable;
exports.getLatestLog = getLatestLog;
exports.isCaptchaVisible = isCaptchaVisible;
exports.hasRanAway = hasRanAway;
exports.getNewEncounterLogs = getNewEncounterLogs;
exports.hasDefeated = hasDefeated,
exports.hasFainted = null
exports.hasFrisked = hasFrisked;
