use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let file = File::open("input.txt").unwrap();
    let reader = BufReader::new(file);

    let mut latest_elf_calories: Vec<i32> = Vec::new();
    let mut max_elf_calori = 0i32;

    for line in reader.lines() {
        let line = line.unwrap();
        if line.is_empty() {
            let total_calori = latest_elf_calories
                .clone()
                .into_iter()
                .reduce(|a, b| a + b)
                .unwrap();
            if total_calori > max_elf_calori {
                max_elf_calori = total_calori;
            }
            latest_elf_calories.clear();
            continue;
        }

        let calori = line.parse::<i32>().unwrap();
        latest_elf_calories.push(calori);
    }

    println!("Most carried calories: {}", max_elf_calori);
}
