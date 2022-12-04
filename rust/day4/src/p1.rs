use std::{io::{BufReader, BufRead}, fs::File, fmt::Display, error::Error, time::Instant, num::ParseIntError, str::FromStr};

const INPUT_FILE: &str = "input.txt"; // Use sample.txt for testing

// #region Parse error
#[derive(Debug)]
enum ParseError {
  InvalidInput(String),
  ParseIntError(ParseIntError)
}

impl Display for ParseError {
  fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
    write!(f, "{}", self)
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
type Input = Vec<Pair>; // Set input type, defaults to Vec<i32>
struct Pair (Assignment, Assignment);

impl Pair {
  fn is_redundant(&self) -> bool {
    self.0.contains(&self.1) || self.1.contains(&self.0)
  }
}

struct Assignment {
  start: u32,
  end: u32
}

impl Assignment {
  fn contains(&self, a2: &Assignment) -> bool {
    self.start <= a2.start && self.end >= a2.end
  }
}

impl FromStr for Assignment {
  type Err = ParseError;

  fn from_str(r: &str) -> Result<Self, Self::Err> {
    let parts: Vec<&str> = r.split("-").collect();
    if parts.len() != 2 {
      return parse_err!("Unexpected number of parts in assignment range: {}", r);
    }
    let start: u32 = parts[0].parse()
      .map_err(Self::Err::ParseIntError)?;
    let end: u32 = parts[1].parse()
      .map_err(Self::Err::ParseIntError)?;
    Ok(Assignment{
      start,
      end
    })
  }
}

// #endregion

// Parse 
fn parse() -> Result<Input, Box<dyn Error>> {
  // Open scanner to read input line by line
  let file = File::open(INPUT_FILE)?;
  let reader = BufReader::new(file);
  
  // Parse lines
  let mut pairs: Input = Vec::new();
  for l in reader.lines().map(|l| -> String {
    l.unwrap_or(String::new())
  }) {
    if l.len() == 0 {
      continue;
    }
    let parts: Vec<&str> = l.split(",").collect();
    if parts.len() != 2 {
      return box_parse_err!("Unexpected number of assignments in pair: {l}");
    }
    let pair = Pair(
      parts[0].parse()?,
      parts[1].parse()?
    );
    pairs.push(pair);
  }
  Ok(pairs)
}

// Solve
fn solve(pairs: &mut Input) -> Result<String, Box<dyn Error>> {
  let mut c: u32 = 0;
  for p in pairs {
    if p.is_redundant() {
      c += 1;
    }
  }
  Ok(c.to_string())
}

pub fn part1() {
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
