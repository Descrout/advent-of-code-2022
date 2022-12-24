const inputLines = readLineByLine('input.txt');

const latestElfCalories = [];
let maxElfCalori = 0;

for (const line of inputLines) {
    if (line.length == 0) {
        const totalCalori = latestElfCalories.reduce((val, sum) => val + sum);
        if (totalCalori > maxElfCalori) {
            maxElfCalori = totalCalori;
        }
        while (latestElfCalories.length) latestElfCalories.pop();
        continue;
    }

    const calori = parseInt(line);
    latestElfCalories.push(calori);
}

console.log(`Most carried calories: ${maxElfCalori}`);

function readLineByLine(filePath) {
    const fs = require('fs');
    const allFileContents = fs.readFileSync(filePath, 'utf-8');
    const inputLines = allFileContents.split(/\r?\n/);

    return inputLines;
}