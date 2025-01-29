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
    const filepath = setDestination(title)

    makeFile(filepath, cropped);
    
}

function setDestination(title) {
    const fileName = `${title}.jpg`;
    if (title.includes("move")) { 
       return path.join(__dirname, `../../../mirrormove-ocr/src/screenshots/moves/${fileName}`)
    }
    return path.join(__dirname, `../../../mirrormove-ocr/src/screenshots/${title}/${fileName}`)
}


async function takePictures() {
    console.log(`${getDate()} | iteration ${i++}`)
    const filePath = path.join(__dirname, `../savedPositions/`)
    let fileNames = await getFileNames(filePath);
    await Promise.all(fileNames.map(async(fileName) => {
        const data = await getFile(`${filePath}/${fileName}`)
        cameraMan(JSON.parse(data))
    }))
}

function getDate() {    
    return dayjs(Date.now()).format('MM-DD-YYYY h:mm:ss A');
}


setInterval(takePictures, 5000)




