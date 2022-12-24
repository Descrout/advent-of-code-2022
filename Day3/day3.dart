import 'dart:io';

void main() async {
  final inputFile = File("input.txt");
  final inputLines = await inputFile.readAsLines();

  int sumErrors = getSumErrors(inputLines);

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
  final occurence = <int, int>{};

  for (final compartment in group) {
    putOccurences(getUniqueAscii(compartment), occurence);
  }

  int maxOccurence = 0;
  int? maxLetter;
  for (final kv in occurence.entries) {
    if (kv.value > maxOccurence) {
      maxOccurence = kv.value;
      maxLetter = kv.key;
    }
  }

  if (maxLetter == null) throw "could not find recurring letter in a group";
  return maxLetter;
}

void putOccurences(List<int> compartment, Map<int, int> occurences) {
  for (final letter in compartment) {
    final occurence = occurences[letter];
    if (occurence != null) {
      occurences[letter] = occurence + 1;
    } else {
      occurences[letter] = 1;
    }
  }
}

int getSumErrors(List<String> texts) => texts.fold(
      0,
      (previousValue, text) => previousValue + getPriority(findError(text)),
    );

List<int> getUniqueAscii(String text) => text.codeUnits.fold(
      <int>{},
      (previousValue, letterCode) => previousValue..add(letterCode),
    ).toList();

int getPriority(int letter) {
  if (letter > 96)
    return letter - 96;
  else
    return letter - 38;
}

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
