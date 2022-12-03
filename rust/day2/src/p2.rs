use std::{io::{BufReader, BufRead}, fs::File, error::{Error}, time::Instant, fmt::Display};

const INPUT_FILE: &str = "input.txt"; // Change to input.txt for final solution

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
macro_rules! box_parse_err {
  ($($arg:tt)*) => {
    Err(Box::new(ParseError::InvalidInput(format!($($arg)*))))
  };
}

// #endregion

// #region Structs
type Input = Vec<Round>; // Set input type, defaults to Vec<i32>
struct Round (char, char);

const ROCK: char = 'A';
const PAPER: char = 'B';
const SCISSORS: char = 'C';

// Desired outcomes
const LOSE: char = 'X';
const DRAW: char = 'Y';
const WIN: char = 'Z';

impl Round {
  fn play(&self) -> Result<i32, ParseError> {
    // Calc potential *self* plays to result in loss/draw/win
    let (l_move, d_move, w_move) = if self.0 == ROCK { // Opponent's move
      (SCISSORS, ROCK, PAPER)
    } else if self.0 == PAPER {
      (ROCK, PAPER, SCISSORS)
    }  else if self.0 == SCISSORS {
      (PAPER, SCISSORS, ROCK)
    } else {
      return parse_err!("Invalid move: {}", self.1)
    };
    // Calc score from desired move
    let score = if self.1 == LOSE {
      self.play_bonus(l_move)
    } else if self.1 == DRAW {
      let bonus = self.play_bonus(d_move)?;
      Ok(3 + bonus)
    } else if self.1 == WIN {
      let bonus = self.play_bonus(w_move)?;
      Ok(6 + bonus)
    } else {
      parse_err!("Invalid move: {}", self.1)
    }?;

    Ok(score)
  }

  fn play_bonus(&self, play: char) -> Result<i32, ParseError> {
    // Match stmt cannot include tuple indexing
    if play == ROCK {
      Ok(1)
    } else if play == PAPER {
      Ok(2)
    } else if play == SCISSORS {
      Ok(3)
    } else {
      parse_err!("Invalid move: {}", play)
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
  let mut rounds: Input = Vec::new();
  for l in reader.lines().map(|l| -> String {
    l.unwrap_or(String::new())
  }) {
    if l.len() == 0 {
      continue
    }
    let parts: Vec<&str> = l.split(' ').collect();

    if parts.len() != 2 {
      return box_parse_err!("Invalid line: {}", l);
    }
    let char0 = if let Some(c) = parts[0].chars().nth(0) {
      c
    } else {
      return box_parse_err!("Failed to get char 0")
    };
    let char1 = if let Some(c) = parts[1].chars().nth(0) {
      c
    } else {
      return box_parse_err!("Failed to get char 1")
    };
    rounds.push(Round(char0, char1));
  }
  Ok(rounds)
}

// Solve
fn solve(rounds: &mut Input) -> String {
  let mut count = 0;
  for round in rounds {
    count += round.play().expect("Failed to play round");
  }
  count.to_string()
}

pub fn part2() {
  // Parse
  let mut rounds = parse().expect("Failed to parse input");

  // Solve
  let start = Instant::now();
  let solution = solve(&mut rounds);

  // Report solve time and solution
  let duration = start.elapsed();
  println!("Solved in \x1b[34m{:?}\x1b[0m", duration);
  println!("Solution: \x1b[32m{}\x1b[0m", solution);
}
