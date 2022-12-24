import 'dart:io';

void main() async {
  final inputFile = File("input.txt");
  final inputLines = await inputFile.readAsLines();

  final latestElfCalories = <int>[];
  int maxElfCalori = 0;

  for (final line in inputLines) {
    if (line.isEmpty) {
      final totalCalori = latestElfCalories.reduce((value, sum) => value + sum);
      if (totalCalori > maxElfCalori) {
        maxElfCalori = totalCalori;
      }
      latestElfCalories.clear();
      continue;
    }

    final calori = int.parse(line);
    latestElfCalories.add(calori);
  }

  print("Most carried calories: $maxElfCalori");
}
