import 'dart:io';

void main() async {
  final inputFile = File("input.txt");
  final inputLines = await inputFile.readAsLines();

  final procedures = <Procedure>[];
  final crates = Crates.fromSize(9);

  for (int i = 0; i < inputLines.length; i++) {
    final line = inputLines[i];

    if (i < 8) {
      crates.fillCrates(line);
      continue;
    }

    if (i > 9) procedures.add(Procedure.fromStr(line));
  }

  final crates9001 = crates.copy();

  for (final proc in procedures) {
    crates.applyProcedure(proc);
    crates9001.applyProcedure(proc, true);
  }

  print("[Part1] Top of each stack CrateMover9000: ${crates.topOfStacks}");
  print("[Part2] Top of each stack CrateMover9001: ${crates9001.topOfStacks}");
}

class Crates {
  final List<List<String>> crates;
  Crates({
    required this.crates,
  });

  factory Crates.fromSize(int size) {
    return Crates(
      crates: List.generate(size, (index) => []),
    );
  }

  Crates copy() {
    return Crates(crates: crates.map((e) => e.map((e) => e).toList()).toList());
  }

  void fillCrates(String line) {
    int j = 0;
    for (int i = 1; i < line.length; i += 4) {
      if (line[i] != " ") crates[j].insert(0, line[i]);
      j++;
    }
  }

  void applyProcedure(Procedure proc, [bool is9001 = false]) {
    final start = crates[proc.from].length - proc.move;
    final end = crates[proc.from].length;
    final sublist = crates[proc.from].sublist(start, end);
    crates[proc.to].addAll(is9001 ? sublist : sublist.reversed);
    crates[proc.from].removeRange(start, end);
  }

  String get topOfStacks => crates.map((e) => e.last).join("");
}

class Procedure {
  final int move;
  final int from;
  final int to;

  Procedure({
    required this.move,
    required this.from,
    required this.to,
  });

  factory Procedure.fromStr(String txt) {
    final splitted = txt.split(" ");
    return Procedure(
      move: int.parse(splitted[1]),
      from: int.parse(splitted[3]) - 1,
      to: int.parse(splitted[5]) - 1,
    );
  }
}
