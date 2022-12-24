const inputLines = readLineByLine('input.txt');

const scores = {
    "A": 1,
    "B": 2,
    "C": 3,
    "X": 1,
    "Y": 2,
    "Z": 3,
};

// Some languages do negative number modulus different
// So instead of % wrap around
// I put "C" as first and "A" as last
const symbols = ["C", "A", "B", "C", "A"];

let scorePart1 = 0;
let scorePart2 = 0;
for (const line of inputLines) {
    const splitted = line.split(" ");
    const other = splitted[0];
    const mine = splitted[1];
    scorePart1 += getMyScore(other, mine);
    scorePart2 += getMyScorePart2(other, mine);
}

console.log(`[Part1] Your score is: ${scorePart1}`);
console.log(`[Part2] Your score is: ${scorePart2}`);

function getMyScore(other, mine) {
    let myScore = scores[mine];
    const otherScore = scores[other];

    const diff = myScore - otherScore;
    if (diff == 0) {
        myScore += 3;
    } else if (diff == 1 || diff == -2) {
        myScore += 6;
    }

    return myScore;
}

function getMyScorePart2(other, mine) {
    const symbolIdx = scores[other];

    switch (mine) {
        case "X":
            return getMyScore(other, symbols[symbolIdx - 1]);
        case "Z":
            return getMyScore(other, symbols[symbolIdx + 1]);
        default:
            return getMyScore(other, other);
    }
}

function readLineByLine(filePath) {
    const fs = require('fs');
    const allFileContents = fs.readFileSync(filePath, 'utf-8');
    const inputLines = allFileContents.split(/\r?\n/);

    return inputLines;
}