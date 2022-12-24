import 'dart:io';

const scores = <String, int>{
  "A": 1,
  "B": 2,
  "C": 3,
  "X": 1,
  "Y": 2,
  "Z": 3,
};

// Some languages do negative number modulus different
// So I put "C" as first, for the wrap around
const symbols = <String>["C", "A", "B", "C", "A"];

int getMyScore(String other, String mine) {
  int myScore = scores[mine]!;
  final otherScore = scores[other]!;

  final diff = myScore - otherScore;
  if (diff == 0) {
    myScore += 3;
  } else if (diff == 1 || diff == -2) {
    myScore += 6;
  }

  return myScore;
}

int getMyScorePart2(String other, String mine) {
  final symbolIdx = scores[other]!;

  switch (mine) {
    case "X":
      return getMyScore(other, symbols[symbolIdx - 1]);
    case "Z":
      return getMyScore(other, symbols[symbolIdx + 1]);
    default:
      return getMyScore(other, other);
  }
}

void main() async {
  final inputFile = File("input.txt");
  final inputLines = await inputFile.readAsLines();

  int scorePart1 = 0;
  int scorePart2 = 0;
  for (final line in inputLines) {
    final splitted = line.split(" ");
    scorePart1 += getMyScore(splitted[0], splitted[1]);
    scorePart2 += getMyScorePart2(splitted[0], splitted[1]);
  }

  print("[Part1] Your score is: $scorePart1");
  print("[Part2] Your score is: $scorePart2");
}