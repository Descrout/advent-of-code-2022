use std::fs::File;
use std::io::{BufRead, BufReader};

const TOP_ELF_CALORIES_LEN: i32 = 3;

fn main() {
    let file = File::open("input.txt").unwrap();
    let reader = BufReader::new(file);

    let mut latest_elf_calories: Vec<i32> = Vec::new();
    let mut top_three_elf_calories = vec![0, 0, 0];

    for line in reader.lines() {
        let line = line.unwrap();
        if line.is_empty() {
            let total_calori = sum(&latest_elf_calories);
            update_top_three(&mut top_three_elf_calories, total_calori);
            latest_elf_calories.clear();
            continue;
        }

        let calori = line.parse::<i32>().unwrap();
        latest_elf_calories.push(calori);
    }

    println!(
        "[Part1] Most carried calories: {}",
        top_three_elf_calories[0]
    );
    let top_three_sum = sum(&top_three_elf_calories);
    println!("[Part2] Top three calories sum: {}", top_three_sum);
}

fn sum(arr: &Vec<i32>) -> i32 {
    arr.clone().into_iter().reduce(|a, b| a + b).unwrap()
}

fn update_top_three(top_three_elf_calories: &mut Vec<i32>, new_calories: i32) {
    for i in 0..TOP_ELF_CALORIES_LEN {
        let i = i as usize;
        if new_calories > top_three_elf_calories[i] {
            let old_calori = top_three_elf_calories[i];
            top_three_elf_calories[i] = new_calories;
            update_top_three(top_three_elf_calories, old_calori);
            break;
        }
    }
}
