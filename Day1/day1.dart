import 'dart:io';

const topElfCaloriesLength = 3;
final topThreeElfCalories = <int>[0, 0, 0];
final latestElfCalories = <int>[];

void main() async {
  final inputFile = File("input.txt");
  final inputLines = await inputFile.readAsLines();

  for (final line in inputLines) {
    if (line.isEmpty) {
      final totalCalori = latestElfCalories.reduce((value, sum) => value + sum);
      updateTopThree(totalCalori);
      latestElfCalories.clear();
      continue;
    }

    final calori = int.parse(line);
    latestElfCalories.add(calori);
  }

  print("[Part1] Most carried calories: ${topThreeElfCalories[0]}");
  final topThreeSum = topThreeElfCalories.reduce((value, sum) => value + sum);
  print("[Part2] Top three calories sum: $topThreeSum");
}

void updateTopThree(int newCalories) {
  for (int i = 0; i < topElfCaloriesLength; i++) {
    if (newCalories > topThreeElfCalories[i]) {
      final oldCalori = topThreeElfCalories[i];
      topThreeElfCalories[i] = newCalories;
      updateTopThree(oldCalori);
      break;
    }
  }
}
