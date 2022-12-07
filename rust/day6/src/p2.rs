use std::{io::{BufReader, BufRead}, fs::File, fmt::Display, error::Error, time::Instant};

const INPUT_FILE: &str = "input.txt"; // Use sample.txt for testing

// #region Parse error
#[derive(Debug)]
enum ParseError {
  InvalidInput(String)
}

impl Display for ParseError {
  fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
    match self {
      Self::InvalidInput(s) => write!(f, "{}", s)
    }
  }
}

impl Error for ParseError {}

macro_rules! parse_err {
  ($($arg:tt)*) => {
    Err(ParseError::InvalidInput(format!($($arg)*)))
  };
}
#[allow(unused_macros)]
macro_rules! box_parse_err {
  ($($arg:tt)*) => {
    Err(Box::new(ParseError::InvalidInput(format!($($arg)*))))
  };
}

// #endregion

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
    todo!();
  }
  Ok(input)
}

// Solve
fn solve(input: &mut Input) -> Result<String, Box<dyn Error>> {
  todo!()
}

pub fn part2() {
  // Parse
  let mut input = parse().expect("Failed to parse input");

  // Solve
  let start = Instant::now();
  let solution = solve(&mut input).expect("Failed to solve");

  // Report solve time and solution
  let duration = start.elapsed();
  println!("Solved in \x1b[34m{:?}\x1b[0m", duration);
  println!("Solution: \x1b[32m{}\x1b[0m", solution);
}
