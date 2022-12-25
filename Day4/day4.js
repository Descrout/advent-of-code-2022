const inputLines = readLineByLine('input.txt');

class Pair {
    constructor(txt) {
        const commaSplit = txt.split(",");
        const first = commaSplit[0].split("-");
        const second = commaSplit[1].split("-");
        this.min1 = parseInt(first[0]);
        this.max1 = parseInt(first[1]);
        this.min2 = parseInt(second[0]);
        this.max2 = parseInt(second[1]);
    }

    get fullyContains() {
        return (this.min1 >= this.min2 && this.max1 <= this.max2) || (this.min2 >= this.min1 && this.max2 <= this.max1);
    }
    get overlaps() {
        return (this.min1 >= this.min2 && this.min1 <= this.max2) || (this.min2 >= this.min1 && this.min2 <= this.max1);
    }
}

const pairs = [];
for (const line of inputLines) {
    pairs.push(new Pair(line));
}

const fullyContains = pairs.filter((element) => element.fullyContains).length;
const overlaps = pairs.filter((element) => element.overlaps).length;
console.log(`[Part1] Fully contained pairs: ${fullyContains}`);
console.log(`[Part2] Overlapped pairs: ${overlaps}`);

function readLineByLine(filePath) {
    const fs = require('fs');
    const allFileContents = fs.readFileSync(filePath, 'utf-8');
    const inputLines = allFileContents.split(/\r?\n/);

    return inputLines;
}