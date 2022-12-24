use std::collections::{HashMap, HashSet};
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let file = File::open("input.txt").unwrap();
    let reader = BufReader::new(file);
    let lines: Vec<String> = reader.lines().collect::<Result<_, _>>().unwrap();

    let sum_errors = lines.iter().fold(0i32, |acc, line| {
        acc + get_priority(find_error(line)) as i32
    });

    println!("[Part1] Sum of priorities: {}", sum_errors);

    let mut sum_recurring = 0i32;
    for i in (0..lines.len()).step_by(3) {
        let recurring = find_recurring_in_group(vec![&lines[i], &lines[i + 1], &lines[i + 2]]);
        sum_recurring += get_priority(recurring) as i32;
    }
    println!("[Part2] Sum of all groups: {}", sum_recurring);
}

fn get_priority(letter: u8) -> u8 {
    if letter > 96 {
        letter - 96
    } else {
        letter - 38
    }
}

fn find_recurring_in_group(group: Vec<&str>) -> u8 {
    let mut occurences: HashMap<u8, i32> = HashMap::new();
    for compartment in group.iter() {
        let compartment_unique = get_unique_ascii(compartment.to_string());

        for letter in compartment_unique.into_iter() {
            occurences
                .entry(letter)
                .and_modify(|occurence| *occurence += 1)
                .or_insert(1);
        }
    }

    return get_three_occurence_letter(occurences);
}

fn get_three_occurence_letter(map: HashMap<u8, i32>) -> u8 {
    for (key, value) in map.into_iter() {
        if value == 3 {
            return key;
        }
    }
    panic!("could not find 3 times recurring letter in a group");
}

fn get_unique_ascii(text: String) -> HashSet<u8> {
    text.as_bytes()
        .iter()
        .fold(HashSet::new(), |mut prev, code| {
            prev.insert(*code);
            prev
        })
}

fn find_error(compartment: &str) -> u8 {
    let mut occurence: HashSet<u8> = HashSet::new();
    let ascii_list = compartment.as_bytes();
    let half = ascii_list.len() / 2;

    for i in 0..ascii_list.len() {
        let code_unit = ascii_list[i];
        if i < half {
            occurence.insert(code_unit);
        } else if occurence.contains(&code_unit) {
            return code_unit;
        }
    }

    panic!("could not find any recurring letter");
}
