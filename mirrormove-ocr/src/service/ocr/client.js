const tesseract = require('tesseract.js')

async function readImage(filepath) {
    return tesseract.recognize(filepath, 'eng');
}

exports.readImage = readImage;
