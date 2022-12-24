const inputLines = readLineByLine('input.txt');

const latestElfCalories = [];
const topThreeElfCalories = [0, 0, 0];

for (const line of inputLines) {
    if (line.length == 0) {
        const totalCalori = latestElfCalories.reduce((val, sum) => val + sum);
        updateTopThree(totalCalori);
        while (latestElfCalories.length) latestElfCalories.pop();
        continue;
    }

    const calori = parseInt(line);
    latestElfCalories.push(calori);
}

console.log(`[Part1] Most carried calories: ${topThreeElfCalories[0]}`);
const topThreeSum = topThreeElfCalories.reduce((value, sum) => value + sum);
console.log(`[Part2] Top three calories sum: ${topThreeSum}`);

function updateTopThree(newCalories) {
    for (let i = 0; i < topThreeElfCalories.length; i++) {
        if (newCalories > topThreeElfCalories[i]) {
            const oldCalori = topThreeElfCalories[i];
            topThreeElfCalories[i] = newCalories;
            updateTopThree(oldCalori);
            break;
        }
    }
}

function readLineByLine(filePath) {
    const fs = require('fs');
    const allFileContents = fs.readFileSync(filePath, 'utf-8');
    const inputLines = allFileContents.split(/\r?\n/);

    return inputLines;
}