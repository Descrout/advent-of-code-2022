// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:io';

class Pair {
  final int min1;
  final int max1;
  final int min2;
  final int max2;
  Pair({
    required this.min1,
    required this.max1,
    required this.min2,
    required this.max2,
  });

  bool get fullyContains =>
      (min1 >= min2 && max1 <= max2) || (min2 >= min1 && max2 <= max1);

  bool get overlaps =>
      (min1 >= min2 && min1 <= max2) || (min2 >= min1 && min2 <= max1);

  factory Pair.fromStr(String txt) {
    final commaSplit = txt.split(",");
    final first = commaSplit[0].split("-");
    final second = commaSplit[1].split("-");
    return Pair(
      min1: int.parse(first[0]),
      max1: int.parse(first[1]),
      min2: int.parse(second[0]),
      max2: int.parse(second[1]),
    );
  }

  @override
  String toString() {
    return 'Pair(min1: $min1, max1: $max1, min2: $min2, max2: $max2)';
  }
}

void main() async {
  final inputFile = File("input.txt");
  final inputLines = await inputFile.readAsLines();

  final pairs = <Pair>[];
  for (final line in inputLines) {
    pairs.add(Pair.fromStr(line));
  }

  final fullyContains = pairs.where((element) => element.fullyContains).length;
  final overlaps = pairs.where((element) => element.overlaps).length;
  print("[Part1] Fully contained pairs: $fullyContains");
  print("[Part2] Overlapped pairs: $overlaps");
}
