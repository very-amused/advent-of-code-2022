use std::{io::{BufReader, BufRead}, fs::File, error::Error, time::Instant};

const INPUT_FILE: &str = "sample.txt"; // Change to input.txt for final solution

// #region Structs
type Input = Vec<i32>; // Set input type, defaults to Vec<i32>

// #endregion

// Parse 
fn parse() -> Result<Input, Box<dyn Error>> {
  // Open scanner to read input line by line
  let file = File::open(INPUT_FILE)?;
  let reader = BufReader::new(file);
  
  // Parse lines
  let mut input: Input = Vec::new();
  for l in reader.lines().map(|l| -> String {
    l.unwrap_or(String::new())
  }) {
    if l.len() == 0 {
    }
  }
  Ok(input)
}

// Solve
fn solve(input: &mut Input) -> String {
  todo!()
}

pub fn part1() {
  // Parse
  let mut input = parse().expect("Failed to parse input");

  // Solve
  let start = Instant::now();
  let solution = solve(&mut input);

  // Report solve time and solution
  let duration = start.elapsed();
  println!("Solved in \x1b[34m{:?}\x1b[0m", duration);
  println!("Solution: \x1b[32m{}\x1b[0m", solution);
}
