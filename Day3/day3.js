const inputLines = readLineByLine('input.txt');

let sumErrors = 0;
for (const line of inputLines) {
    sumErrors += getPriority(findError(line));
}

console.log(`[Part1] Sum of priorities: ${sumErrors}`);

let sumRecurring = 0;
for (let i = 0; i < inputLines.length; i += 3) {
    const recurring = findRecurringInGroup([
        inputLines[i],
        inputLines[i + 1],
        inputLines[i + 2],
    ]);

    sumRecurring += getPriority(recurring.charCodeAt(0));
}

console.log(`[Part2] Sum of all groups: ${sumRecurring}`);

function getPriority(letter) { return letter > 96 ? letter - 96 : letter - 38; }

function findError(compartment) {
    const occurence = new Set();
    const half = compartment.length / 2;

    for (let i = 0; i < compartment.length; i++) {
        const codeUnit = compartment.charCodeAt(i);
        if (i < half) {
            occurence.add(codeUnit);
        } else if (occurence.has(codeUnit)) {
            return codeUnit;
        }
    }

    throw "could not find any recurring letter";
}

function getUniqueAscii(text) {
    const uniqueLetters = new Set();
    for (const letter of text) {
        uniqueLetters.add(letter);
    }
    return uniqueLetters.values();
}

function findRecurringInGroup(group) {
    const occurences = new Map();

    for (const compartment of group) {
        const compartmentUnique = getUniqueAscii(compartment);
        for (const letter of compartmentUnique) {
            const occurence = occurences.get(letter);
            occurences.set(letter, (occurence ?? 0) + 1);
        }
    }

    return getThreeOccurenceLetter(occurences);
}

function getThreeOccurenceLetter(map) {
    for (const [key, value] of map.entries()) {
        if (value == 3) return key;
    }
    throw "could not find 3 times recurring letter in a group";
}

function readLineByLine(filePath) {
    const fs = require('fs');
    const allFileContents = fs.readFileSync(filePath, 'utf-8');
    const inputLines = allFileContents.split(/\r?\n/);

    return inputLines;
}