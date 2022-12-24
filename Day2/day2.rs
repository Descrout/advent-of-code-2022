use std::fs::File;
use std::io::{BufRead, BufReader};

// Some languages do negative number modulus different
// So instead of % wrap around
// I put "C" as first and "A" as last
const SYMBOLS: &'static [&str] = &["C", "A", "B", "C", "A"];

fn main() {
    let file = File::open("input.txt").unwrap();
    let reader = BufReader::new(file);

    let mut score_p1 = 0;
    let mut score_p2 = 0;
    for line in reader.lines() {
        let line = line.unwrap();
        let splitted = line.split(" ").collect::<Vec<&str>>();
        let other = splitted[0];
        let mine = splitted[1];
        score_p1 += get_my_score(other, mine);
        score_p2 += get_my_score_part2(other, mine);
    }

    println!("[Part1] Your score is: {}", score_p1);
    println!("[Part2] Your score is: {}", score_p2);
}

fn get_score(letter: &str) -> i32 {
    match letter {
        "A" | "X" => 1,
        "B" | "Y" => 2,
        "C" | "Z" => 3,
        _ => panic!("invalid letter for score"),
    }
}

fn get_my_score(other: &str, mine: &str) -> i32 {
    let my_score = get_score(mine);
    let other_score = get_score(other);

    let diff = my_score - other_score;

    if diff == 0 {
        my_score + 3
    } else if diff == 1 || diff == -2 {
        my_score + 6
    } else {
        my_score
    }
}

fn get_my_score_part2(other: &str, mine: &str) -> i32 {
    let symbol_idx = get_score(other);
    match mine {
        "X" => get_my_score(other, SYMBOLS[(symbol_idx - 1) as usize]),
        "Z" => get_my_score(other, SYMBOLS[(symbol_idx + 1) as usize]),
        _ => get_my_score(other, other),
    }
}
