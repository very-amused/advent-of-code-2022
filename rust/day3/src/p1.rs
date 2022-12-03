use std::{io::{BufReader, BufRead}, fs::File, fmt::Display, error::Error, time::Instant, collections::HashMap};

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

#[allow(unused_macros)]
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
type Input = Vec<Rucksack>;
struct Rucksack (String);

impl Rucksack {
  fn common_item(&self) -> Result<char, ParseError> {
    // Split into compartments
    let c1 = &self.0[..self.0.len()/2];
    let c2 = &self.0[self.0.len()/2..];
    let mut item_map: HashMap<char, bool> = HashMap::new();
    for c in c1.chars() {
      item_map.insert(c, true);
    }
    for c in c2.chars() {
      if item_map.contains_key(&c) {
        return Ok(c);
      }
    }
    parse_err!("No common item found: {}", self.0)
  }

  fn item_priority(&self, item: char) -> Result<u32, ParseError> {
    if item >= 'a' && item <= 'z' {
      Ok((item as u32 - 'a' as u32) + 1) // 1-26
    } else if item >= 'A' && item <= 'Z' {
      Ok((item as u32 - 'A' as u32) + 27) // 27-52
    } else {
      parse_err!("Invalid item char: {}", item)
    }
  }
}

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
      continue
    }
    input.push(Rucksack(l));
  }
  Ok(input)
}

// Solve
fn solve(input: &mut Input) -> String {
  let mut c = 0;
  for sack in input {
    let item = sack.common_item().expect("Failed to find common item");
    c += sack.item_priority(item).expect("Failed to get item priority");
  }
  c.to_string()
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
