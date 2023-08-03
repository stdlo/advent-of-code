use regex::Regex;
pub fn process_part1(input: &str) -> i32 {
    let re = Regex::new(r"(\d+)-(\d+) (\w+): (\w+)$").unwrap();

    let mut valid_passwords = 0;

    for line in input.lines() {
        if let Some((_, [min, max, letter, password])) =
            re.captures(line).map(|caps| caps.extract())
        {
            let letter_matches: Vec<_> = password.match_indices(letter).collect();
            let letter_matches_count = letter_matches.len();
            if letter_matches_count < min.parse::<usize>().unwrap() {
                // println!("TOO LITTLE MATCHES");
            } else if letter_matches_count > max.parse::<usize>().unwrap() {
                // println!("TOO MANY MATCHES");
            } else {
                // println!("VALID PASSWORD");
                valid_passwords += 1;
            }
        }
    }
    valid_passwords
}

pub fn process_part2(input: &str) -> i32 {
    let re = Regex::new(r"(\d+)-(\d+) (\w+): (\w+)$").unwrap();

    let mut valid_passwords = 0;

    for line in input.lines() {
        if let Some((_, [pos_1, pos_2, letter, password])) =
            re.captures(line).map(|caps| caps.extract())
        {
            let password_vec: Vec<char> = password.chars().collect();
            let index_1 = pos_1.parse::<usize>().unwrap() - 1;
            let index_2 = pos_2.parse::<usize>().unwrap() - 1;
            let char = letter.chars().nth(0).unwrap();
            let mut matches = 0;
            if password_vec[index_1] == char {
                matches += 1;
            }
            if password_vec[index_2] == char {
                matches += 1;
            }
            if matches == 1 {
                valid_passwords += 1;
            }
        }
    }
    valid_passwords
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc";

    #[test]
    fn part1() {
        let result = process_part1(INPUT);
        assert_eq!(result, 2);
    }

    #[test]
    fn part2() {
        let result = process_part2(INPUT);
        assert_eq!(result, 1);
    }
}
