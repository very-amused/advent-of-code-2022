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
type Input = Vec<Rucksacks>;
struct Rucksacks ([String; 3]);

impl Rucksacks {
  fn common_item(&self) -> Result<char, ParseError> {
    let mut item_map: HashMap<char, u32> = HashMap::new();
    // Split into sacks
    for rs in self.0.iter() {
      let mut char_map: HashMap<char, bool> = HashMap::new();
      for c in rs.chars() {
        if char_map.contains_key(&c) {
          continue
        }
        char_map.insert(c, true);
        let n = item_map.entry(c).and_modify(|n| {
          *n += 1
        }).or_insert(1);
        if *n == 3 {
          return Ok(c);
        }
      }
    }
    parse_err!("No common item found: {:?}", self.0)
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
  let mut group: Vec<String> = Vec::new();
  let mut i = 0;
  for l in reader.lines().map(|l| -> String {
    l.unwrap_or(String::new())
  }) {
    if l.len() == 0 {
      continue;
    }
    group.push(l);
    i += 1;
    if i == 3 {
      let rucksacks = Rucksacks([&group[0], &group[1], &group[2]]
        .map(|s| s.clone()));
      input.push(rucksacks);
      group.clear();
      i = 0;
    }
  }
  Ok(input)
}

// Solve
fn solve(input: &mut Input) -> Result<String, Box<dyn Error>> {
  let mut c = 0;
  for sack in input {
    let badge = sack.common_item()?;
    c += sack.item_priority(badge)?;
  }
  Ok(c.to_string())
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
