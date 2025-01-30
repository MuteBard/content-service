const screenshot = require('screenshot-desktop');
const fs = require('fs');
const sharp = require('sharp');
const { getFileNames, getFile, makeFile, removeFile } = require('./fileManager')
const path = require("path");
const dayjs = require('dayjs')
let i = 1;

async function cameraMan({ x, y, width, height, title}) {
    const left = x
    const top = y

    const picture = await screenshot();
    const cropped = await sharp(picture).extract({ left, top, width, height}).toBuffer();
    const fileName = `${title}.jpg`;
    const filepath = path.join(__dirname, `../../../mirrormove-ocr/src/screenshots/${fileName}`)
    makeFile(filepath, cropped);
    
}

async function takePictures() {
    console.log(`${getDate()} | iteration ${i++}`)
    const filePath = path.join(__dirname, `../savedPositions/`)
    let fileNames = await getFileNames(filePath);
    // console.log(fileNames)
    await Promise.all(fileNames.map(async(fileName) => {
        const data = await getFile(`${filePath}/${fileName}`)
        cameraMan(JSON.parse(data))
    }))
}

function getDate() {    
    return dayjs(Date.now()).format('MM-DD-YYYY h:mm:ss A');
}


setInterval(takePictures, 2500)




