use day01::process_part1;
use std::fs;

fn main() {
    let input = fs::read_to_string("./input.txt").unwrap();
    println!("{}", process_part1(&input));
}