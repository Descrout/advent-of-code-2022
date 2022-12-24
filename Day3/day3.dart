import 'dart:io';

void main() async {
  final inputFile = File("input.txt");
  final inputLines = await inputFile.readAsLines();

  final sumErrors = inputLines.fold(
    0,
    (prev, text) => prev + getPriority(findError(text)),
  );

  print("[Part1] Sum of priorities: $sumErrors");

  int sumRecurring = 0;
  for (int i = 0; i < inputLines.length; i += 3) {
    final recurring = findRecurringInGroup([
      inputLines[i],
      inputLines[i + 1],
      inputLines[i + 2],
    ]);

    sumRecurring += getPriority(recurring);
  }

  print("[Part2] Sum of all groups: $sumRecurring");
}

int findRecurringInGroup(List<String> group) {
  final occurences = <int, int>{};
  for (final compartment in group) {
    final compartmentUnique = getUniqueAscii(compartment);
    for (final letter in compartmentUnique) {
      final occurence = occurences[letter];
      occurences[letter] = (occurence ?? 0) + 1;
    }
  }

  return getThreeOccurenceLetter(occurences);
}

int getThreeOccurenceLetter(Map<int, int> map) {
  for (final kv in map.entries) {
    if (kv.value == 3) {
      return kv.key;
    }
  }
  throw "could not find 3 times recurring letter in a group";
}

List<int> getUniqueAscii(String text) => text.codeUnits.fold(
      <int>{},
      (prev, letterCode) => prev..add(letterCode),
    ).toList();

int getPriority(int letter) => letter > 96 ? letter - 96 : letter - 38;

int findError(String compartment) {
  final occurence = <int>{};
  final asciiList = compartment.codeUnits;
  final half = asciiList.length / 2;

  for (int i = 0; i < asciiList.length; i++) {
    final codeUnit = asciiList[i];
    if (i < half) {
      occurence.add(codeUnit);
    } else if (occurence.contains(codeUnit)) {
      return codeUnit;
    }
  }

  throw "could not find any recurring letter";
}
