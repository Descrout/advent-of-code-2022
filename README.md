# Advent of Code 2022 in 5 languages

  

Advent of Code is an Advent calendar of small programming puzzles for a variety of skill sets and skill levels that can be solved in any programming language you like. People use them as interview prep, company training, university coursework, practice problems, a speed contest, or to challenge each other.

In this repo, I will try to complete all challenges in 5 different programming languages.
* Dart
* Go
* Javascript
* Python
* Rust

### Run
`cd Day(x)` and run them.  Of course you will need dart, go, nodejs, python and rust to run each desired language.
```makefile
cd Day1

make dart
make go
make js
make python
make rust
```
If you don't have Makefile:
```
cd Day1

dart day1.dart
go run day1.go
node day1.js
python3 day1.py
rustc day1.rs --crate-type bin -o bin/day1_rust.exe && ./bin/day1_rust.exe
```