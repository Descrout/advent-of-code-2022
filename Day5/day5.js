const inputLines = readLineByLine('input.txt');


class Procedure {
    constructor(txt) {
        const splitted = txt.split(" ");
        this.move = parseInt(splitted[1]);
        this.from = parseInt(splitted[3]) - 1;
        this.to = parseInt(splitted[5]) - 1;
    }
}

class Crates {
    constructor(size) {
        this.crates = [];
        for (let i = 0; i < size; i++) {
            this.crates[i] = [];
        }
    }

    copy() {
        const crates = new Crates(this.crates.length);
        for (let i = 0; i < this.crates.length; i++) {
            crates.crates[i] = [];
            for (let j = 0; j < this.crates[i].length; j++) {
                crates.crates[i].push(this.crates[i][j]);
            }
        }
        return crates;
    }

    fillCrates(line) {
        let j = 0;
        for (let i = 1; i < line.length; i += 4) {
            if (line[i] != " ") this.crates[j].unshift(line[i]);
            j++;
        }
    }

    applyProcedure(proc, is9001 = false) {
        const start = this.crates[proc.from].length - proc.move;
        const end = this.crates[proc.from].length;
        const sublist = this.crates[proc.from].slice(start, end);
        Array.prototype.push.apply(this.crates[proc.to], is9001 ? sublist : sublist.reverse());
        this.crates[proc.from].splice(start, end);
    }

    get topOfStacks() { return this.crates.map((e) => e[e.length - 1]).join(""); }
}


const procedures = [];
const crates = new Crates(9);

for (let i = 0; i < inputLines.length; i++) {
    const line = inputLines[i];

    if (i < 8) {
        crates.fillCrates(line);
        continue;
    }

    if (i > 9) procedures.push(new Procedure(line));
}

const crates9001 = crates.copy();

for (const proc of procedures) {
    crates.applyProcedure(proc);
    crates9001.applyProcedure(proc, true);
}

console.log(`[Part1] Top of each stack CrateMover9000: ${crates.topOfStacks}`);
console.log(`[Part2] Top of each stack CrateMover9001: ${crates9001.topOfStacks}`);

function readLineByLine(filePath) {
    const fs = require('fs');
    const allFileContents = fs.readFileSync(filePath, 'utf-8');
    const inputLines = allFileContents.split(/\r?\n/);

    return inputLines;
}