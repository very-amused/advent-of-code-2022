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

const ROCK: (char, char) = ('A', 'X');
const PAPER: (char, char) = ('B', 'Y');
const SCISSORS: (char, char) = ('C', 'Z');

impl Round {
  fn play(&self) -> Result<i32, ParseError> {
    // Calc potential opponent plays to result in opponent draw/loss
    let (d_move, l_move) = if self.1 == ROCK.1 { // Your move
      (ROCK.0, SCISSORS.0)
    } else if self.1 == PAPER.1 {
      (PAPER.0, ROCK.0)
    }  else if self.1 == SCISSORS.1 {
      (SCISSORS.0, PAPER.0)
    } else {
      return parse_err!("Invalid move: {}", self.1)
    };
    // Calc score from opponent move
    let mut score = if self.0 == d_move { 3 }
      else if self.0 == l_move { 6 }
      else { 0 };

    score += self.play_bonus(self.1)?;

    Ok(score)
  }

  fn play_bonus(&self, play: char) -> Result<i32, ParseError> {
    // Match stmt cannot include tuple indexing
    if play == ROCK.1 {
      Ok(1)
    } else if play == PAPER.1 {
      Ok(2)
    } else if play == SCISSORS.1 {
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

pub fn part1() {
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
