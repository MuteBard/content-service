

const { readImage } = require('./client')
const path = require("path");
let isTesting = false;
let testDir = isTesting ? "testing/" : "";

async function getHorde(){
    return Promise.all([...Array(5)].fill().map(async (_, i) => {
        const filepath = path.join(__dirname, `../../screenshots/${testDir}horde_h${i + 1}.jpg`)
        const res = await readImage(filepath);
        return  {id: i + 1, value: res?.data?.text} || null
    }))
}


async function getSingle(){
    const filepath = path.join(__dirname, `../../screenshots/${testDir}single_n.jpg`)
    const res = await readImage(filepath);
    return  res?.data?.text || null
}


async function getMoves(){
    return Promise.all([...Array(4)].fill().map(async (_, i) => {
        const filepath = path.join(__dirname, `../../screenshots/${testDir}moves_m${i + 1}.jpg`)
        const res = await readImage(filepath);
        return {id: i + 1, value: res?.data?.text} || null
    }))
}

async function getCaptcha(){
    const filepath = path.join(__dirname, `../../screenshots/${testDir}captcha.jpg`)
    const res = await readImage(filepath);
    return res?.data?.text || null
}

async function getLogs(){
    const filepath = path.join(__dirname, `../../screenshots/${testDir}logs.jpg`)
    const res = await readImage(filepath);
    return res?.data?.text || null
}


exports.getHorde = getHorde;
exports.getSingle = getSingle;
exports.getMoves = getMoves;
exports.getCaptcha = getCaptcha;
exports.getLogs = getLogs;
