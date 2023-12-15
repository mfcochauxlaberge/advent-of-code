use std::fs::File;
use std::io::{self, BufRead};

fn main() -> io::Result<()> {
    let file = File::open("input.txt")?;

    let mut total = 0;

    let reader = io::BufReader::new(file);
    for line in reader.lines() {
        let line = line?;
        let mut first_digit = 0;
        let mut first_digit_at = usize::MAX;
        let mut last_digit = 0;
        let mut last_digit_at = 0;

        println!("Before: {}", line);

        let digits = [
            "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
        ];

        for (i, digit) in digits.iter().enumerate() {
            update_digit_positions(
                line.as_str(),
                digit,
                (i + 1) as u32,
                &mut first_digit,
                &mut first_digit_at,
                &mut last_digit,
                &mut last_digit_at,
            );
        }

        for (i, ch) in line.chars().enumerate() {
            if !ch.is_digit(10) {
                continue;
            }

            let d = ch.to_digit(10);

            if i < first_digit_at {
                first_digit = d.unwrap();
                first_digit_at = i;
            }

            if i > last_digit_at {
                last_digit = d.unwrap();
                last_digit_at = i;
            }
        }

        if last_digit_at == 0 {
            last_digit = first_digit;
        }

        println!("{} is {}+{}", line, first_digit, last_digit);

        total += (first_digit * 10) + last_digit;
    }

    println!("Total: {}", total);

    Ok(())
}

fn update_digit_positions(
    line: &str,
    word: &str,
    digit_value: u32,
    first_digit: &mut u32,
    first_digit_at: &mut usize,
    last_digit: &mut u32,
    last_digit_at: &mut usize,
) {
    if let Some(pos) = line.find(word) {
        if pos < *first_digit_at {
            *first_digit = digit_value;
            *first_digit_at = pos;
        }
        if pos > *last_digit_at {
            *last_digit = digit_value;
            *last_digit_at = pos;
        }
    }
    if let Some(pos) = line.rfind(word) {
        if pos < *first_digit_at {
            *first_digit = digit_value;
            *first_digit_at = pos;
        }
        if pos > *last_digit_at {
            *last_digit = digit_value;
            *last_digit_at = pos;
        }
    }
}
