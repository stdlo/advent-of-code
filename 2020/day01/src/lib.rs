pub fn process_part1(input: &str) -> String {
    "it works".to_string()
}

pub fn process_part2(input: &str) -> String {
    "it works".to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "";

    #[test]
    fn part1() {
        let result = process_part1(INPUT);
        assert_eq!(result, "it works");        
    }
}