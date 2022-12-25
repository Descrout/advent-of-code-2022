use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let file = File::open("input.txt").unwrap();
    let reader = BufReader::new(file);
    let lines: Vec<String> = reader.lines().collect::<Result<_, _>>().unwrap();

    let mut crates = Crates::new(9);
    let mut procs: Vec<Procedure> = Vec::new();

    for (i, line) in lines.iter().enumerate() {
        if i < 8 {
            crates.fill_crates(line);
            continue;
        }

        if i > 9 {
            procs.push(Procedure::new(line));
        }
    }

    let mut crates9001 = crates.clone();

    for proc in procs.iter() {
        crates.apply_proc(proc, false);
        crates9001.apply_proc(proc, true);
    }

    println!(
        "[Part1] Top of each stack CrateMover9000: {}",
        crates.top_of_stacks()
    );
    println!(
        "[Part2] Top of each stack CrateMover9001: {}",
        crates9001.top_of_stacks()
    );
}

#[derive(Clone)]
struct Crates {
    crates: Vec<Vec<String>>,
}

impl Crates {
    pub fn new(size: usize) -> Self {
        Self {
            crates: vec![Vec::new(); size],
        }
    }

    pub fn fill_crates(&mut self, line: &str) {
        let mut j = 0;
        for i in (1..line.chars().count()).step_by(4) {
            let ch = line.chars().nth(i).unwrap();
            if ch != ' ' {
                self.crates[j].insert(0, ch.to_string());
            }
            j += 1;
        }
    }

    pub fn apply_proc(&mut self, proc: &Procedure, is9001: bool) {
        let end = self.crates[proc.from].len();
        let start = end - proc.mv;

        let mut sublist: Vec<String> = self.crates[proc.from].drain(start..end).collect();

        if !is9001 {
            sublist.reverse();
        }

        self.crates[proc.to].append(&mut sublist);
    }

    pub fn top_of_stacks(&self) -> String {
        let mut txt = String::new();
        for c in self.crates.iter() {
            txt += &c[c.len() - 1]
        }
        txt
    }
}

struct Procedure {
    mv: usize,
    from: usize,
    to: usize,
}

impl Procedure {
    pub fn new(txt: &str) -> Self {
        let splitted = txt.split(" ").collect::<Vec<&str>>();
        Self {
            mv: splitted[1].parse::<usize>().unwrap(),
            from: splitted[3].parse::<usize>().unwrap() - 1,
            to: splitted[5].parse::<usize>().unwrap() - 1,
        }
    }
}
