const { writeFile, readFile, listFiles, deleteFile } = require('./fileUtil');

async function makeFile(path, content) {
    await writeFile(path, content);
}

async function getFile(path) {
    return readFile(path);
}

async function getFileNames(path) {
    return listFiles(path)
}

async function removeFile(path) {
    return deleteFile(path)
}

exports.makeFile = makeFile;
exports.getFile = getFile;
exports.getFileNames = getFileNames;
exports.removeFile = removeFile;