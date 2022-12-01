use std::{io::{BufReader, BufRead}, fs::File, error::Error, time::Instant};

const INPUT_FILE: &str = "input.txt";

// #region Structs
type Input = Vec<i32>;

// #endregion

// Parse 
fn parse() -> Result<Input, Box<dyn Error>> {
  // Open scanner to read input line by line
  let file = File::open(INPUT_FILE)?;
  let reader = BufReader::new(file);
  
  // Parsing state vars go here (if any)
  let mut c = 0;
  let mut elves: Input = Vec::new();
  for l in reader.lines().map(|l| -> String {
    l.unwrap_or(String::new())
  }) {
    if l.len() == 0 {
      elves.push(c);
      c = 0;
      continue;
    }

    let cals: i32 = l.parse()?;
    c += cals;
  }
  Ok(elves)
}

// Solve
fn solve(input: &mut Input) -> String {
  // sum top 3 elves
  input.sort();
  let mut sum = 0;
  for c in input[input.len()-3..].iter() {
    sum += c;
  }
  format!("{}", sum)
}

pub fn part2() {
  // Parse
  let mut input = parse().expect("Failed to parse input");

  // Solve
  let start = Instant::now();
  let solution = solve(&mut input);

  // Report solve time and solution
  let duration = start.elapsed();
  println!("Solved in \x1b[34m{:?}\x1b[0m", duration);
  println!("Solution: \x1b[32m{}\x1b[0m", solution)
}