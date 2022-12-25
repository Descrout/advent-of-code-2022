use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let file = File::open("input.txt").unwrap();
    let reader = BufReader::new(file);
    let lines: Vec<String> = reader.lines().collect::<Result<_, _>>().unwrap();

    let mut pairs: Vec<Pair> = Vec::new();
    for line in lines.iter() {
        pairs.push(Pair::new(line));
    }

    let mut fully_contains = 0;
    let mut overlaps = 0;
    for pair in pairs.iter() {
        if pair.fully_contains() {
            fully_contains += 1;
        }
        if pair.overlaps() {
            overlaps += 1;
        }
    }
    println!("[Part1] Fully contained pairs: {}", fully_contains);
    println!("[Part2] Overlapped pairs: {}", overlaps);
}

struct Pair {
    min1: i32,
    max1: i32,
    min2: i32,
    max2: i32,
}

impl Pair {
    pub fn new(txt: &str) -> Self {
        let comma_split = txt.split(",").collect::<Vec<&str>>();
        let first = comma_split[0].split("-").collect::<Vec<&str>>();
        let second = comma_split[1].split("-").collect::<Vec<&str>>();
        Self {
            min1: first[0].parse::<i32>().unwrap(),
            max1: first[1].parse::<i32>().unwrap(),
            min2: second[0].parse::<i32>().unwrap(),
            max2: second[1].parse::<i32>().unwrap(),
        }
    }

    pub fn overlaps(&self) -> bool {
        (self.min1 >= self.min2 && self.min1 <= self.max2)
            || (self.min2 >= self.min1 && self.min2 <= self.max1)
    }

    pub fn fully_contains(&self) -> bool {
        (self.min1 >= self.min2 && self.max1 <= self.max2)
            || (self.min2 >= self.min1 && self.max2 <= self.max1)
    }
}
