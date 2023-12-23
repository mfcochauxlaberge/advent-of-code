use std::fs::File;
use std::io::{self, BufRead};

fn main() -> io::Result<()> {
    let file = File::open("input.txt")?;

    let mut total = 0;
    let mut total_power = 0;

    let reader = io::BufReader::new(file);
    for line in reader.lines() {
        let line = line?;

        let name = line.split(":").nth(0).unwrap();
        let id_str = name.split(" ").last().unwrap();
        let id = id_str.parse::<u32>().unwrap();

        let grabs = line.split(":").nth(1).unwrap().split(";");

        let mut max_red = 0;
        let mut max_green = 0;
        let mut max_blue = 0;

        let mut valid = true;

        grabs.enumerate().for_each(|(_, grab)| {
            let mut num_red = 0;
            let mut num_green = 0;
            let mut num_blue = 0;

            grab.split(',').enumerate().for_each(|(_, mut color)| {
                color = color.trim();

                let num = color.split(" ").nth(0).unwrap().parse::<u32>().unwrap();

                match color.split(" ").nth(1).unwrap().trim() {
                    color if color.contains("red") => num_red += num,
                    color if color.contains("green") => num_green += num,
                    color if color.contains("blue") => num_blue += num,
                    _ => (),
                }
            });

            max_red = max_red.max(num_red);
            max_green = max_green.max(num_green);
            max_blue = max_blue.max(num_blue);

            // 12 red cubes, 13 green cubes, and 14 blue cubes
            if num_red > 12 || num_green > 13 || num_blue > 14 {
                valid = false;
            }
        });

        let power = max_red * max_green * max_blue;
        total_power += power;

        if valid {
            total += id;
        }
    }

    println!("Total: {}", total);
    println!("Total power: {}", total_power);

    Ok(())
}
